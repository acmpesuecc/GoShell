package core

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"os"
	"path/filepath"
)

var CdCmd = &cobra.Command{
	Use:   "cd",
	Short: "Change the current working directory",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		targetDir := args[0]
		
		// Handle parent directory shortcut
		if targetDir == ".." {
			currentDir, err := os.Getwd()
			if err != nil {
				log.Fatalf("Error getting current directory: %v", err)
			}
			targetDir = filepath.Dir(currentDir)
		}
		
		// Convert to absolute path
		absPath, err := filepath.Abs(targetDir)
		if err != nil {
			log.Fatalf("Error resolving path: %v", err)
		}
		
		// Check if directory exists
		fileInfo, err := os.Stat(absPath)
		if os.IsNotExist(err) {
			log.Fatalf("Directory does not exist: %s", absPath)
		}
		if err != nil {
			log.Fatalf("Error accessing directory: %v", err)
		}
		if !fileInfo.IsDir() {
			log.Fatalf("Not a directory: %s", absPath)
		}
		
		// Output the path with special prefix for shell wrapper
		fmt.Printf("GOSHELL_CD:%s\n", absPath)
	},
}

