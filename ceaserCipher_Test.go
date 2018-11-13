package main

import "testing"

var plainText = "helloooo"
var cipherTxt = encodeCaesar(plainText, 3)

func Test_encodeCaeser(t *testing.T) {
	err := encodeCaesar(plainText, 3)
	if err != "" {
		t.Errorf("Error during Encode proces with Ceaser, %s", err)
	}
}

func Test_decodeCaeser(t *testing.T) {
	err := decodeCaesar(plainText, 3)
	if err != "" {
		t.Errorf("Error during Decode proces with Ceaser, %s", err)
	}
	if err != plainText {
		t.Error("Error decoding")
	}
}

// func Test_cipherEncoding(t *testing.T) {
// 	err := cipher(plainText, -1, 3)
// 	if err != "" {
// 		t.Error("Error: Not encode!")
// 	}
// 	if err != cipherTxt {
// 		t.Error("Encoding not correct")
// 	}
// }
// func Test_cipherDecoding(t *testing.T) {
// 	err := cipher(cipherTxt, +1, 3)
// 	if err != "" {
// 		t.Errorf("Error: Not encode!,%s", err)
// 	}
// 	if err != plainText {
// 		t.Error("Decoding not correct")
// 	}
// }
