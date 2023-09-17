// This is the main package for the gobfs application.
//
// This package utilizes a SHA-1 hash algorithm and the custom hash utility from
// "github.com/ecorreiax/gobfs/internal/hash" to create a hash index.
// The hash index is then stored in a boolean bitset.
package main

import (
	"crypto/sha1"
	"fmt"

	"github.com/ecorreiax/gobfs/internal/hash"
	"github.com/ecorreiax/gobfs/internal/list"
)

// bitset is a boolean array used to store hash indexes.
var bitset []bool

func init() {
	var h = sha1.New()
	for _, s := range list.Invalid_usernames {
		idx := hash.CreateHash(h, s)
		if idx >= len(bitset) {
			newBitset := make([]bool, idx+1)
			copy(newBitset, bitset)
			bitset = newBitset
		}
		bitset[idx] = true
	}
}

// The main function of the gobfs application.
// It creates a hash index for a string and check if
// this index is already taken in the bitset.
func main() {
	var h = sha1.New()
	username := "ecorreia"
	idx := hash.CreateHash(h, username)
	p := bitset[idx]

	if p {
		fmt.Printf("username %v is not allowed\n", username)
		return
	}

	fmt.Printf("username %v is ok to be used\n", username)
}
