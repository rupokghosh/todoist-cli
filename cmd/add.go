package cmd

import (
	"fmt"
	"os"
	"bufio"
	"github.com/spf13/cobra"
)

func init () {
	rootCmd.AddCommand(addCmd)
}
var addCmd = &cobra.Command{
	Use: "add", 
	Short: "Add your to-do to the list",
	Run: func (cmd *cobra.Command, args []string){
		fmt.Println("add cmd")
		var todo string
		scanner := bufio.NewReader(os.Stdin)
		fmt.Println("Add new todo:")
		todo, _ = scanner.ReadString('\n')
		fmt.Println(todo)		
	},
}
