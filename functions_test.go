package main

import (
	"io/ioutil"
	"os"
	"testing"
)

var plainText = "helloooo"
var cipherTxt = encodeCaesar(plainText, 3)

func Test_encodeCaeser(t *testing.T) {
	err := encodeCaesar(plainText, 3)
	if err == "" {
		t.Errorf("Error during Encode proces with Ceaser, %s", err)
	}
}

func Test_decodeCaeser(t *testing.T) {
	err := decodeCaesar(plainText, 3)
	if err == "" {
		t.Errorf("Error during Decode proces with Ceaser, %s", err)
	}

}
func Test_errorPanic(t *testing.T) {
	var err error
	if err != nil {
		panic(err)
	}

}
func Test_exists(t *testing.T) {
	boolValue, err := exists("./tmp/asdf.png")
	if err != nil {
		t.Error("Error:Doesn't exist")
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
func Test_validateSignup2(t *testing.T) {

	err := validateSignup("uhajzeraj", "etnik123", "etnik1s3", "besnikk@stud.ntnu.no")
	if err == nil {
		t.Errorf("Error validating Signup:%s", err)
	}
}

func Test_validateLogin(t *testing.T) {

	err := validateLogin("etnik5", "etnik123")
	if err == nil {
		t.Errorf("Error validating Login:%s", err)
	}
}

func Test_validateLogin2(t *testing.T) {

	err := validateLogin("uhajzeraj", "urani123")
	if err != nil {
		t.Errorf("Error validating Login:%s", err)
	}
}
func Test_validateChangePassword(t *testing.T) {
	err := validateChangePassword("urani123", "urani123", "urani123", "uhajzeraj")
	if err != nil {
		t.Errorf("Error during Change Password, %s", err)
	}
}

func Test_cipherEncoding(t *testing.T) {
	err := cipher(plainText, -1, 3)
	if err == "" {
		t.Error("Error: Not encode!")
	}
	// if err != cipherTxt {
	// 	t.Error("Encoding not correct")
	// }
}
func Test_cipherDecoding(t *testing.T) {

	err := cipher(cipherTxt, +1, 3)
	if err == "" {
		t.Errorf("Error: Not encode!, %s", err)
	}
	// if err != plainText {
	// 	t.Error("Decoding not correct")
	// }
}

func Test_encode(t *testing.T) {
	_, err := ioutil.ReadFile("tmp/asdf.png")
	if err != nil {
		panic(err)
	}
	file, err := os.Open("tmp/asdf.png")
	if err != nil {
		panic(err)
	}
	_, err = encode(file, "hello")
	if err != nil {
		t.Error("Encode function doesn't work")
	}
}

func Test_decode(t *testing.T) {

	_, err := ioutil.ReadFile("tmp/asdf.png")
	if err != nil {
		panic(err)
	}
	file, err := os.Open("tmp/asdf.png")
	if err != nil {
		panic(err)
	}
	bytes, er := encode(file, "hello")
	if er != nil {
		t.Error("Encode function doesn't work")
	}

	_, err = decode(bytes)
	if err != nil {
		t.Error("Decode function doesn't work")
	}
}

func Test_rot13(t *testing.T) {
	r := rot13(45)

	if r == 0 {
		t.Error("rot13 isn't working properly!")
	}
}

func Test_encoding(t *testing.T) {
	var j rune = 1
	i := encoding(&j, 2, 3)
	if i == 0 {
		t.Error("Encoding isn't working properly!")
	}
}

func Test_decoding(t *testing.T) {
	var j rune = 1
	i := decoding(&j, 2, 3)
	if i == 0 {
		t.Error("Encoding isn't working properly!")
	}
}
