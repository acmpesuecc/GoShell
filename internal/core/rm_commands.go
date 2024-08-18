package core

import (
	"bufio"
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"os"
)

var RmCmd = &cobra.Command{
	Use:   "rm",
	Short: "Remove files or directories",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("Usage: rm [flags] [file ...]")
			return
		}

		// Check for flags
		interactive, _ := cmd.Flags().GetBool("interactive")
		confirmAll, _ := cmd.Flags().GetBool("confirm-all")
		recursive, _ := cmd.Flags().GetBool("recursive")
		force, _ := cmd.Flags().GetBool("force")

		if interactive {
			for _, file := range args {
				RemoveI_Method(file)
			}
		} else if confirmAll {
			fmt.Print("Remove all specified files? (y/n): ")
			reader := bufio.NewReader(os.Stdin)
			char, _, err := reader.ReadRune()
			if err != nil {
				fmt.Println("Error reading input:", err)
				return
			}
			if char == 'Y' || char == 'y' {
				for _, file := range args {
					removeFile(file)
				}
			} else {
				fmt.Println("Files not removed")
			}
		} else if recursive {
			if len(args) != 1 {
				fmt.Println("Usage: rm -r [directory]")
				return
			}
			removeDir(args[0])
		} else if force {
			for _, file := range args {
				removeFile(file)
			}
		} else {
			for _, file := range args {
				removeFile(file)
			}
		}
	},
}

func init() {
	RmCmd.Flags().BoolP("interactive", "i", false, "Prompt before each removal")
	RmCmd.Flags().BoolP("confirm-all", "I", false, "Prompt once before removing all files")
	RmCmd.Flags().BoolP("recursive", "r", false, "Remove directories and their contents recursively")
	RmCmd.Flags().BoolP("force", "f", false, "Ignore nonexistent files and arguments, never prompt")
}

func removeFile(fileName string) {
	fi, err := os.Lstat(fileName)
	if err != nil {
		log.Fatalf("Error checking file: %v", err)
	}
	if fi.Mode().IsDir() {
		fmt.Printf("%s is a directory\n", fileName)
	} else {
		err := os.Remove(fileName)
		if err != nil {
			log.Fatalf("Error removing file: %v", err)
		}
		fmt.Printf("%s removed successfully\n", fileName)
	}
}

func removeDir(path string) {
	err := os.RemoveAll(path)
	if err != nil {
		log.Fatalf("Error removing directory: %v", err)
	}
	fmt.Printf("Directory %s removed successfully\n", path)
}

func RemoveI_Method(file string) {
	fmt.Printf("Remove %s? (y/n): ", file)
	reader := bufio.NewReader(os.Stdin)
	char, _, err := reader.ReadRune()
	if err != nil {
		fmt.Println("Error reading input:", err)
	}
	if char == 'Y' || char == 'y' {
		removeFile(file)
	} else {
		fmt.Printf("%s not removed\n", file)
	}
}
