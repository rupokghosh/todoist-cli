package cmd

import (
	"github.com/spf13/cobra"
	"fmt"
	
)

func init() {
	rootCmd.AddCommand(doneCmd)
}
var doneCmd = &cobra.Command{
	Use: "done", 
	Short: "Command to mark a task as done", 
	Run: func (cmd *cobra.Command, args []string){
		fmt.Println("done!")
	},
 }