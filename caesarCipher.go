package main

// cipher takes in the text to be ciphered along with the direction that
// is being taken; -1 means decoding, +1 means encoding.
func cipher(text string, direction int, shiftSize int) string {
	// shift -> number of letters to move to right or left
	// offset -> size of the alphabet, in this case the plain ASCII
	shift, offset := rune(shiftSize), rune(26)

	// string->rune conversion
	runes := []rune(text)

	for index, char := range runes {
		// Iterate over all runes, and perform substitution
		// wherever possible. If the letter is not in the range
		// [1 .. 25], the offset defined above is added or
		// subtracted.
		switch direction {
		case -1: // decoding
			decoding(&char, shift, offset)
		case +1: // encoding
			encoding(&char, shift, offset)
		}

		// Above `if`s handle both upper and lower case ASCII
		// characters; anything else is returned as is (includes
		// numbers, punctuation and space).
		runes[index] = char
	}

	return string(runes)
}

func decoding(char *rune, shift rune, offset rune) rune {

	if *char >= 'a'+shift && *char <= 'z' ||
		*char >= 'A'+shift && *char <= 'Z' {
		*char = *char - shift
	} else if *char >= 'a' && *char < 'a'+shift ||
		*char >= 'A' && *char < 'A'+shift {
		*char = *char - shift + offset
	}

	return *char

}

func encoding(char *rune, shift rune, offset rune) rune {

	if *char >= 'a' && *char <= 'z'-shift ||
		*char >= 'A' && *char <= 'Z'-shift {
		*char = *char + shift
	} else if *char > 'z'-shift && *char <= 'z' ||
		*char > 'Z'-shift && *char <= 'Z' {
		*char = *char + shift - offset
	}

	return *char

}

// encode and decode provide the API for encoding and decoding text using
// the Caesar Cipher algorithm.
func encodeCaesar(text string, shiftSize int) string { return cipher(text, +1, shiftSize) }
func decodeCaesar(text string, shiftSize int) string { return cipher(text, -1, shiftSize) }
