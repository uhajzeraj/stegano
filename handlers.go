package main

import (
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
)

// Wrap mux router in a function for testing
func newRouter() *mux.Router {
	r := mux.NewRouter()
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

	t.Execute(w, nil)

	// fmt.Fprintf(w, "Hello world!")
}
