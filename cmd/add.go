package cmd

import (
	"log"
	"fmt"
	"os"
	"bufio"
	"github.com/spf13/cobra"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
)

const todoistAPIBase = "https://api.todoist.com/rest/v2"

type Task struct {
	Content string `json:"content"`
	Due string `json:"due_string"`
	Priority string `json:"priority"`
	DueLang string `json:"due_lang"`
}

func init () {
	rootCmd.AddCommand(addCmd)
}

var addCmd = &cobra.Command{
	Use: "add", 
	Short: "Add your to-do to the list",
	Run: func (cmd *cobra.Command, args []string){

		var tasks Task
		id := uuid.New().String()	// generate uuid for header

		scanner := bufio.NewReader(os.Stdin)
		fmt.Println("Name your task:")
		tasks.Content, _ = scanner.ReadString('\n')
		tasks.Due, _ = scanner.ReadString('\n')
		tasks.Priority, _ = scanner.ReadString('\n')
		tasks.DueLang = "en"
		fmt.Println(tasks)		
		apiToken := os.Getenv("TODOIST_API_TOKEn")
		if apiToken == "" {
			log.Fatal("token not found")
		}
		client := resty.New()
		resp, err := client.R().
			SetHeader("Content-Type", "application/json").
			SetHeader("X-Request-Id", id ).
			SetHeader("Authorizaition", "Bearer" + apiToken).
			SetBody(tasks).
			Post(todoistAPIBase + "/tasks")
		
		if err != nil {
			log.Fatal("Couldn't add task")
		}
	},
}
