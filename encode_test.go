package main

import (
	"testing"
)

func Test_encode(t *testing.T) {

	inputFileString := "images/plain/mario.png"
	messageFileString := "text/Harry Potter.txt"

	inputFile := &inputFileString
	messageFile := &messageFileString

	encode(inputFile, messageFile)
}
