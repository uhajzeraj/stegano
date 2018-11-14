package main

import (
	"bytes"
	"errors"
	"image"
	"image/color"
	_ "image/gif"
	_ "image/jpeg"
	"image/png"
	"mime/multipart"
)

// The bitmask that will be used (last two bits)
var lsbMask = ^(uint32(3))

func encode(inputReader multipart.File, inputMessage string) ([]byte, error) {

	message := []byte(inputMessage)

	img, _, err := image.Decode(inputReader) // decode the image
	if err != nil {
		return nil, err
	}

	err = inputReader.Close() // close the reader
	if err != nil {
		return nil, err
	}

	bounds := img.Bounds()                        // get the bounds of the image
	outputImage := image.NewNRGBA64(img.Bounds()) // create output image

	var messageIndex = 0 // get the rows and columns of the image

	totalPixels := bounds.Size().X * bounds.Size().Y // Get the total number of pixels in the image

	if totalPixels < len(message) {
		return nil, errors.New("The text is larger than what can be hidden in the image")
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

	buf := new(bytes.Buffer)
	err = png.Encode(buf, outputImage)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}
