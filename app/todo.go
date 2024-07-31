package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

var tasks []string

func main() {
    reader := bufio.NewReader(os.Stdin)

    for {
        fmt.Println("1. Add Task")
        fmt.Println("2. View Tasks")
        fmt.Println("3. Exit")
        fmt.Print("Choose an option: ")

        option, _ := reader.ReadString('\n')
        option = strings.TrimSpace(option)

        switch option {
        case "1":
            addTask(reader)
        case "2":
            viewTasks()
        case "3":
            fmt.Println("Exiting...")
            return
        default:
            fmt.Println("Invalid option. Please choose again.")
        }
    }
}

func addTask(reader *bufio.Reader) {
    fmt.Print("Enter task description: ")
    task, _ := reader.ReadString('\n')
    task = strings.TrimSpace(task)
    if task != "" {
        tasks = append(tasks, task)
        fmt.Println("Task added.")
    } else {
        fmt.Println("Task description cannot be empty.")
    }
}

func viewTasks() {
    if len(tasks) == 0 {
        fmt.Println("No tasks to display.")
        return
    }
    fmt.Println("To-Do List:")
    for i, task := range tasks {
        fmt.Printf("%d. %s\n", i+1, task)
    }
}

