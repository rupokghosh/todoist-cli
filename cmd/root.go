package cmd

import (
    "github.com/spf13/cobra"
    "log"
    "os"
)

var rootCmd = &cobra.Command{
    Use:   "todoist-cli",
    Short: "A simple CLI to interact with Todoist",
}

func Execute() {
    err := rootCmd.Execute()
    if err != nil {
        log.Fatal(err)
        os.Exit(1)
    }
}

