package main

import (
	"testing"
)

func Test_outputName(t *testing.T) {

	// Some dummy names
	inputNames := []string{"test.png",
		"rick.and.morty.png",
		"rias.//.213.#@!3/k21.3123k1o3(*$.jpeg"}

	// Some dummy names to test against
	outputNames := []string{"test_stegano.png",
		"rick.and.morty_stegano.png",
		"rias.//.213.#@!3/k21.3123k1o3(*$_stegano.png"}

	for i := 0; i < len(inputNames); i++ {
		if outputName("images/plain/"+inputNames[i]) != "images/encoded/"+outputNames[i] {
			t.Errorf("Bad output name, wanted %s, got %s", outputName(inputNames[i]), "images/encoded/"+outputNames[i])
		}
	}

}

func Test_errorPanic(t *testing.T) {
	var err error
	errorPanic(err)
}
func Test_saveile(t *testing.T){

}