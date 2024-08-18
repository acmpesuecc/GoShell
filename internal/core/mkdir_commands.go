package core

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"os"
)

var MkdirCmd = &cobra.Command{
	Use:   "mkdir [dirName]",
	Short: "Create a new file",
	Long:  `The touch command creates a new file with the given dir_Name.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 1 {
			Dir_Maker(args[0])
		} else if len(args) > 1 {
			for _, i := range args[:] {
				Dir_Maker(i)
			}
		}
	},
}

func Dir_Maker(dir_Name string) {
	_, err := os.Stat(dir_Name)
	if os.IsNotExist(err) {
		errs := os.Mkdir(dir_Name, 0750)
		if errs != nil {
			log.Fatalf("Failed to create directory: %v", err)
		} else {
			fmt.Println("File created successfully")
		}
	} else if err == nil {
		fmt.Println("File already exists")
	} else {
		log.Fatalf("Failed to check if file exists: %v", err)
	}
}
