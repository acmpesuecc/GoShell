package core

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/thanhpk/ascii"
)

var CatCmd = &cobra.Command{
	Use:   "cat",
	Short: "Concatenate and display file contents",
	Args:  cobra.ArbitraryArgs,
	Run: func(cmd *cobra.Command, args []string) {
		// Get flag values
		number, _ := cmd.Flags().GetBool("number")
		numberNonBlank, _ := cmd.Flags().GetBool("number-nonblank")
		squeezeBlank, _ := cmd.Flags().GetBool("squeeze-blank")

		if len(args) > 1 && args[len(args)-2] == "mx" {
			mergeFiles(args[:len(args)-2], args[len(args)-1])
		} else {
			for _, file := range args {
				displayFile(file, number, numberNonBlank, squeezeBlank)
			}
		}
	},
}

func init() {
	// Define the flags
	CatCmd.Flags().BoolP("number", "n", false, "Number all output lines")
	CatCmd.Flags().BoolP("number-nonblank", "b", false, "Number non-empty output lines, overrides -n")
	CatCmd.Flags().BoolP("squeeze-blank", "s", false, "Suppress repeated empty output lines")
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

func displayFile(file string, number, numberNonBlank, squeezeBlank bool) {
	contents, err := os.ReadFile(file)
	if err != nil {
		log.Fatalf("Error reading file %s: %v", file, err)
	}

	lines := strings.Split(string(ascii.Convert(string(contents))), "\n")
	fmt.Println("Contents of:", file)

	lineNumber := 1
	for i, line := range lines {
		if squeezeBlank && i > 0 && lines[i-1] == "" && line == "" {
			continue
		}
		if numberNonBlank && line != "" {
			fmt.Printf("    %d %s\n", lineNumber, line)
			lineNumber++
		} else if number {
			fmt.Printf("    %d %s\n", i+1, line)
		} else {
			fmt.Println(line)
		}
	}
	fmt.Println()
}
