package main

import (
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"net/smtp"
	"regexp"

	"github.com/gorilla/mux"
)

// example encode usage: go run png-lsb-steg.go -operation encode -image-input-file test.png -image-output-file steg.png -message-input-file hide.txt
// example decode usage: go run main.go -operation decode -image-input-file steg.png

// command line options
var inputFilename = flag.String("in", "", "input image file")
var messageFilename = flag.String("msg", "", "message input file")
var operation = flag.String("op", "encode", "encode or decode")

type User struct {
	Username string
	Password string
	Errors   map[string]string
}

//Validate is used
func (usr *User) Validate() bool {
	usr.Errors = make(map[string]string)

	re := regexp.MustCompile(".+@.+\\..+")
	matched := re.Match([]byte(usr.Username))
	if matched == false {
		usr.Errors["Username"] = "Please enter a valid email address"
	}
	// re1 := regexp.MustCompile("^(?=.*[A-Z].*[A-Z])(?=.*[!@#$&*])(?=.*[0-9].*[0-9])(?=.*[a-z].*[a-z].*[a-z]).{8}")
	// matched1 := re1.Match([]byte(usr.Password))
	if usr.Password == "" {
		usr.Errors["Password"] = "Password is not strong"
	}

	return len(usr.Errors) == 0
}

func (usr *User) Deliver() error {
	to := []string{"kelmendi.besnik3@gmail.com"}
	body := fmt.Sprintf("Reply-To: %v\r\nSubject: New Message\r\n%v", usr.Username, "Someone is trying to get into your account!")

	username := "nickmkelmendi@gmail.com"
	password := "TestTest123"
	auth := smtp.PlainAuth("", username, password, "smtp.gmail.com")

	return smtp.SendMail("smtp.gmail.com:587", auth, usr.Username, to, []byte(body))
}

func main() {

	// // Connect to mongo before doing anything
	// client, err := mongoConnect()
	// if err != nil {
	// 	panic(err)
	// }

	// img, err := getImage(client)
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println(img["imgEncoding"])

	// path := "assets/images/plain/rick_morty.jpg"

	// encoding, err := image64Encode(path)

	// if err != nil {
	// 	panic(err)
	// }

	// storeImage(client, encoding)

	// // parse the command line options
	// flag.Parse()

	// switch *operation {
	// case "encode":
	// 	fmt.Println("encoding!")
	// 	err := encode(inputFilename, messageFilename)
	// 	errorPanic(err)

	// case "decode":
	// 	fmt.Println("decoding!")
	// 	err := decode(inputFilename)
	// 	errorPanic(err)
	// }

	// Create a router
	r := mux.NewRouter()

	// muxs := pat.New()
	// muxs.Get("/", http.HandlerFunc(index))
	// muxs.Get("/login.html", http.HandlerFunc(login))
	// muxs.Get("/signup.html", http.HandlerFunc(signup))
	// muxs.Post("/login.html", http.HandlerFunc(signin))

	r.HandleFunc("/", index).Methods("GET")
	r.HandleFunc("/login.html", login).Methods("GET")
	r.HandleFunc("/signup.html", signup).Methods("GET")
	r.HandleFunc("/login.html", signin).Methods("POST")
	r.HandleFunc("/signup.html", signup).Methods("GET")
	r.HandleFunc("/home.html", home).Methods("GET")

	// srv := &http.Server{
	// 	Handler: r,
	// 	Addr:    "127.0.0.1:8080",
	// 	// Good practice: enforce timeouts for servers you create!
	// 	WriteTimeout: 15 * time.Second,
	// 	ReadTimeout:  15 * time.Second,
	// }
	// log.Fatal(srv.ListenAndServe())

	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}

}

func index(w http.ResponseWriter, r *http.Request) {
	render(w, "front-end/index.html", nil)
}
func login(w http.ResponseWriter, r *http.Request) {
	render(w, "front-end/login.html", nil)
}
func signup(w http.ResponseWriter, r *http.Request) {
	render(w, "front-end/signup.html", nil)
}

func home(w http.ResponseWriter, r *http.Request) {
	render(w, "front-end/home.html", nil)
}

func signin(w http.ResponseWriter, r *http.Request) {
	usr := &User{
		Username: r.FormValue("username"),
		Password: r.FormValue("password"),
	}
	if usr.Validate() == false {
		render(w, "front-end/login.html", nil)
		if err := usr.Deliver(); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		return
	}

	http.Redirect(w, r, "/home.html", http.StatusSeeOther)
}
func render(w http.ResponseWriter, filename string, data interface{}) {
	tmpl, err := template.ParseFiles(filename)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	if err := tmpl.Execute(w, data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
