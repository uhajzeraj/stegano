package main

func rot13(r rune) rune {

	if r >= 'a' && r <= 'z' {
		// Rotate lowercase letters 13 places.
		if r >= 'm' {
			return r - 13
		}
		return r + 13
	} else if r >= 'A' && r <= 'Z' {
		// Rotate uppercase letters 13 places.
		if r >= 'M' {
			return r - 13
		}
		return r + 13
	}
	// Do nothing.
	return r
}
