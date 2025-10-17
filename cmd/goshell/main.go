package main

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/IshaanNene/GoShell/internal/core"
)

func main() {
	rootCmd := &cobra.Command{
		Use:   "./goshell",
		Short: "For all usage",
	}
	core.InitCommands(rootCmd)
	rootCmd.AddCommand(core.LsCmd)
	rootCmd.AddCommand(core.CdCmd)
	rootCmd.AddCommand(core.MkdirCmd)
	rootCmd.AddCommand(core.RmCmd)
	rootCmd.AddCommand(core.TouchCmd)
	rootCmd.AddCommand(core.PwdCmd)
	rootCmd.AddCommand(core.CatCmd)
	rootCmd.AddCommand(core.DateCmd)
	rootCmd.AddCommand(core.Iamwho)
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
