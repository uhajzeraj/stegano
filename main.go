package main

import (
	"bytes"
	"flag"
	"html/template"
	_ "image/gif"
	_ "image/jpeg"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

// example encode usage: go run png-lsb-steg.go -operation encode -image-input-file test.png -image-output-file steg.png -message-input-file hide.txt
// example decode usage: go run main.go -operation decode -image-input-file steg.png

// command line options
var inputFilename = flag.String("in", "", "input image file")
var messageFilename = flag.String("msg", "", "message input file")
var operation = flag.String("op", "encode", "encode or decode")

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
		image := r.PostFormValue("file")
		plText := r.FormValue("text")
		r := bytes.NewReader([]byte(plText))
		//using goquery to get value from plaintext element
		doc, _ := goquery.NewDocumentFromReader(r)
		text := doc.Find("textarea").Text()
		encode(&image, &text)
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
}
