package cmd

import (
    "fmt"
    "log"
    "todoist-cli/todoist"
    "github.com/spf13/cobra"
)

var viewCmd = &cobra.Command{
    Use:   "view",
    Short: "View all Todoist tasks",
    Run: func(cmd *cobra.Command, args []string) {
        tasks, err := todoist.GetTasks()
        if err != nil {
            log.Fatalf("Error fetching tasks: %v", err)
        }

        for _, task := range tasks {
            fmt.Printf("%d: %s\n", task.ID, task.Content)
        }
    },
}
