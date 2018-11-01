package main

import (
	"strings"
)

// Helper for error handling
func panicOnError(e error) {
	if e != nil {
		panic(e)
	}
}

// Output file name
func outputName(inputFN *string) string {
	output := *inputFN

	len := len("images/plain/") // Count the characters, and remove it using the `slice cutter`

	// Remove the file extension and add `_stegano.png` at the end
	output = "images/encoded/" + output[len:strings.LastIndex(output, ".")] + `_stegano.png`

	return output
}
