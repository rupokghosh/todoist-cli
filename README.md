# Todoist CLI

This is a terminal app that connects your todoist tasks with your terminal. 


Made with:
- Go
- Libraries:
    - Cobra
    - Resty

## How to use and API TOKEN setup

To use this app, you would need to create an .env file
and in that env file store your api token as:
``` 
TODOIST_API_TOKEN = whatever your token is
```

To get your token:
- Go to Settings > Integrations > Developer 
- Then copy your API TOKEn

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

  
  
