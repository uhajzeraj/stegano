package main

import (
	"fmt"
	"log"
	"os"
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

// Save the decoded test
func saveFile(fileName string, stringOutput string) {
	file, err := os.Create(fileName)
	if err != nil {
		log.Fatal("Cannot create file", err)
	}
	defer file.Close()

	fmt.Fprintf(file, stringOutput)
}
