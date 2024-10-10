package cmd

import (
    "encoding/json"
    "fmt"
    "log"
    "os"
    "github.com/go-resty/resty/v2"
    "github.com/spf13/cobra"
)

const todoistAPIBase = "https://api.todoist.com/rest/v2"

var listCmd = &cobra.Command{
    Use:   "list",
    Short: "List current active tasks",
    Run: func(cmd *cobra.Command, args []string) {
        apiToken := os.Getenv("TODOIST_API_TOKEN")
        if apiToken == "" {
            log.Fatal("API token not found. Please set TODOIST_API_TOKEN in your environment variables.")
        }

        client := resty.New()
        resp, err := client.R().
            SetHeader("Authorization", "Bearer "+apiToken).
            Get(todoistAPIBase + "/tasks")

        if err != nil {
            log.Fatalf("Error fetching tasks: %v", err)
        }

        var tasks []struct {
            ID      string `json:"id"`
            Content string `json:"content"`
        }

        err = json.Unmarshal(resp.Body(), &tasks)
        if err != nil {
            log.Fatalf("Error parsing tasks: %v", err)
        }

        if len(tasks) == 0 {
            fmt.Println("No active tasks found.")
            return
        }

        for i, task := range tasks {
            fmt.Printf("%d. %s\n", i+1, task.Content)
        }
    },
}

func init() {
    // Register listCmd to the rootCmd
    rootCmd.AddCommand(listCmd)
}
