package core

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"os"
)

var TouchCmd = &cobra.Command{
	Use:   "touch [filename]",
	Short: "Create a new file",
	Long:  `The touch command creates a new file with the given filename.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 1 {
			File_Maker(args[0])
		} else if len(args) > 1 {
			for _, i := range args[:] {
				File_Maker(i)
			}
		}
	},
}

func File_Maker(fileName string) {
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
}

//
//func FileRangeMake(file_names string) {
//	//	touch <filename{<start>..<finish>}>
//
//}
