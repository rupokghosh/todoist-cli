package cmd

import (
	"fmt"
	"os"
	"bufio"
	"github.com/spf13/cobra"
)
var addCmd = &cobra.Command{
	Use: "add", 
	Short: "Add your to-do to the list",
	Run: func (cmd *cobra.Command, args []string){
		fmt.Println("add cmd")
		var todo string
		scan := bufio.NewReader(os.Stdin)
		fmt.Println("Add new todo:")
		todo, _ = scan.ReadString('\n')
		fmt.Println(todo)		
	},
}
