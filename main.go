package main

import (
	"flag"
<<<<<<< HEAD
	"fmt"
	"html/template"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	"image/png"
	"io"
	"log"
	"net/http"
	"os"
=======
	"log"
	"net/http"
	"os"
	"time"
>>>>>>> 2f9be89f9667689eeceeb83ad34e4c031d5091ce
)

// example encode usage: go run png-lsb-steg.go -operation encode -image-input-file test.png -image-output-file steg.png -message-input-file hide.txt
// example decode usage: go run main.go -operation decode -image-input-file steg.png

// command line options
var inputFilename = flag.String("in", "", "input image file")
var messageFilename = flag.String("msg", "", "message input file")
var operation = flag.String("op", "encode", "encode or decode")

<<<<<<< HEAD
// the bitmask we will use (last two bits)
var lsbMask = ^(uint32(3))

func handler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		t, err := template.ParseFiles("front-end/stegano.html")
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
		file, err := os.Create("C:/Users/Etnik Gashi/Desktop/stegano1/tmp/asdf.png")
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
		existingImageFile, err := os.Open("C:/Users/Etnik Gashi/Desktop/stegano1/tmp/asdf.png")
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

// main, based on operation flag will encode data into image, or decode data from an image
func main() {

	// parse the command line options
	// flag.Parse()

	// switch *operation {
	// case "encode":
	// 	fmt.Println("encoding!")
	// 	encode(inputFilename, messageFilename)

	// case "decode":
	// 	fmt.Println("decoding!")
	// 	decode(inputFilename)
	// }

	http.HandleFunc("/submit", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
=======
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
>>>>>>> 2f9be89f9667689eeceeb83ad34e4c031d5091ce
}
