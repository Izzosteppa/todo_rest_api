package handlers

import (
    "encoding/json"
    "net/http"
    "go-todo-api/models"
)

var todos []models.Todo

func GetTodos(w http.ResponseWriter, r *http.Request) {
    json.NewEncoder(w).Encode(todos)
}

func CreateTodo(w http.ResponseWriter, r *http.Request) {
    var todo models.Todo
    json.NewDecoder(r.Body).Decode(&todo)
    todos = append(todos, todo)
    json.NewEncoder(w).Encode(todo)
}
