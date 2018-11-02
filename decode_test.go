package main

import (
	"testing"
)

func Test_decode(t *testing.T) {

	inputFileString := "images/encoded/mario_stegano.png"

	inputFile := &inputFileString

	decode(inputFile)
}
