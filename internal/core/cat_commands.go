package core

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
	"github.com/thanhpk/ascii"
)

var CatCmd = &cobra.Command{
	Use:   "cat",
	Short: "Concatenate and display file contents",
	Args:  cobra.ArbitraryArgs,
	Run: func(cmd *cobra.Command, args []string) {
		number, _ := cmd.Flags().GetBool("number")
		numberNonBlank, _ := cmd.Flags().GetBool("number-nonblank")
		squeezeBlank, _ := cmd.Flags().GetBool("squeeze-blank")

		if len(args) > 1 && args[len(args)-2] == "mx" {
			inputFiles := args[:len(args)-2]

			baseNames := make([]string, len(inputFiles))
			for i, file := range inputFiles {
				base := filepath.Base(file)
				baseNames[i] = strings.TrimSuffix(base, ".txt")
			}
			fileNameParts := strings.Join(baseNames, "_")
			outputFileName := fmt.Sprintf("concatenated_%s.txt", fileNameParts)

			outputDir := "concatenated_files"

			outputFile := filepath.Join(outputDir, outputFileName)

			mergeFiles(inputFiles, outputFile)

		} else {
			if len(args) == 0 {
				return
			}

			baseNames := make([]string, len(args))
			for i, file := range args {
				base := filepath.Base(file)
				baseNames[i] = strings.TrimSuffix(base, ".txt")
			}
			fileNameParts := strings.Join(baseNames, "_")
			outputFileName := fmt.Sprintf("concatenated_%s.txt", fileNameParts)

			outputDir := "concatenated_files"

			outputFile := filepath.Join(outputDir, outputFileName)

			mergeFiles(args, outputFile)
			finalContents, err := os.ReadFile(outputFile)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error displaying contents of %s: %v\n", outputFile, err)
			} else {
				displayContents(string(finalContents), number, numberNonBlank, squeezeBlank)
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
	outputDir := "concatenated_files"

	err := os.MkdirAll(outputDir, 0755)
	if err != nil {
		log.Fatalf("Error creating directory %s: %v", outputDir, err)
	}

	var container []byte
	for i, file := range files {
		contents, err := os.ReadFile(file)
		if err != nil {
			log.Fatalf("Error reading file %s: %v", file, err)
		}

		container = append(container, contents...)

		if len(contents) > 0 && i < len(files)-1 {
			if contents[len(contents)-1] != '\n' {
				container = append(container, '\n')
			}
		}
	}

	err = os.WriteFile(outputFile, container, 0644)

	if err != nil {
		log.Fatalf("Error writing to file %s: %v", outputFile, err)
	}
	fmt.Println("Contents successfully written to", outputFile)
}

func displayContents(contents string, number, numberNonBlank, squeezeBlank bool) {
	lines := strings.Split(string(ascii.Convert(contents)), "\n")

	lineNumber := 1
	var lastLineBlank bool = true

	for i, line := range lines {
		currentLineBlank := (line == "")

		if squeezeBlank && lastLineBlank && currentLineBlank {
			continue
		}
		if numberNonBlank && !currentLineBlank {
			fmt.Printf("    %d %s\n", lineNumber, line)
			lineNumber++
		} else if number {
			fmt.Printf("    %d %s\n", i+1, line)
		} else {
			fmt.Println(line)
		}
		lastLineBlank = currentLineBlank
	}
	fmt.Println()
}
