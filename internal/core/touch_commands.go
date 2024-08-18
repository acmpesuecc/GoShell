package core

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
)

var TouchCmd = &cobra.Command{
	Use:   "touch [filename]",
	Short: "Create a new file",
	Long:  `The touch command creates a new file with the given filename.`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fileName := args[0]
		_, err := os.Stat(fileName)
		if os.IsNotExist(err) {
			err = os.WriteFile(fileName, []byte(""), 0666)
			if err != nil {
				log.Fatalf("Failed to create file: %v", err)
			} else {
				fmt.Println("File created successfully")
			}
		} else if err == nil {
			fmt.Println("File already exists")
		} else {
			log.Fatalf("Failed to check if file exists: %v", err)
		}
	},
}

//
//func FileRangeMake(file_names string) {
//	//	touch <filename{<start>..<finish>}>
//
//}
