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

// Helper for error handling
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

		inputReader, inputErr := os.Open(*inputFilename) // read the input file
		panicOnError(inputErr)                           // panic on an error
		defer inputReader.Close()                        // close the reader

		message, inputMessageErr := ioutil.ReadFile(*messageFilename) // read the input message file
		panicOnError(inputMessageErr)                                 // panic on an erro

		img, _, imageDecodeErr := image.Decode(inputReader) // decode the image
		panicOnError(imageDecodeErr)                        // panic if image isn't decoded

		bounds := img.Bounds() // get the bounds of the image

		outputImage := image.NewNRGBA64(img.Bounds()) // create output image

		var messageIndex = 0 // get the rows and columns of the image

		for y := bounds.Min.Y; y < bounds.Max.Y; y++ { // loop over rows
			for x := bounds.Min.X; x < bounds.Max.X; x++ { // loop over columns

				r, g, b, a := img.At(x, y).RGBA() // get the rgba values from the input image

				if messageIndex < len(message) { // if we have bytes in message

					newr := uint32(message[messageIndex]>>6) + (r & lsbMask) // first two bits (R)

					newg := uint32(message[messageIndex]>>4) & ^lsbMask + (g & lsbMask) // second two bits (G)

					newb := uint32(message[messageIndex]>>2) & ^lsbMask + (b & lsbMask) // third two bits (B)

					newa := uint32(message[messageIndex]) & ^lsbMask + (a & lsbMask) // last two bits (A - alfa)
					messageIndex++

					outputImage.SetNRGBA64(x, y, color.NRGBA64{uint16(newr), uint16(newg), uint16(newb), uint16(newa)}) // set the color in the new output image
				} else if messageIndex == len(message) {
					// if we are done with our message bytes
					messageIndex++

					outputImage.SetNRGBA64(x, y, color.NRGBA64{uint16(0), uint16(0), uint16(0), uint16(0)}) // set a null ascii char to know if we are done
				} else {
					outputImage.SetNRGBA64(x, y, color.NRGBA64{uint16(r), uint16(g), uint16(b), uint16(a)}) // otherwise, just put the exact values in the new image
				}
			}
		}

		if messageIndex < len(message) { // We have more data then what can fit the image
			panicOnError(errors.New("out of space in input image"))
		}

		outputWriter, outputErr := os.Create(*outputFilename) // write the new file out

		panicOnError(outputErr) // Panic if there is an error

		defer outputWriter.Close() // Close output writer

		png.Encode(outputWriter, outputImage) // Png encode the writer

	case "decode":
		fmt.Println("decoding!")

		inputReader, inputErr := os.Open(*inputFilename) // read the input file
		panicOnError(inputErr)                           // panic on an error
		defer inputReader.Close()                        // close the reader

		img, _, imageDecodeErr := image.Decode(inputReader) // decode the image
		panicOnError(imageDecodeErr)                        // panic if image isn't decoded

		bounds := img.Bounds() // get the bounds of the image

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

				// If the char is valid ascii print it out
				if (ch >= 32 && ch <= 126) || ch == '\n' {
					fmt.Printf("%c", ch)
				}
			}
		}
		// newline
		fmt.Printf("\n")
	}
}
