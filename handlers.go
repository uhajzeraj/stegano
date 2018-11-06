package main

import (

	//"github.com/mongodb/mongo-go-driver/core/result"
	//"context"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"regexp"
	"strconv"
	"github.com/gorilla/sessions"
	"github.com/globalsign/mgo/bson"
	"github.com/gorilla/mux"

	//"github.com/mongodb/mongo-go-driver/mongo"
	//"github.com/mongodb/mongo-go-driver/bson"
	//"github.com/mongodb/mongo-go-driver/mongo"
	"golang.org/x/crypto/bcrypt"
	//"log"
)

// Test struct for testing
type Test struct {
	Title     string
	ImgEncode []string
}

type PassHASH struct {
	User string `bson:"user"`
}

//Session var
var store = sessions.NewCookieStore([]byte("asdaskdhasdhgsajdgasdsadksakdhasidoajsdousahdopj"))

// Wrap mux router in a function for testing
func newRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	// Diferent path - method handlers
	router.HandleFunc("/", rootHandler).Methods("GET")
	router.HandleFunc("/signup", signupGetHandler).Methods("GET")
	router.HandleFunc("/signup", signupPostHandler).Methods("POST")
	router.HandleFunc("/login", loginHandler).Methods("GET")
	router.HandleFunc("/test", testHandler).Methods("GET")
	router.HandleFunc("/stegano", steganoGetHandler).Methods("GET")
	router.HandleFunc("/stegano", steganoPostHandler).Methods("POST")

	router.HandleFunc("/caesar", caesarGetHandler).Methods("GET")
	router.HandleFunc("/caesar", caesarPostHandler).Methods("POST")
	router.HandleFunc("/login", postLoginHandler).Methods("POST")
	router.HandleFunc("/home", homeHandler)

	// Static file directory
	staticFileDirectory := http.Dir("./assets/")
	staticFileHadler := http.StripPrefix("/assets/", http.FileServer(staticFileDirectory))
	router.PathPrefix("/assets/").Handler(staticFileHadler).Methods("GET")

	return router
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("assets/html/index.html")
	if err != nil {
		panic(err)
	}
	t.Execute(w, nil /*Test{"Best page title", `assets/images/plain/smaller_image.jpeg`}*/)
}

func signupGetHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("assets/html/signup.html")
	if err != nil {
		panic(err)
	}
	t.Execute(w, nil)
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("assets/html/login.html")
	if err != nil {
		panic(err)
	}
	t.Execute(w, nil)
}

func steganoGetHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("assets/html/stegano.html")
	if err != nil {
		panic(err)
	}
	t.Execute(w, nil)
}

func steganoPostHandler(w http.ResponseWriter, r *http.Request) {

	// Check if textfield is empty
	if len(r.FormValue("text")) == 0 {
		fmt.Println("Empty text field")
		return
	}

	// Check if image is uploaded successfully
	file, _, err := r.FormFile("image")
	if err != nil {
		fmt.Println("Empty image")
		return
	}

	text := r.FormValue("text") // Get the text received

	err = file.Close() // Close the file
	if err != nil {
		return
	}

	imgBin, err := encode(file, text)
	if err != nil {
		return
	}

	// Store the image into DB
	err = storeImage("uranii", imgBin)
	if err != nil {
		return
	}

}

func signupPostHandler(w http.ResponseWriter, r *http.Request) {

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

	session, err := store.Get(r, "user-login")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Set some session values.
	session.Values["user"] = user
	//session.Values[] = 43
	// Save it before we write to the response/return from the handler.
	session.Save(r, w)

	http.Redirect(w, r, "/home", http.StatusSeeOther)

	// w.Write([]byte("This seems to work"))
	fmt.Fprintf(w, "This seems to work")

}

func testHandler(w http.ResponseWriter, r *http.Request) {

	counter := 1

	// Get the user based on the session info
	sessionUser := "uranii" // This needs to be changed

	err := decode(sessionUser)
	if err != nil {
		panic(err)
	}

	// Fetch images from Mongo
	images, err := getImages(sessionUser)
	returnEmptyError(err)

	for _, val := range images {

		// Save new image here
		err = ioutil.WriteFile("assets/images/"+sessionUser+"img"+strconv.Itoa(counter)+".png", val, 0644)
		if err != nil {
			fmt.Println(err)
		}

		counter++

	}

	// HTML Templating
	t, err := template.ParseFiles("assets/html/test.html")
	returnEmptyError(err)

	test := Test{Title: "Best page title"}

	for i := 1; i < counter; i++ {
		test.ImgEncode = append(test.ImgEncode, "assets/images/"+sessionUser+"img"+strconv.Itoa(i)+".png")
	}

	err = t.Execute(w, test)
	// if err == nil {
	// 	// os.RemoveAll("assets/images/uranii")
	// }
}

/* CAESAR's CIPHER */
func caesarGetHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("assets/html/caesar.html")
	if err != nil {
		panic(err)
	}
	t.Execute(w, nil)
}

func caesarPostHandler(w http.ResponseWriter, r *http.Request) {

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

func postLoginHandler(w http.ResponseWriter, r *http.Request) {

	var errorSlice []string

	email := r.FormValue("username")
	pass := r.FormValue("pass")

	fmt.Println(email)

	match, err := regexp.MatchString("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$", email)
	returnEmptyError(err)
	if !match {
		errorSlice = append(errorSlice, "Email does not meet the requirements")
	}

	hashPass, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	returnEmptyError(err)

	coll := conn.DB("stegano").C("users")
	result := PassHASH{}
	err = coll.Find(bson.M{"email": email}).Select(bson.M{"user": 1}).One(&result)
	fmt.Println(result)
	if err != nil {
		panic(err)
	}

	if result.User == "" {
		//TODO handle username or password incorrect
		http.Error(w, "404",404)
		return
	}

	session, err := store.Get(r, "user-login")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Set some session values.
	session.Values["hash"] = string(hashPass)
	//session.Values[] = 43
	// Save it before we write to the response/return from the handler.
	session.Save(r, w)

	http.Redirect(w, r, "/home", http.StatusSeeOther)

}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	session, err := store.Get(r, "user-login")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if session.Values["hash"] == "" {
		//TODO: don't show anything
		http.Error(w, err.Error(), 404)
		return
	}

	t, err := template.ParseFiles("assets/html/home.html")
	if err != nil {
		panic(err)
	}
	t.Execute(w, nil)
}

