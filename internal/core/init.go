package core

import (
	"github.com/spf13/cobra"
)

// InitCommands initializes and adds all commands to the provided rootCmd
func InitCommands(rootCmd *cobra.Command) {
	rootCmd.AddCommand(LsCmd)
	rootCmd.AddCommand(CatCmd)
	rootCmd.AddCommand(RmCmd)
	rootCmd.AddCommand(CdCmd)
	rootCmd.AddCommand(dateCmd)
	rootCmd.AddCommand(pwdCmd)
}
