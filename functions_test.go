package main

import (
	"testing"
)

func Test_errorPanic(t *testing.T) {
	var err error
	errorPanic(err)
}
func Test_exists(t *testing.T) {
	boolValue, err := exists("../stegano/.idea")
	if err != nil {
		t.Error("Error:Doesnt exist")
	}
	if boolValue != true {
		t.Error("Path dont exist")
	}
}

func Test_validateSignup(t *testing.T) {

	err := validateSignup("besniku7", "etnik123", "etnik123", "besnikk@stud.ntnu.no")
	if err != nil {
		t.Errorf("Error validating Signup:%s", err)
	}
}

func Test_validateLogin(t *testing.T) {

	err := validateLogin("etnik5", "etnik123")
	if err == nil {
		t.Errorf("Error validating Login:%s", err)
	}
}

func Test_encodeCaeser(t *testing.T) {
	err := encodeCaesar(plainText, 3)
	if err == "" {
		t.Errorf("Error during Encode proces with Ceaser, %s", err)
	}

}

func Test_decodeCaeser(t *testing.T) {
	err := decodeCaesar(cipherTxt, 3)
	if err == "" {
		t.Errorf("Error during Decode proces with Ceaser, %s", err)
	}
	if err != plainText {
		t.Error("Error decoding")
	}
}

func Test_cipherEncoding(t *testing.T) {
	err := cipher(plainText, -1, 3)
	if err == "" {
		t.Error("Error: Not encode!")
	}
	if err != cipherTxt {
		t.Error("Encoding not correct")
	}
}
func Test_cipherDecoding(t *testing.T) {

	err := cipher(cipherTxt, +1, 3)
	if err == "" {
		t.Errorf("Error: Not encode!, %s", err)
	}
	if err != plainText {
		t.Error("Decoding not correct")
	}
}
