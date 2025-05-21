package main

import (
    "net/http"
    "github.com/gorilla/mux"
    "go-todo-api/handlers"
)

func main() {
    r := mux.NewRouter()
    r.HandleFunc("/todos", handlers.GetTodos).Methods("GET")
    r.HandleFunc("/todos", handlers.CreateTodo).Methods("POST")
    http.ListenAndServe(":8080", r)
}
