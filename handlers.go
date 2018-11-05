package main

import (
	"context"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

// Test struct for testing
type Test struct {
	Title     string
	ImgEncode []string
}

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
	err = storeImage(imgBin)
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

	// w.Write([]byte("This seems to work"))
	fmt.Fprintf(w, "This seems to work")
}

func testHandler(w http.ResponseWriter, r *http.Request) {

	counter := 1

	// Create connection to mongo
	conn, err := mongo.Connect(context.Background(), "mongodb://admin:connecttome123@ds151533.mlab.com:51533/stegano", nil)
	if err != nil {
		panic(err)
	}
	coll := conn.Database("stegano").Collection("images")

	// Fetch image from mongo
	cur, err := coll.Find(context.Background(), nil)
	if err != nil {
		panic(err)
	}
	var img map[string]interface{}       // Here we'll store fetched images
	for cur.Next(context.Background()) { // Iterate the cursor
		err := cur.Decode(&img) // Store fetched images
		if err != nil {
			panic(err)
		}

		// Save new image here
		err = ioutil.WriteFile("assets/images/plain/image"+strconv.Itoa(counter)+".png", bson.Binary(img["imgBin"].(bson.Binary)).Data, 0644)
		if err != nil {
			panic(err)
		}
		counter++

	}

	// HTML Templating
	t, err := template.ParseFiles("assets/html/test.html")
	if err != nil {
		panic(err)
	}

	test := Test{Title: "Best page title"}

	for i := 1; i < counter; i++ {
		test.ImgEncode = append(test.ImgEncode, "assets/images/plain/image"+strconv.Itoa(i)+".png")
	}

	t.Execute(w, test)

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

	ciphertext := encodeCaesar(plaintext)

	fmt.Fprint(w, ciphertext)

}

