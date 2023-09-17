// Package parser provides utilities for parsing strings based on specific conditions.
//
// This package depends on the "github.com/rivo/uniseg" library to handle Unicode grapheme clusters.
package parser

import (
	"strings"
	"unicode"

	"github.com/rivo/uniseg"
)

// ParseString takes an input string and returns a new string
// containing only letters and digits from the original.
//
// This function employs the github.com/rivo/uniseg package for handling Unicode grapheme clusters.
//
// Example:
//
//	input := "Hello, Wörld!"
//	output := ParseString(input)
//	// Output: "HelloWörld"
//
// Parameters:
//
//	s: The input string to be parsed.
//
// Returns:
//
//	A new string containing only letters and digits.
func ParseString(s string) string {
	var b strings.Builder
	gr := uniseg.NewGraphemes(s)
	for gr.Next() {
		r := gr.Runes()[0]
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			b.WriteString(gr.Str())
		}
	}

	return b.String()
}
