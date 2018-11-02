package main

import (
	"strings"
)

// Damn you gometalinter and your cyclomatic complexity
func errorPanic(err error) {
	if err != nil {
		panic(err)
	}
}

// Output file name
func outputName(inputFN string) string {
	output := inputFN

	len := len("images/plain/") // Count the characters, and remove it using the `slice cutter`

	// Remove the file extension and add `_stegano.png` at the end
	output = "images/encoded/" + output[len:strings.LastIndex(output, ".")] + `_stegano.png`

	return output
}
