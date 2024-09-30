package main

import (
    "log"

    "github.com/joho/godotenv"
    "todoist-cli/cmd"
)

func main() {
    // Load .env file
    err := godotenv.Load()
    if err != nil {
        log.Fatalf("Error loading .env file")
    }
    cmd.Execute()
}
