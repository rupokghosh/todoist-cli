package cmd

import (
    "os"
    "github.com/spf13/cobra"
)
var rootCmd = &cobra.Command{
    Use:   "todoist-cli",
    Short: "A CLI for interacting with Todoist",
}
// Execute runs the root command
func Execute() {
    if err := rootCmd.Execute(); err != nil {
        os.Exit(1)
    }
}

