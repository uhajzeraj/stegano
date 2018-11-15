package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strings"
	"time"

	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"github.com/gorilla/context"
	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
)

var (
	conn, _ = mgo.Dial("mongodb://admin:connecttome123@ds151533.mlab.com:51533/stegano")
)

// UserInfo struct
type UserInfo struct {
	User     string `bson:"user"`
	HashPass string `bson:"passHash"`
}

func main() {

	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/admin", adminHandler).Methods("POST")
	router.HandleFunc("/admin/{user}", adminDeleteHandler).Methods("DELETE")

	fmt.Print("Admin")

	srv := &http.Server{
		Handler: context.ClearHandler(router),
		Addr:    ":8080",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}

func adminDeleteHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	pathVars := mux.Vars(r)

	if len(pathVars) != 1 {
		http.Error(w, "400 - Bad Request, too many URL arguments.", http.StatusBadRequest)
		return
	}

	err := conn.DB("stegano").C("user").Remove(bson.M{"user": pathVars["user"]})

	if err != nil {
		http.Error(w, "404 - User not found!", 404)
		return
	}

	resp := `{`
	resp += `"user":`
	resp += pathVars["user"]
	resp += `}`

	fmt.Fprint(w, resp)
}

func adminHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	adminUser := &UserInfo{}

	err := json.NewDecoder(r.Body).Decode(&adminUser)

	if err != nil {
		http.Error(w, err.Error(), 400) //checking for errors in the process and returning bad request if so
		return
	}

	errorSlice := validateLogin(adminUser.User, adminUser.HashPass)

	if len(errorSlice) > 0 {
		for _, val := range errorSlice {
			fmt.Println(val)
		}
		return
	}
	iter := conn.DB("stegano").C("user").Find(nil).Iter()

	resp := `{`
	resp += `"user":[`
	for iter.Next(&adminUser) {
		resp += adminUser.User
		resp += `,`
	}
	if err := iter.Close(); err != nil {
		return
	}
	resp = strings.TrimRight(resp, ",")
	resp += `]}`

	fmt.Fprint(w, resp)

}

func validateLogin(user, pass string) []string {

	var errorSlice []string

	// Check if username is OK
	match, err := regexp.MatchString(`^[a-zA-Z0-9_-]{6,30}$`, user)
	returnEmptyError(err)
	if !match {
		errorSlice = append(errorSlice, "Username does not meet the requirements")
	}

	// Check if password is OK
	match, err = regexp.MatchString(`^.{6,40}$`, pass)
	returnEmptyError(err)
	if !match {
		errorSlice = append(errorSlice, "Password does not meet the requirements")
	}

	// Continue further if there are no errors
	if len(errorSlice) == 0 {

		result := UserInfo{}

		conn.DB("stegano").C("admin").Find(bson.M{"user": user}).Select(bson.M{"user": user, "passHash": 1}).One(&result)

		err := bcrypt.CompareHashAndPassword([]byte(result.HashPass), []byte(pass))
		if err != nil {
			errorSlice = append(errorSlice, "Username or password is not correct")
		}
	}

	return errorSlice
}

func returnEmptyError(err error) {
	if err != nil {
		return
	}
}
