// This is the main package for the gobfs application.
//
// This package utilizes a SHA-1 hash algorithm and the custom hash utility from
// "github.com/ecorreiax/gobfs/internal/hash" to create a hash index.
// The hash index is then stored in a boolean bitset.
package main

import (
	"crypto/sha1"

	"github.com/ecorreiax/gobfs/internal/hash"
)

// h is the hash function being used, in this case, SHA-1.
var h = sha1.New()

// bitset is a boolean array used to store hash indexes.
var bitset [100]bool

// The main function of the gobfs application.
// It creates a hash index for a string and stores it in the bitset.
func main() {
	h1 := hash.CreateHash(h, "e_correia") // TODO: remove hard coded username
	bitset[h1] = true
}
