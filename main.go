package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"time"
)

// example encode usage: go run png-lsb-steg.go -operation encode -image-input-file test.png -image-output-file steg.png -message-input-file hide.txt
// example decode usage: go run main.go -operation decode -image-input-file steg.png

// command line options
var inputFilename = flag.String("in", "", "input image file")
var messageFilename = flag.String("msg", "", "message input file")
var operation = flag.String("op", "encode", "encode or decode")

func main() {

	// Create a router
	r := newRouter()

	srv := &http.Server{
		Handler: r,
		Addr:    ":" + os.Getenv("PORT"),
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
