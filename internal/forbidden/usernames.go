// Package forbidden contains a pre-defined list of usernames to be considered invalid.
//
// This list is used to pre-populate the bloom filter for rapid username checking.
//
// Variables:
//   - Invalid_usernames: An array of strings containing usernames that are not allowed.
package forbidden

// Invalid_usernames is an array containing usernames that are not to be allowed.
// These usernames are used to initialize the bloom filter in the application.
var ForbiddenUsernames = [...]string{
	"admin",
	"server",
	"master",
	"super",
	"worker",
	"slave",
	"root",
	"default",
	"verified",
	"support",
	"help",
	"helpdesk",
	"cto",
	"ceo",
	"cfo",
	"chief",
	"owner",
	"me",
}
