// Package hash offers utilities to generate hash values based on provided inputs.
//
// This package utilizes the standard library's hash interface and custom parser utilities
// from the "github.com/ecorreiax/gobfs/internal/parser" package.
package hash

import (
	"bytes"
	"encoding/binary"
	"hash"

	"github.com/ecorreiax/gobfs/internal/parser"
)

// bitset is a boolean array used to store hash indexes.
var bitset []bool

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
	return int(result)
}

// AddHash sets a bit in the global bitset based on the given index.
//
// This function first ensures that the index is non-negative. It then checks
// if the index falls outside the bounds of the existing bitset. If it does,
// the bitset is resized to accommodate the new index.
//
// Parameters:
//
//	idx int: The index at which the bit should be set. Negative values are converted to positive.
//
// Example:
//
//	AddHash(42)
func AddHash(idx int) {
	if idx < 0 {
		idx = -idx
	}

	if idx >= len(bitset) {
		newBitset := make([]bool, idx+1)
		copy(newBitset, bitset)
		bitset = newBitset
	}

	bitset[idx] = true
}

// VerifyHash checks if a bit is set at a given index in the global bitset.
//
// The function returns true if the bit at the specified index is set,
// otherwise false. It is the caller's responsibility to ensure that the
// index is within the bounds of the bitset.
//
// Parameters:
//
//	idx int: The index of the bit to be verified.
//
// Returns:
//
//	bool: True if the bit at the given index is set, otherwise false.
//
// Example:
//
//	result := VerifyHash(42)
func VerifyHash(idx int) bool {
	return bitset[idx]
}
