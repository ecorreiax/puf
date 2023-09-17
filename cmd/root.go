// Package cmd contains the Cobra-based command-line interface for gobfs.
//
// gobfs is designed for fast username validation using a bloom filter.
//
// Functions:
//   - Execute: Initiates the root Cobra command.
//   - init:    Adds child commands to the root command during package initialization.
//
// Commands:
//   - rootCmd: The root command for the gobfs CLI.
package cmd

import (
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands.
//
// It provides a brief description and a longer help text detailing the program's functionality.
var rootCmd = &cobra.Command{
	Use:   "gobfs",
	Short: "Gobfs is a very fast username validator",
	Long: `A Fast bloom's filter to validate if a specific username 
		   is definitely not present in a dataset`,
}

// Execute runs the root command and returns any errors.
//
// This is the entry point for the CLI, and should be invoked in your main function.
func Execute() error {
	return rootCmd.Execute()
}

// init adds child commands to the root command during package initialization.
//
// This function is automatically called when the package is imported and attaches 'checkCmd' as a subcommand of 'rootCmd'.
func init() {
	rootCmd.AddCommand(checkCmd)
}
