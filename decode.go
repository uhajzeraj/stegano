package main

import (
	"fmt"
	"image"
	"image/color"
	"os"
)

func decode() {
	inputReader, inputErr := os.Open(*inputFilename) // read the input file
	panicOnError(inputErr)                           // panic on an error
	defer inputReader.Close()                        // close the reader

	img, _, imageDecodeErr := image.Decode(inputReader) // decode the image
	panicOnError(imageDecodeErr)                        // panic if image isn't decoded

	bounds := img.Bounds() // get the bounds of the image

	// get the rows and columns of the image
	// loop over rows we will break here if done reading message
LoopBreak:
	// fmt.Println("Is it coming here")
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
				break LoopBreak
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
