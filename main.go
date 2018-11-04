package main

import (
	"flag"
	"fmt"
	"image"
	"image/png"
	"io"
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

/*
func handler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		t, err := template.ParseFiles("web-dev/front-end/stegano.html")
		if err != nil {
			http.Error(w, "Error processing template", http.StatusInternalServerError)
		}
		t.Execute(w, nil)

	case http.MethodPost:
		url := "http://i.imgur.com/m1UIjW1.png"

		response, e := http.Get(url)
		if e != nil {
			log.Fatal(e)
		}

		defer response.Body.Close()

		//open a file for writing
		file, err := os.Create("/tmp/asdf.png")
		if err != nil {
			log.Fatal(err)
		}
		// Use io.Copy to just dump the response body to the file. This supports huge files
		_, err = io.Copy(file, response.Body)
		if err != nil {
			log.Fatal(err)
		}
		file.Close()
		//fmt.Println("Success!")
		// Read image from file that already exists
		existingImageFile, err := os.Open("/tmp/asdf.png")
		if err != nil {
			// Handle error
			log.Fatal(err)
		}
		defer existingImageFile.Close()

		// Calling the generic image.Decode() will tell give us the data
		// and type of image it is as a string. We expect "png"
		imageData, imageType, err := image.Decode(existingImageFile)
		if err != nil {
			// Handle error
			log.Fatal(err)
		}
		fmt.Println(imageData)
		fmt.Println(imageType)

		// We only need this because we already read from the file
		// We have to reset the file pointer back to beginning
		existingImageFile.Seek(0, 0)

		// Alternatively, since we know it is a png already
		// we can call png.Decode() directly
		loadedImage, err := png.Decode(existingImageFile)
		if err != nil {
			// Handle error
			log.Fatal(err)
		}
		fmt.Println(loadedImage)
	default:
		http.Error(w, "Not implemented", http.StatusNotImplemented)
	}

}
*/
