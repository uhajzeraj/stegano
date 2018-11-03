package main

import (
	"flag"
	"fmt"
	_ "image/gif"
	_ "image/jpeg"
)

// example encode usage: go run png-lsb-steg.go -operation encode -image-input-file test.png -image-output-file steg.png -message-input-file hide.txt
// example decode usage: go run main.go -operation decode -image-input-file steg.png

// command line options
var inputFilename = flag.String("in", "", "input image file")
var messageFilename = flag.String("msg", "", "message input file")
var operation = flag.String("op", "encode", "encode or decode")

// the bitmask we will use (last two bits)
var lsbMask = ^(uint32(3))

// main, based on operation flag will encode data into image, or decode data from an image
func main() {

	// Connect to mongo before doing anything
	_, err := mongoConnect()
	if err != nil {
		panic(err)
	}

	// parse the command line options
	flag.Parse()

	switch *operation {
	case "encode":
		fmt.Println("encoding!")
		err := encode(inputFilename, messageFilename)
		errorPanic(err)

	case "decode":
		fmt.Println("decoding!")
		err := decode(inputFilename)
		errorPanic(err)
	}
}
