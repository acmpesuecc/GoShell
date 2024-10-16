package main

import (
	"os"

	"github.com/IshaanNene/GoShell/internal/core"
	"github.com/spf13/cobra"
)

func main() {
	rootCmd := &cobra.Command{
		Use:   "mycli",
		Short: "A CLI tool for file and directory operations",
	}
	rootCmd.AddCommand(core.LsCmd)
	rootCmd.AddCommand(core.CdCmd)
	rootCmd.AddCommand(core.MkdirCmd)
	rootCmd.AddCommand(core.RmCmd)
	rootCmd.AddCommand(core.TouchCmd)
	rootCmd.AddCommand(core.CatCmd)

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
