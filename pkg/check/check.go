package check

import (
	"crypto/sha1"

	"github.com/ecorreiax/gobfs/internal/hash"
)

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
	h := sha1.New()
	idx := hash.CreateHash(h, u)
	return hash.VerifyHash(idx)
}
