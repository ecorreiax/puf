package cmd

import (
	"crypto/sha1"
	"fmt"

	"github.com/ecorreiax/gobfs/internal/hash"
	"github.com/spf13/cobra"
)

var checkCmd = &cobra.Command{
	Use:   "check",
	Short: "Print the usage status of a username",
	Long:  `Check if a given username should be allowed or not`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		h := sha1.New()
		u := args[0]
		idx := hash.CreateHash(h, u)

		result := hash.VerifyHash(idx)

		if result {
			fmt.Printf("Username %s is not allowed\n", u)
		} else {
			fmt.Printf("Username %s is ok to be used\n", u)
		}
	},
}
