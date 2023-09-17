// This is the main package for the gobfs application.
//
// This package utilizes a SHA-1 hash algorithm and the custom hash utility from
// "github.com/ecorreiax/gobfs/internal/hash" to create a hash index.
// The hash index is then stored in a boolean bitset.
package main

import (
	"crypto/sha1"

	"github.com/ecorreiax/gobfs/cmd"
	"github.com/ecorreiax/gobfs/internal/forbidden"
	"github.com/ecorreiax/gobfs/internal/hash"
)

// init initializes the global bitset with precomputed hash values
// from a list of invalid usernames.
//
// This function is automatically called when the package is imported.
// It uses SHA-1 hashing via the CreateHash function to generate indexes
// for each invalid username. These indexes are then added to the bitset
// using the AddHash function.
//
// Note:
//
//	This function depends on 'list.Invalid_usernames' to be populated with
//	the set of usernames that are to be marked as invalid.
//
// Example:
//
//	Automatically invoked upon package import.
func init() {
	h := sha1.New()
	for _, s := range forbidden.ForbiddenUsernames {
		idx := hash.CreateHash(h, s)
		hash.AddHash(idx)
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
	h := sha1.New()
	idx := hash.CreateHash(h, u)
	return hash.VerifyHash(idx)
}

// main is the entry point for the gobfs application.
//
// This function simply invokes the Execute function from the cmd package, which initializes
// and runs the root Cobra command. Any errors returned from Execute are logged to the standard output.
//
// Example:
//
//	Invoked automatically when the application is run.
func main() {
	cmd.Execute()
}
