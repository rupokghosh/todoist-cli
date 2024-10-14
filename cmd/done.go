package cmd

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/go-resty/resty/v2"
	"github.com/spf13/cobra"
)

var taskIDs []string // To store task IDs globally

func init() {
	rootCmd.AddCommand(doneCmd)
}

// doneCmd represents the done command
var doneCmd = &cobra.Command{
	Use:   "done [task_index]",
	Short: "Command to mark a task as done",
	Args:  cobra.ExactArgs(1), 
	Run: func(cmd *cobra.Command, args []string) {

		err := fetchTasks()
		if err != nil {
			log.Fatalf("Failed to fetch tasks: %v", err)
		}
		// Parse the task index provided by the user
		index, err := strconv.Atoi(args[0])
		if err != nil || index <= 0 || index > len(taskIDs) {
			log.Fatalf("Choose from the list of current tasks please%d", len(taskIDs))
		}

		// Get the task ID based on the selected index 
		taskID := taskIDs[index-1]	// ( we do index -1 since our tasks start with 1 but array index start with 0)

		client := resty.New()
		apiToken := os.Getenv("TODOIST_API_TOKEN")

		resp, err := client.R().
			SetHeader("Authorization", "Bearer "+apiToken).
			Post(fmt.Sprintf("%s/tasks/%s/close", todoistAPIBase, taskID))
		if err != nil {
			log.Fatalf("Couldn't close the task: %v", err)
		}
		_ = resp
		fmt.Println("Task marked as finished!")
	},
}

// fetchTasks fetches tasks and populates the taskIDs slice
func fetchTasks() error {
	client := resty.New()
	apiToken := os.Getenv("TODOIST_API_TOKEN")

	var tasks []struct {
		ID string `json:"id"`
	}

	// Make a GET request to fetch the tasks
	resp, err := client.R().
		SetHeader("Authorization", "Bearer "+apiToken).
		SetResult(&tasks). 
		Get(fmt.Sprintf("%s/tasks", todoistAPIBase))

	if err != nil {
		return err
	}
	_ = resp
	// Populate the taskIDs slice with task IDs from the response
	taskIDs = []string{}
	for _, task := range tasks {
		taskIDs = append(taskIDs, task.ID)
	
	}
	return nil
}
