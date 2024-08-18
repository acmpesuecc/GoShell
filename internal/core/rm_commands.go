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

		flag, files := args[0], args[1:]

		switch flag {
		case "-i":
			for _, file := range files {
				RemoveI_Method(file)
			}

		case "-I":
			fmt.Print("Remove all specified files? (y/n): ")
			reader := bufio.NewReader(os.Stdin)
			char, _, err := reader.ReadRune()
			if err != nil {
				fmt.Println("Error reading input:", err)
				return
			}
			if char == 'Y' || char == 'y' {
				for _, file := range files {
					removeFile(file)
				}
			} else {
				fmt.Println("Files not removed")
			}

		case "-r":
			if len(files) != 1 {
				fmt.Println("Usage: rm -r [directory]")
				return
			}
			removeDir(files[0])

		case "-f":
			for _, file := range files {
				removeFile(file)
			}

		default:
			for _, file := range args {
				removeFile(file)
			}
		}
	},
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
