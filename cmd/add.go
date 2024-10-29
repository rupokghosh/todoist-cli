package cmd

import (
	"log"
	"fmt"
	"os"
	"bufio"
	"github.com/spf13/cobra"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
	"time"
	"strconv"
	"strings"
	
)

func init () {
	rootCmd.AddCommand(addCmd)
}

var addCmd = &cobra.Command{
	Use: "add", 
	Short: "Add your to-do to the list",
	Run: func (cmd *cobra.Command, args []string){

		var tasks Task
		id := uuid.New().String()	// generate uuid for header
		time := time.Now().Format("2006-01-02")

		// read input and give it to task
		scanner := bufio.NewReader(os.Stdin)
		fmt.Println("Name your task:")
		tasks.Content, _ = scanner.ReadString('\n')
		tasks.Content = strings.TrimSpace(tasks.Content)

		fmt.Println("Enter the due date(in the format: 'tomorrow at 12:00 or YYYY-MM-DD')\nOr leave empty for today's date:")
		tasks.Due, _ = scanner.ReadString('\n')
		tasks.Due = strings.TrimSpace(tasks.Due)
		if tasks.Due == "" {
			tasks.Due = time
		}
		
		fmt.Println("Enter tasks priority(4 is highest, 1 is lowest)")
		priority, _ := scanner.ReadString('\n')
		priority = strings.TrimSpace(priority)
		Priority, err := strconv.Atoi(priority)
		if Priority < 1 || Priority > 4 {
			fmt.Println("Priority has to be between 1 and 4")
			os.Exit(1)
		}
		if err != nil {
			log.Fatalf("error %s", err)
		}
		tasks.Priority = Priority

		tasks.DueLang = "en"

		apiToken := os.Getenv("TODOIST_API_TOKEN")
		if apiToken == "" {
			log.Fatal("token not found")
		}
		client := resty.New()
		resp , err := client.R().
			SetHeader("Content-Type", "application/json").
			SetHeader("X-Request-Id", id ).
			SetHeader("Authorization", " Bearer "+apiToken).
			SetBody(tasks).
			Post(todoistAPIBase + "/tasks")
		
		if err != nil {
			log.Fatal("Couldn't add task")
		}
		_ = resp 
		fmt.Println("\033[36mTask successfully added!")
	},
}
