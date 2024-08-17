package main

import (
	"fmt"
	"github.com/IshaanNene/GoShell/internal/core"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "goshell",
	Short: "A simple shell command implementation",
}

func init() {
	// Register commands
	rootCmd.AddCommand(core.CatCmd)
	rootCmd.AddCommand(core.RmCmd)
	rootCmd.AddCommand(core.CdCmd)
	rootCmd.AddCommand(core.LsCmd)
	// Add other commands as needed
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
