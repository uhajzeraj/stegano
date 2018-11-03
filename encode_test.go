package main

import (
	"testing"
)

func Test_encode(t *testing.T) {

	inputFileString := "images/plain/mario.png"
	messageFileString := "text/Harry Potter.txt"

	inputFile := &inputFileString
	messageFile := &messageFileString

	err := encode(inputFile, messageFile)

	if err != nil {
		t.Error("Something went wrong", err)
	}
}
