package cmd

import (
    "github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
    Use:   "todoist-cli",
    Short: "A CLI tool to interact with Todoist",
    Long:  "Manage your Todoist tasks directly from the command line.",
}

func init() {
    RootCmd.AddCommand(viewCmd)
    RootCmd.AddCommand(addCmd)
    RootCmd.AddCommand(deleteCmd)
    RootCmd.AddCommand(updateCmd)
}
