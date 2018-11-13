package main

import (
	"encoding/json"

	"github.com/globalsign/mgo/bson"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gorilla/mux"
	"github.com/gorilla/securecookie"
	"github.com/gorilla/sessions"
	"golang.org/x/crypto/bcrypt"
)

// SavedData structure for showing images and their info
type SavedData struct {
	Images []string
}

// User struct
type User struct {
	User string `bson:"user"`
}

//Session var
var store = sessions.NewCookieStore(securecookie.GenerateRandomKey(32))

// Wrap mux router in a function for testing
func newRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	// Diferent path - method handlers
	router.HandleFunc("/", rootHandler).Methods("GET")
	router.HandleFunc("/home", homeHandler).Methods("GET")
	router.HandleFunc("/user", userHandler).Methods("GET")
	router.HandleFunc("/saved", savedHandler).Methods("GET")

	router.HandleFunc("/logout", logoutHandler).Methods("GET")

	router.HandleFunc("/login", loginGetHandler).Methods("GET")
	router.HandleFunc("/login", loginPostHandler).Methods("POST")

	router.HandleFunc("/signup", signupGetHandler).Methods("GET")
	router.HandleFunc("/signup", signupPostHandler).Methods("POST")

	router.HandleFunc("/stegano", steganoGetHandler).Methods("GET")
	router.HandleFunc("/stegano", steganoPostHandler).Methods("POST")

	router.HandleFunc("/caesar", caesarGetHandler).Methods("GET")
	router.HandleFunc("/caesar", caesarPostHandler).Methods("POST")

	router.HandleFunc("/deleteImg", deleteImgPostHandler).Methods("POST")
	router.HandleFunc("/admin", adminHandler).Methods("POST")
	router.HandleFunc("/admin/{user}", adminDeleteHandler).Methods("DELETE")


	// Static file directory
	staticFileDirectory := http.Dir("./assets/")
	staticFileHadler := http.StripPrefix("/assets/", http.FileServer(staticFileDirectory))
	router.PathPrefix("/assets/").Handler(staticFileHadler).Methods("GET")

	return router
}

func rootHandler(w http.ResponseWriter, r *http.Request) {

	// Check if session is set
	session, err := store.Get(r, "user-login")
	returnEmptyError(err)
	if len(session.Values) != 0 {
		http.Redirect(w, r, "/home", http.StatusSeeOther) // Redirect to root
		return
	}

	t, err := template.ParseFiles("assets/html/index.html")
	if err != nil {
		panic(err)
	}
	t.Execute(w, nil)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {

	// Check if session is set
	session, err := store.Get(r, "user-login")
	returnEmptyError(err)
	if len(session.Values) == 0 {
		http.Redirect(w, r, "/", http.StatusSeeOther) // Redirect to root
		return
	}

	t, err := template.ParseFiles("assets/html/home.html")
	if err != nil {
		panic(err)
	}
	t.Execute(w, nil)
}

func userHandler(w http.ResponseWriter, r *http.Request) {

	// Check if session is set
	session, err := store.Get(r, "user-login")
	returnEmptyError(err)
	if len(session.Values) == 0 {
		http.Redirect(w, r, "/", http.StatusSeeOther) // Redirect to root
		return
	}

	t, err := template.ParseFiles("assets/html/user.html")
	if err != nil {
		panic(err)
	}
	t.Execute(w, nil)
}

func savedHandler(w http.ResponseWriter, r *http.Request) {

	// Check if session is set
	session, err := store.Get(r, "user-login")
	returnEmptyError(err)
	if len(session.Values) == 0 {
		http.Redirect(w, r, "/", http.StatusSeeOther) // Redirect to root
		return
	}

	sessionUser := session.Values["user"].(string)

	// Fetch images from Mongo
	images, err := getImages(sessionUser)
	returnEmptyError(err)

	savedImages := SavedData{}

	for _, val := range images {

		// Save new image here
		err = ioutil.WriteFile("assets/images/"+val.Name+".png", val.Img, 0644)
		returnEmptyError(err)

		// Append the path for templating
		savedImages.Images = append(savedImages.Images, "assets/images/"+val.Name+".png")
	}

	t, err := template.ParseFiles("assets/html/savedData.html")
	if err != nil {
		panic(err)
	}
	t.Execute(w, savedImages)
}

func logoutHandler(w http.ResponseWriter, r *http.Request) {

	// Check if session is set
	session, err := store.Get(r, "user-login")
	returnEmptyError(err)
	if len(session.Values) == 0 {
		http.Redirect(w, r, "/", http.StatusSeeOther) // Redirect to root
		return
	}

	// Remove session
	delete(session.Values, "user")
	session.Save(r, w)
	// Redirect to index
	http.Redirect(w, r, "/", http.StatusSeeOther) // Redirect to root
}

func loginGetHandler(w http.ResponseWriter, r *http.Request) {

	// Check if session is set
	session, err := store.Get(r, "user-login")
	returnEmptyError(err)
	if len(session.Values) != 0 {
		http.Redirect(w, r, "/home", http.StatusSeeOther) // Redirect to home
		return
	}

	t, err := template.ParseFiles("assets/html/login.html")
	if err != nil {
		panic(err)
	}
	t.Execute(w, nil)
}

func loginPostHandler(w http.ResponseWriter, r *http.Request) {

	// Check if session is set
	session, err := store.Get(r, "user-login")
	returnEmptyError(err)
	if len(session.Values) != 0 {
		http.Redirect(w, r, "/home", http.StatusSeeOther) // Redirect to root
		return
	}

	user := r.FormValue("username")
	pass := r.FormValue("pass")

	errorSlice := validateLogin(user, pass)

	if len(errorSlice) > 0 {
		for _, val := range errorSlice {
			fmt.Println(val)
		}
		return
	}

	// Set some session values.
	session.Values["user"] = user
	// Save it before we write to the response/return from the handler.
	session.Save(r, w)

	fmt.Fprint(w, 1)

}

func signupGetHandler(w http.ResponseWriter, r *http.Request) {

	// Check if session is set
	session, err := store.Get(r, "user-login")
	returnEmptyError(err)
	if len(session.Values) != 0 {
		http.Redirect(w, r, "/home", http.StatusSeeOther) // Redirect to home
		return
	}

	t, err := template.ParseFiles("assets/html/signup.html")
	if err != nil {
		panic(err)
	}
	t.Execute(w, nil)
}

func signupPostHandler(w http.ResponseWriter, r *http.Request) {

	session, err := store.Get(r, "user-login")
	returnEmptyError(err)
	// Check if session is set
	if len(session.Values) != 0 {
		http.Redirect(w, r, "/", http.StatusSeeOther) // Redirect to homepage
		return
	}

	// Get the fields
	user := r.FormValue("displayName")
	email := r.FormValue("email")
	pass := r.FormValue("pass")
	passConfirm := r.FormValue("passConfirm")

	// Validate the fields
	errorSlice := validateSignup(user, pass, passConfirm, email)

	// Check if there are any errors
	if len(errorSlice) > 0 {
		for _, val := range errorSlice {
			fmt.Println(val)
		}
		return
	}

	// Hash the password
	hashPass, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	returnEmptyError(err)

	// Add the user in the database
	err = addUser(user, email, string(hashPass))
	returnEmptyError(err)

	// Set some session values.
	session.Values["user"] = user
	// Save it before we write to the response/return from the handler.
	session.Save(r, w)

	fmt.Fprint(w, 1)

}

func steganoGetHandler(w http.ResponseWriter, r *http.Request) {

	session, err := store.Get(r, "user-login")
	returnEmptyError(err)
	// Check if session is set
	if len(session.Values) == 0 {
		http.Redirect(w, r, "/", http.StatusSeeOther) // Redirect to homepage
		return
	}

	t, err := template.ParseFiles("assets/html/stegano.html")
	if err != nil {
		panic(err)
	}
	t.Execute(w, nil)
}

func steganoPostHandler(w http.ResponseWriter, r *http.Request) {

	session, err := store.Get(r, "user-login")
	returnEmptyError(err)
	// Check if session is set
	if len(session.Values) == 0 {
		http.Redirect(w, r, "/", http.StatusSeeOther) // Redirect to homepage
		return
	}

	// Check if textfield is empty
	if len(r.FormValue("text")) == 0 {
		fmt.Println("Empty text field")
		return
	}

	// Check if image is uploaded successfully
	file, fileHeader, err := r.FormFile("image")
	if err != nil {
		return
	}

	title := []byte(time.Now().String() + fileHeader.Filename)
	shaSum := sha256.Sum224(title)

	fileName := hex.EncodeToString(shaSum[:])

	text := r.FormValue("text") // Get the text received

	err = file.Close() // Close the file
	if err != nil {
		return
	}

	imgBin, err := encode(file, text)
	if err != nil {
		return
	}

	sessionUser := session.Values["user"].(string)

	// Store the image into DB
	err = storeImage(sessionUser, fileName, imgBin)
	if err != nil {
		return
	}

	fmt.Fprint(w, 1)
}

/* CAESAR's CIPHER */
func caesarGetHandler(w http.ResponseWriter, r *http.Request) {

	// Check if session is set
	session, err := store.Get(r, "user-login")
	returnEmptyError(err)
	if len(session.Values) == 0 {
		http.Redirect(w, r, "/", http.StatusSeeOther) // Redirect to root
		return
	}

	t, err := template.ParseFiles("assets/html/caesar.html")
	if err != nil {
		panic(err)
	}
	t.Execute(w, nil)
}

func caesarPostHandler(w http.ResponseWriter, r *http.Request) {

	// Check if session is set
	session, err := store.Get(r, "user-login")
	returnEmptyError(err)
	if len(session.Values) == 0 {
		http.Redirect(w, r, "/", http.StatusSeeOther) // Redirect to root
		return
	}

	// Check if textfield is empty
	if len(r.FormValue("plaintext")) == 0 {
		fmt.Println("Empty text field")
		return
	}

	plaintext := r.FormValue("plaintext") // Get the text received
	shiftSize, err := strconv.Atoi(r.FormValue("shiftSize"))
	if err != nil {
		return
	}

	ciphertext := encodeCaesar(plaintext, shiftSize)

	fmt.Fprint(w, ciphertext)

}

func deleteImgPostHandler(w http.ResponseWriter, r *http.Request) {

	// Check if session is set
	session, err := store.Get(r, "user-login")
	returnEmptyError(err)
	if len(session.Values) == 0 {
		http.Redirect(w, r, "/", http.StatusSeeOther) // Redirect to root
		return
	}

	// Check if textfield is empty
	if len(r.FormValue("imgName")) == 0 {
		return
	}

	sessionUser := session.Values["user"].(string)

	// Trim the name of the image from the unneeded parts
	imgName := r.FormValue("imgName")
	imgName = strings.TrimPrefix(imgName, "assets/images/")
	imgName = strings.TrimSuffix(imgName, ".png")

	err = removeImage(sessionUser, imgName)
	returnEmptyError(err)

	fmt.Fprint(w, 1)
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

	resp := "{\n"
	resp += `"user":\n`
	resp += pathVars["user"]
	resp += "\n}"

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

	resp := "{\n"
	resp += `"user":[\n`
	for iter.Next(&adminUser) {
		resp += adminUser.User
		resp += ","
	}
	if err := iter.Close(); err != nil {
		return
	}
	resp = strings.TrimRight(resp, ",")
	resp += "]\n}"

	fmt.Fprint(w, resp)

}
