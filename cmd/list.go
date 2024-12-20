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

	
func init () {
    rootCmd.AddCommand(listCmd)
}
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
			Priority int `json:"priority"`
			DueDate struct {
				Date string `json:"date"`
                IsRecurring bool `json:"is_recurring"`
			} `json:"due"`
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
           
            recurring := ""
            if task.DueDate.IsRecurring {
                recurring = "recurring task"
            }
            // everything starting with \033 is just color coding
            fmt.Printf("\033[36m%d %s \033[35m[Due: %s] \033[31m[%s%d] \033[36m%s\n\n", i+1, task.Content, task.DueDate.Date, "priority-", task.Priority, recurring)
        }
       
    },
}


