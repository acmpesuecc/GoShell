package core

import (
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"syscall"
	"time"

	"github.com/dustin/go-humanize"
	"github.com/spf13/cobra"
)

type FileTimeStruct struct {
	Fname string
	Ftime time.Time
}

func checkError(err error, context string) {
	if err != nil {
		log.Fatalf("Error %s: %v", context, err)
	}
}
func listFiles(dir string, showHidden bool, appendSlashToDir bool, sortByTime bool, listInode bool, humanReadable bool, out io.Writer) {
	files, err := os.ReadDir(dir)
	checkError(err, "reading directory")

	if sortByTime {
		var fileTimes []FileTimeStruct
		for _, file := range files {
			info, err := file.Info()
			checkError(err, "getting file info")
			fileTimes = append(fileTimes, FileTimeStruct{Fname: file.Name(), Ftime: info.ModTime()})
		}
		sort.Slice(fileTimes, func(i, j int) bool {
			return fileTimes[i].Ftime.Before(fileTimes[j].Ftime)
		})
		for _, file := range fileTimes {
			fmt.Fprintln(out, file.Fname)
		}
		return
	}

	for _, file := range files {
		name := file.Name()

		if !showHidden && name[0] == '.' {
			continue
		}

		if listInode {
			info, err := os.Stat(file.Name())
			checkError(err, "getting file stat")
			stat := info.Sys().(*syscall.Stat_t)
			fmt.Fprintf(out, "%d ", stat.Ino)
		}

		if appendSlashToDir && file.IsDir() {
			name += "/"
		}
		if humanReadable {
			info, err := os.Stat(file.Name())
			checkError(err, "getting file info")
			fmt.Fprintf(out, "%s %s\n", humanize.Bytes(uint64(info.Size())), name)
		} else {
			fmt.Fprintln(out, name)
		}
	}
}

var LsCmd = &cobra.Command{
	Use:   "ls",
	Short: "List directory contents",
	Run: func(cmd *cobra.Command, args []string) {
		dir, _ := cmd.Flags().GetString("directory")
		showHidden, _ := cmd.Flags().GetBool("a")
		appendSlashToDir, _ := cmd.Flags().GetBool("F")
		sortByTime, _ := cmd.Flags().GetBool("t")
		listInode, _ := cmd.Flags().GetBool("i")
		humanReadable, _ := cmd.Flags().GetBool("V")

		listFiles(dir, showHidden, appendSlashToDir, sortByTime, listInode, humanReadable, cmd.OutOrStdout())
	},
}

func init() {
	LsCmd.Flags().StringP("directory", "d", ".", "Directory to list")
	LsCmd.Flags().BoolP("a", "a", false, "Include hidden files")
	LsCmd.Flags().BoolP("F", "F", false, "Append indicator (one of */=>@|) to entries")
	LsCmd.Flags().BoolP("t", "t", false, "Sort by modification time, newest first")
	LsCmd.Flags().BoolP("i", "i", false, "Print the index number of each file")
	LsCmd.Flags().BoolP("V", "V", false, "print sizes like 1K 234M 2G")
}
