package main

import (
    "fmt"
    "log"
    "os"
    "todoist-cli/cmd"
)

func main() {
    if err := cmd.RootCmd.Execute(); err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
}
