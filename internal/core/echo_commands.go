package core

import (
    "fmt"
    "strings"

    "github.com/spf13/cobra"
)

// EchoCmd prints its arguments separated by spaces.
var EchoCmd = &cobra.Command{
    Use:   "echo",
    Short: "Print arguments to the console",
    Run: func(cmd *cobra.Command, args []string) {
        fmt.Println(strings.Join(args, " "))
    },
}
