package main

import (
	"fmt"
	"image"
	"image/color"
	"os"
)

func decode(inputFile *string) error {
	inputReader, err := os.Open(*inputFile) // read the input file
	if err != nil {
		panic(err)
	}

	img, _, err := image.Decode(inputReader) // decode the image
	if err != nil {
		panic(err)
	}

	bounds := img.Bounds() // get the bounds of the image

	// get the rows and columns of the image
	// loop over rows we will break here if done reading message
	stringOutput := ""

LoopBreak:

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
				stringOutput += string(ch)
			}
		}
	}
	// newline
	fmt.Printf("\n")
	err = inputReader.Close() // close the reader
	if err != nil {
		panic(err)
	}

	saveFile(outputName(*inputFile), stringOutput)
	return nil
}
