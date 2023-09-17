// This is the main package for the gobfs application.
//
// This package utilizes a SHA-1 hash algorithm and the custom hash utility from
// "github.com/ecorreiax/gobfs/internal/hash" to create a hash index.
// The hash index is then stored in a boolean bitset.
package gobfs

import (
	"crypto/sha1"

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

// Check validates a given username against a predefined dataset
// and returns a boolean value indicating its status.
//
// This function employs SHA-1 hashing along with a bitset for efficient
// lookups. If the username exists in the dataset, it returns true, signaling
// that the username should not be allowed in your database.
//
// Parameters:
//
//	u string: The username to be checked.
//
// Returns:
//
//	bool: True if username exists in dataset, otherwise false.
//
// Example:
//
//	result := Check("someUsername")
func Check(u string) bool {
	var h = sha1.New()
	idx := hash.CreateHash(h, u)
	return bitset[idx]
}
