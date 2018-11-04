package main

import (
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
)

// Test struct for testing
type Test struct {
	Title     string
	ImgEncode string
}

// Wrap mux router in a function for testing
func newRouter() *mux.Router {
	r := mux.NewRouter()

	// Diferent path - method handlers
	r.HandleFunc("/", rootHandler).Methods("GET")

	// Static file directory
	staticFileDirectory := http.Dir("./assets/")
	staticFileHadler := http.StripPrefix("/assets/", http.FileServer(staticFileDirectory))
	r.PathPrefix("/assets/").Handler(staticFileHadler).Methods("GET")

	return r
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	// w.Header().Set("Content-Type", "text/html; charset=utf-8")

	t, err := template.ParseFiles("assets/html/index.html")
	if err != nil {
		panic(err)
	}

	// Connect to mongo before doing anything
	client, err := mongoConnect()
	if err != nil {
		panic(err)
	}

	_, err = getImage(client)
	if err != nil {
		panic(err)
	}

	t.Execute(w, Test{"Best page title", `assets/images/plain/smaller_image.jpeg`})

	// fmt.Fprintf(w, "Hello world!")
}
