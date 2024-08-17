package core

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"os"
	"path/filepath"
)

// Command to change the current working directory
var cdCmd = &cobra.Command{
	Use:   "cd",
	Short: "Change the current working directory",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		targetDir := args[0]

		if targetDir == ".." {
			currentDir, err := os.Getwd()
			if err != nil {
				log.Fatalf("Error getting current directory: %v", err)
			}
			targetDir = filepath.Dir(currentDir)
		}

		err := os.Chdir(targetDir)
		if err != nil {
			log.Fatalf("Error changing directory: %v", err)
		}

		newDir, err := os.Getwd()
		if err != nil {
			log.Fatalf("Error getting new directory: %v", err)
		}

		fmt.Println("Directory changed successfully")
		fmt.Println(newDir)
	},
}

func init() {
	// Register the cd command
	rootCmd.AddCommand(cdCmd)
}
