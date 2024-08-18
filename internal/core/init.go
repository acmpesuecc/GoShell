package core

import (
	"github.com/spf13/cobra"
)

func InitCommands(rootCmd *cobra.Command) {
	rootCmd.AddCommand(LsCmd)
	rootCmd.AddCommand(CatCmd)
	rootCmd.AddCommand(RmCmd)
	rootCmd.AddCommand(CdCmd)
	rootCmd.AddCommand(dateCmd)
	rootCmd.AddCommand(pwdCmd)
	rootCmd.AddCommand(TouchCmd)
	rootCmd.AddCommand(MkdirCmd)
}
