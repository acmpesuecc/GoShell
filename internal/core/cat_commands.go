package core

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/thanhpk/ascii"
)

// Command to concatenate and display file contents
var catCmd = &cobra.Command{
	Use:   "cat",
	Short: "Concatenate and display file contents",
	Args:  cobra.ArbitraryArgs,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 1 && args[len(args)-2] == "mx" {
			mergeFiles(args[:len(args)-2], args[len(args)-1])
		} else if len(args) > 0 && args[0] == "-n" {
			displayFilesWithLineNumbers(args[1:])
		} else {
			displayFiles(args)
		}
	},
}

func mergeFiles(files []string, outputFile string) {
	var container []byte
	for _, file := range files {
		contents, err := os.ReadFile(file)
		if err != nil {
			log.Fatalf("Error reading file %s: %v", file, err)
		}
		container = append(container, contents...)
	}

	err := os.WriteFile(outputFile, container, 0644)
	if err != nil {
		log.Fatalf("Error writing to file %s: %v", outputFile, err)
	}
	fmt.Println("Contents successfully written to", outputFile)
}

func displayFiles(files []string) {
	for _, file := range files {
		contents, err := os.ReadFile(file)
		if err != nil {
			log.Fatalf("Error reading file %s: %v", file, err)
		}
		fmt.Println("Contents of:", file)
		fmt.Println(ascii.Convert(string(contents)))
		fmt.Println()
	}
}

func displayFilesWithLineNumbers(files []string) {
	for _, file := range files {
		contents, err := os.ReadFile(file)
		if err != nil {
			log.Fatalf("Error reading file %s: %v", file, err)
		}
		fmt.Println("Contents of:", file)
		lines := strings.Split(string(ascii.Convert(string(contents))), "\n")
		for i, line := range lines {
			fmt.Printf("    %d %s\n", i+1, line)
		}
	}
}

func init() {
	// Register the cat command
	rootCmd.AddCommand(catCmd)
}
