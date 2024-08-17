package core

import (
	"github.com/spf13/cobra"
)

// InitCommands initializes and adds all commands to the provided rootCmd
func InitCommands(rootCmd *cobra.Command) {
	rootCmd.AddCommand(lsCmd)
	rootCmd.AddCommand(catCmd)
	rootCmd.AddCommand(rmCmd)
	rootCmd.AddCommand(cdCmd)
}
