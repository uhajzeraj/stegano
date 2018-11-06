package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/gorilla/securecookie"
	"github.com/gorilla/sessions"
	"golang.org/x/crypto/bcrypt"
)

// Test struct for testing
type Test struct {
	Title     string
	ImgEncode []string
}

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

	router.HandleFunc("/test", testHandler).Methods("GET")

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
	counter := 1

	// Fetch images from Mongo
	images, err := getImages(sessionUser)
	returnEmptyError(err)

	for _, val := range images {

		// Save new image here
		err = ioutil.WriteFile("assets/images/"+sessionUser+strconv.Itoa(counter)+".png", val, 0644)
		if err != nil {
			fmt.Println(err)
		}

		counter++
	}

	savedImages := SavedData{}

	for i := 1; i < counter; i++ {
		savedImages.Images = append(savedImages.Images, "assets/images/"+sessionUser+strconv.Itoa(i)+".png")
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

	http.Redirect(w, r, "/home", http.StatusSeeOther)

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

	sessionUser := session.Values["user"].(string)

	// Store the image into DB
	err = storeImage(sessionUser, imgBin)
	if err != nil {
		return
	}

	fmt.Fprint(w, 1)
}

func testHandler(w http.ResponseWriter, r *http.Request) {

	// Check if session is set
	session, err := store.Get(r, "user-login")
	returnEmptyError(err)
	if len(session.Values) == 0 {
		http.Redirect(w, r, "/", http.StatusSeeOther) // Redirect to root
		return
	}

	counter := 1

	// Get the user based on the session info
	sessionUser := session.Values["user"].(string)

	err = decode(sessionUser)
	if err != nil {
		panic(err)
	}

	// Fetch images from Mongo
	images, err := getImages(sessionUser)
	returnEmptyError(err)

	for _, val := range images {

		// Save new image here
		err = ioutil.WriteFile("assets/images/"+sessionUser+strconv.Itoa(counter)+".png", val, 0644)
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
		test.ImgEncode = append(test.ImgEncode, "assets/images/"+sessionUser+strconv.Itoa(i)+".png")
	}

	err = t.Execute(w, test)
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
