package main

import (
    "encoding/json"
    "fmt"
    "net/http"
    "strconv"
    "sync"
)

type Task struct {
    ID          int    `json:"id"`
    Description string `json:"description"`
    Completed   bool   `json:"completed"`
}

var (
    tasks  []Task
    nextID int
    mu     sync.Mutex
)

func main() {
    http.HandleFunc("/add", addTaskHandler)
    http.HandleFunc("/view", viewTasksHandler)
    http.HandleFunc("/complete", completeTaskHandler)

    fmt.Println("Starting server on :8080")
    http.ListenAndServe(":8080", nil)
}

func addTaskHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
        return
    }

    description := r.FormValue("description")
    if description == "" {
        http.Error(w, "Description is required", http.StatusBadRequest)
        return
    }

    mu.Lock()
    defer mu.Unlock()

    nextID++
    task := Task{ID: nextID, Description: description, Completed: false}
    tasks = append(tasks, task)

    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(task)
}

func viewTasksHandler(w http.ResponseWriter, r *http.Request) {
    mu.Lock()
    defer mu.Unlock()

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(tasks)
}

func completeTaskHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
        return
    }

    idStr := r.FormValue("id")
    id, err := strconv.Atoi(idStr)
    if err != nil || id <= 0 {
        http.Error(w, "Invalid ID", http.StatusBadRequest)
        return
    }

    mu.Lock()
    defer mu.Unlock()

    for i, task := range tasks {
        if task.ID == id {
            tasks[i].Completed = true
            w.WriteHeader(http.StatusOK)
            json.NewEncoder(w).Encode(tasks[i])
            return
        }
    }

    http.Error(w, "Task not found", http.StatusNotFound)
}
