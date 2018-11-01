package main

import (
	"errors"
	"flag"
	"fmt"
	"image"
	"image/color"
	_ "image/jpeg"
	"image/png"
	"io/ioutil"
	"os"
)

// command line options
var inputFilename = flag.String("in", "", "input image file")
var outputFilename = flag.String("out", "", "output image file")
var messageFilename = flag.String("msg", "", "message input file")
var operation = flag.String("op", "encode", "encode or decode")

// example encode usage: go run png-lsb-steg.go -operation encode -image-input-file test.png -image-output-file steg.png -message-input-file hide.txt
// example decode usage: go run main.go -operation decode -image-input-file steg.png

// the bitmask we will use (last two bits)
var lsbMask = ^(uint32(3))

//helper to dry error handling up
func panicOnError(e error) {
	if e != nil {
		panic(e)
	}
}

// main, based on operation flag will encode data into image, or decode data from an image
func main() {

	// parse the command line options
	flag.Parse()

	switch *operation {
	case "encode":
		fmt.Println("encoding!")
		// read the input file
		inputReader, inputErr := os.Open(*inputFilename)
		// panic on an error
		panicOnError(inputErr)
		// close the reader
		defer inputReader.Close()

		// read the input message file
		message, inputMessageErr := ioutil.ReadFile(*messageFilename)
		// panic on an error
		panicOnError(inputMessageErr)

		// decode the image
		img, _, imageDecodeErr := image.Decode(inputReader)
		// panic if image isn't decoded
		panicOnError(imageDecodeErr)
		// get the bounds of the image
		bounds := img.Bounds()
		// create output image
		outputImage := image.NewNRGBA64(img.Bounds())
		// get the rows and columns of the image
		var messageIndex = 0
		// loop over rows
		for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
			// loop over columns
			for x := bounds.Min.X; x < bounds.Max.X; x++ {
				// get the rgba values from the input image
				r, g, b, a := img.At(x, y).RGBA()
				// if we have bytes in message
				if messageIndex < len(message) {
					// first two bits
					newr := uint32(message[messageIndex]>>6) + (r & lsbMask)
					// second two bits
					newg := uint32(message[messageIndex]>>4) & ^lsbMask + (g & lsbMask)
					// third two bits
					newb := uint32(message[messageIndex]>>2) & ^lsbMask + (b & lsbMask)
					// last two bits
					newa := uint32(message[messageIndex]) & ^lsbMask + (a & lsbMask)
					messageIndex++
					// set the color in the new output image
					outputImage.SetNRGBA64(x, y, color.NRGBA64{uint16(newr), uint16(newg), uint16(newb), uint16(newa)})
				} else if messageIndex == len(message) {
					// if we are done with our message bytes
					messageIndex++
					// set a null ascii char to know if we are done
					outputImage.SetNRGBA64(x, y, color.NRGBA64{uint16(0), uint16(0), uint16(0), uint16(0)})
				} else {
					// otherwise, just put the exact values in the new image
					outputImage.SetNRGBA64(x, y, color.NRGBA64{uint16(r), uint16(g), uint16(b), uint16(a)})
				}
			}
		}
		// we obviously have more data that won't fit in the image
		if messageIndex < len(message) {
			panicOnError(errors.New("out of space in input image"))
		}

		// write the new file out
		outputWriter, outputErr := os.Create(*outputFilename)
		// panic if fails..
		panicOnError(outputErr)
		// close output file when donw
		defer outputWriter.Close()
		// encode the png
		png.Encode(outputWriter, outputImage)
		// operation was decode that we passed in

	case "decode":
		fmt.Println("decoding!")
		// read the input file
		inputReader, inputErr := os.Open(*inputFilename)
		// panic on an error
		panicOnError(inputErr)
		// close the reader
		defer inputReader.Close()
		// decode the image
		img, _, imageDecodeErr := image.Decode(inputReader)
		// panic if image isn't decoded
		panicOnError(imageDecodeErr)
		// get the bounds of the image
		bounds := img.Bounds()
		// get the rows and columns of the image
		// loop over rows we will break here if done reading message
	OUTER:
		for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
			// loop over columns
			for x := bounds.Min.X; x < bounds.Max.X; x++ {
				// get the rgba values from the input image
				c := img.At(x, y).(color.NRGBA64)
				r := uint32(c.R)
				g := uint32(c.G)
				b := uint32(c.B)
				a := uint32(c.A)
				// build the byte from the color lsbs
				ch := (r & ^lsbMask) << 6
				ch += (g & ^lsbMask) << 4
				ch += (b & ^lsbMask) << 2
				ch += (a & ^lsbMask)
				// if we come across a zero byte
				if ch == 0 {
					break OUTER
				}
				// if the char is valid ascii print it out
				if ch >= 32 && ch <= 126 {
					fmt.Printf("%c", ch)
				}
			}
		}
		// newline
		fmt.Printf("\n")
	}
}
