package main

import (
	"flag"
	"log"
	"net/http"
	"time"
)

// example encode usage: go run png-lsb-steg.go -operation encode -image-input-file test.png -image-output-file steg.png -message-input-file hide.txt
// example decode usage: go run main.go -operation decode -image-input-file steg.png

// command line options
var inputFilename = flag.String("in", "", "input image file")
var messageFilename = flag.String("msg", "", "message input file")
var operation = flag.String("op", "encode", "encode or decode")

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
	r := newRouter()

	srv := &http.Server{
		Handler: r,
		Addr:    "127.0.0.1:8080",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
