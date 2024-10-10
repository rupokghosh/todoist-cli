package cmd

import (
    "github.com/spf13/cobra"
    "log"
    "os"
)

var rootCmd = &cobra.Command{
    Use:   "todoist",
    Short: "A simple CLI to interact with Todoist",
}

func Execute() {
    if err := rootCmd.Execute(); err != nil {
        log.Fatal(err)
        os.Exit(1)
    }
}

