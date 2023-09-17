// Package hash offers utilities to generate hash values based on provided inputs.
//
// This package utilizes the standard library's hash interface and custom parser utilities
// from the "github.com/ecorreiax/gobfs/internal/parser" package.
package hash

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"hash"

	"github.com/ecorreiax/gobfs/internal/parser"
)

// CreateHash computes a hash index based on a given string and hash algorithm.
// It uses the parser from "github.com/ecorreiax/gobfs/internal/parser" to sanitize the input
// before hashing it and then calculates an integer index from the hashed bytes.
//
// Example:
//
//	h := sha256.New()
//	input := "some random string"
//	idx := CreateHash(h, input)
//
// Parameters:
//
//	h: A hash.Hash compliant hashing algorithm.
//	s: The input string to hash.
//
// Returns:
//
//	An integer hash index generated from the hashed bytes.
func CreateHash(h hash.Hash, s string) int {
	h.Write([]byte(parser.ParseString(s)))
	bits := h.Sum(nil)
	buf := bytes.NewBuffer(bits)
	result, _ := binary.ReadVarint(buf)
	idx := int(result)

	fmt.Printf("%v:%v\n", "idx", idx)

	return idx
}
