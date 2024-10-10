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

// Execute runs the root command (called from main)
func Execute() {
    if err := rootCmd.Execute(); err != nil {
        log.Fatal(err)
        os.Exit(1)
    }
}

func init() {
    
    rootCmd.AddCommand(listCmd)
    
}
