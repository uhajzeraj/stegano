package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {

	key, _ := getImages("username")
	fmt.Println(key)

	// Create a router
	r := newRouter()

	srv := &http.Server{
		Handler: r,
		Addr:    ":8080",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
