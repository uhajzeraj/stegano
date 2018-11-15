package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/context"
)

func main() {

	// Create a router
	r := newRouter()

	srv := &http.Server{
		Handler: context.ClearHandler(r),
		Addr:    ":" + os.Getenv("PORT"),
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
