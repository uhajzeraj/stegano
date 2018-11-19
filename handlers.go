package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"html/template"
	"io"
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
	Image   string
	Message string
}

// User struct
type User struct {
	User string `bson:"user"`
}

// UserTemplate structure for showing user info
type UserTemplate struct {
	User  string
	Email string
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
	router.HandleFunc("/steganoDecode", steganoDecodeHandler).Methods("POST")

	router.HandleFunc("/caesar", caesarGetHandler).Methods("GET")
	router.HandleFunc("/caesar", caesarPostHandler).Methods("POST")
	router.HandleFunc("/caesarDecode", caesarDecodeHandler).Methods("POST")

	router.HandleFunc("/rot13", rot13GetHandler).Methods("GET")
	router.HandleFunc("/rot13", rot13PostHandler).Methods("POST")

	router.HandleFunc("/deleteImg", deleteImgPostHandler).Methods("POST")

	router.HandleFunc("/changePass", changePassPostHandler).Methods("POST")
	router.HandleFunc("/deleteAcc", deleteAccDeleteHandler).Methods("DELETE")

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
		return
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
		return
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
		return
	}

	sessionUser := session.Values["user"].(string)
	email, err := findEmail(sessionUser)
	returnEmptyError(err)

	tmpl := UserTemplate{sessionUser, email}

	t.Execute(w, tmpl)
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

	savedImages := []SavedData{}
	for _, val := range images {

		savedImage := SavedData{}

		// Decode the mesagge
		message, err := decode(val.Img)
		returnEmptyError(err)

		// Save new image here
		err = ioutil.WriteFile("assets/images/"+val.Name+".png", val.Img, 0644)
		returnEmptyError(err)

		// Append the path and message for templating
		savedImage.Image = "assets/images/" + val.Name + ".png"
		savedImage.Message = message
		savedImages = append(savedImages, savedImage)
	}

	t, err := template.ParseFiles("assets/html/savedData.html")
	if err != nil {
		return
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
		return
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
		response, err := json.Marshal(errorSlice)
		returnEmptyError(err)
		
		storeFailedLogin(user, time.Now())
		i := countLogins(user)
		fmt.Println(i)
		if i%3 == 0 {
			http.Redirect(w, r, "10.212.138.222:8080/admin/email/"+user+"", http.StatusSeeOther)
			return
		}
		fmt.Fprint(w, string(response))
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
		return
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
		response, err := json.Marshal(errorSlice)
		returnEmptyError(err)
		fmt.Fprint(w, string(response))
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
		return
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

	var errorSlice []string

	// Check if textfield is empty
	if len(r.FormValue("text")) == 0 {
		errorSlice = append(errorSlice, "Empty text field")
	}

	// Check if image is uploaded successfully
	file, fileHeader, err := r.FormFile("image")
	if err != nil {
		errorSlice = append(errorSlice, "Empty image")
	}

	// If there are any errors, don't continue any further
	// Return to the user the errors
	if len(errorSlice) > 0 {
		response, err := json.Marshal(errorSlice)
		returnEmptyError(err)
		fmt.Fprint(w, string(response))
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

func steganoDecodeHandler(w http.ResponseWriter, r *http.Request) {

	session, err := store.Get(r, "user-login")
	returnEmptyError(err)
	// Check if session is set
	if len(session.Values) == 0 {
		http.Redirect(w, r, "/", http.StatusSeeOther) // Redirect to homepage
		return
	}

	// Check if image is uploaded successfully
	file, fileHeader, err := r.FormFile("image")
	if err != nil {
		return
	}

	fileExtension := strings.Split(fileHeader.Filename, ".")

	if fileExtension[len(fileExtension)-1] != "png" {
		return
	}

	buf := bytes.NewBuffer(nil)
	_, err = io.Copy(buf, file)
	if err != nil {
		return
	}

	// Decode the message
	message, err := decode(buf.Bytes())
	returnEmptyError(err)

	err = file.Close() // Close the file
	if err != nil {
		return
	}

	fmt.Fprint(w, message)
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
		return
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
		return
	}

	plaintext := r.FormValue("plaintext") // Get the text received
	shiftSize, err := strconv.Atoi(r.FormValue("shiftSize"))
	if err != nil {
		return
	}

	if shiftSize < 0 || shiftSize > 25 {
		return
	}

	ciphertext := encodeCaesar(plaintext, shiftSize)

	fmt.Fprint(w, ciphertext)

}

func caesarDecodeHandler(w http.ResponseWriter, r *http.Request) {

	// Check if session is set
	session, err := store.Get(r, "user-login")
	returnEmptyError(err)
	if len(session.Values) == 0 {
		http.Redirect(w, r, "/", http.StatusSeeOther) // Redirect to root
		return
	}

	// Check if textfield is empty
	if len(r.FormValue("ciphertext")) == 0 {
		fmt.Println("Empty text field")
		return
	}

	ciphertext := r.FormValue("ciphertext") // Get the text received
	shiftSize, err := strconv.Atoi(r.FormValue("shiftSizeD"))
	if err != nil {
		return
	}

	plaintext := decodeCaesar(ciphertext, shiftSize)

	fmt.Fprint(w, plaintext)

}

/* ROT 13 CIPHER */
func rot13GetHandler(w http.ResponseWriter, r *http.Request) {

	// Check if session is set
	session, err := store.Get(r, "user-login")
	returnEmptyError(err)
	if len(session.Values) == 0 {
		http.Redirect(w, r, "/", http.StatusSeeOther) // Redirect to root
		return
	}

	t, err := template.ParseFiles("assets/html/rot13.html")
	if err != nil {
		return
	}
	t.Execute(w, nil)
}

func rot13PostHandler(w http.ResponseWriter, r *http.Request) {

	// Check if session is set
	session, err := store.Get(r, "user-login")
	returnEmptyError(err)
	if len(session.Values) == 0 {
		http.Redirect(w, r, "/", http.StatusSeeOther) // Redirect to root
		return
	}

	// Check if textfield is empty
	if len(r.FormValue("plaintext")) == 0 {
		return
	}

	input := r.FormValue("plaintext") // Get the text received

	ciphertext := strings.Map(rot13, input)

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

func changePassPostHandler(w http.ResponseWriter, r *http.Request) {

	session, err := store.Get(r, "user-login")
	returnEmptyError(err)
	// Check if session is set
	if len(session.Values) == 0 {
		http.Redirect(w, r, "/", http.StatusSeeOther) // Redirect to homepage
		return
	}

	// Get the fields
	currentPass := r.FormValue("currentPass")
	newPass := r.FormValue("newPass")
	confirmPass := r.FormValue("confirmPass")

	sessionUser := session.Values["user"].(string)

	// Validate the fields
	errorSlice := validateChangePassword(currentPass, newPass, confirmPass, sessionUser)

	// Check if there are any errors
	if len(errorSlice) > 0 {
		response, err := json.Marshal(errorSlice)
		returnEmptyError(err)
		fmt.Fprint(w, string(response))
		return
	}

	// Hash the password
	hashPass, err := bcrypt.GenerateFromPassword([]byte(newPass), bcrypt.DefaultCost)
	returnEmptyError(err)

	// Change the password
	err = changeUserPassword(sessionUser, string(hashPass))
	returnEmptyError(err)

	fmt.Fprint(w, 1)

}

func deleteAccDeleteHandler(w http.ResponseWriter, r *http.Request) {

	session, err := store.Get(r, "user-login")
	returnEmptyError(err)
	// Check if session is set
	if len(session.Values) == 0 {
		http.Redirect(w, r, "/", http.StatusSeeOther) // Redirect to homepage
		return
	}

	// Get the fields
	sessionUser := session.Values["user"].(string)

	// Delete account
	err = deleteUser(sessionUser)
	returnEmptyError(err)

	// Remove session
	delete(session.Values, "user")
	session.Save(r, w)

	fmt.Fprint(w, 1)

}
