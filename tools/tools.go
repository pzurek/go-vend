package tools

import "strings"

// Returns the original string s padded on the left with a given string p to a total length of l
func PadLeft(s, p string, l int) string {

	// If the length of s is already longer of equal the required length just return s
	if l < len(s) {
		return s
	}

	pad := createPadding(s, p, l)
	return pad + s
}

// Returns the original string s padded on the right with a given string p to a total length of l
func PadRight(s, p string, l int) string {

	// If the length of s is already longer of equal the required length just return s
	if l < len(s) {
		return s
	}

	pad := createPadding(s, p, l)
	return s + pad
}

func createPadding(s, p string, l int) string {
	// How many additional characters do we need?
	delta := l - len(s)

	// How many times do we need to repeat the padding string to get it?
	count := (delta / len(p)) + 1

	pad := strings.Repeat(p, count)

	// Grab just as many characters as we need
	return pad[0:delta]
}
