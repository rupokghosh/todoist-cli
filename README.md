# Todoist CLI

This is a terminal app that connects your todoist tasks with your terminal.

Made with:

- Go
- Libraries:
  - Cobra
  - Resty

## How to use and API TOKEN setup

1. Clone the repository:

   ```zsh
   git clone https://github.com/yourusername/todoist-cli.git
   cd todoist-cli
   ```

2. Create a `.env` file at the root of the project and add your Todoist API token:
   ```zsh
   TODOIST_API_TOKEN=your_api_token_here
   ```
   To get your token:

- Go to Settings > Integrations > Developer
- Then copy your API TOKEN

3. Install the required dependencies:

   ```zsh
   go get
   ```

4. Build the binary app:
   ```zsh
   go build -o todoist-cli
   ```
5. Run the binary from the current directory:

    ```zsh
    ./todoist-cli list
    ./todoist-cli add "Your task"
    ./todoist-cli done 1
    ```

6. Additionally you can also use it without building a binary by doing:

    ```zsh
    go run main.go list
    go run main.go add
    ```

## Features and Commands

- List all active tasks

```zsh
    todoist-cli list
```

- Add new tasks.

```zsh
    todoist-cli add "My new task"
```

- Mark tasks as complete.

```zsh
    todoist-cli done 2 // number of the task in the list
```

## Contributing

Feel free to fork this repository and submit pull requests to improve the functionality or fix bugs.
