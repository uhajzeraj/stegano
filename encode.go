package main

import (
	"errors"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"io/ioutil"
	"os"
)

// Encoding function
func encode() {
	inputReader, inputErr := os.Open(*inputFilename) // read the input file
	panicOnError(inputErr)                           // panic on an error
	defer inputReader.Close()                        // close the reader

	message, inputMessageErr := ioutil.ReadFile(*messageFilename) // read the input message file
	panicOnError(inputMessageErr)                                 // panic on an erro

	img, _, imageDecodeErr := image.Decode(inputReader) // decode the image
	panicOnError(imageDecodeErr)                        // panic if image isn't decoded

	bounds := img.Bounds()                        // get the bounds of the image
	outputImage := image.NewNRGBA64(img.Bounds()) // create output image

	var messageIndex = 0 // get the rows and columns of the image

	fmt.Printf("The message is %d characters long\n", len(message))
	fmt.Println(bounds.Size().X * bounds.Size().Y)

	totalPixels := bounds.Size().X * bounds.Size().Y // Get the total number of pixels in the image

	if totalPixels < len(message) {
		fmt.Println("The text is larger than what can be hidden in the image")
		return
	}

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ { // loop over rows
		for x := bounds.Min.X; x < bounds.Max.X; x++ { // loop over columns

			r, g, b, a := img.At(x, y).RGBA() // get the rgba values from the input image

			if messageIndex < len(message) { // if we have bytes in message
				newr := uint32(message[messageIndex]>>6) + (r & lsbMask)            // first two bits (R)
				newg := uint32(message[messageIndex]>>4) & ^lsbMask + (g & lsbMask) // second two bits (G)
				newb := uint32(message[messageIndex]>>2) & ^lsbMask + (b & lsbMask) // third two bits (B)
				newa := uint32(message[messageIndex]) & ^lsbMask + (a & lsbMask)    // last two bits (A - alfa)
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

	outName := outputName(inputFilename)

	outputWriter, outputErr := os.Create(outName) // write the new file out
	panicOnError(outputErr)                       // Panic if there is an error
	defer outputWriter.Close()                    // Close output writer

	err := png.Encode(outputWriter, outputImage) // Png encode the writer
	if err != nil {
		panic(err)
	}
}
