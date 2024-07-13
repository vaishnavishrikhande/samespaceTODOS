package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

type TodoStatus string

const (
	TodoStatusPending  TodoStatus = "pending"
	TodoStatusCompleted TodoStatus = "completed"
)

type Todo struct {
	ID          string     `json:"id"`
	UserID      string     `json:"user_id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Status      TodoStatus `json:"status"`
}

func createTodo(filename string, todo *Todo) {
	saveTodoToFile(filename, todo)
}

func getTodo(filename string) (*Todo, error) {
	return loadTodoFromFile(filename)
}

func updateTodo(filename string, todo *Todo) {
	saveTodoToFile(filename, todo)
}

func listTodos(dir string, status TodoStatus) []*Todo {
	if dir == "" {
		dir = "."
	}
	var todos []*Todo

	files, _ := ioutil.ReadDir(dir)
	for _, file := range files {
		if !file.IsDir() && strings.HasPrefix(file.Name(), "todo-") {
			todo, err := loadTodoFromFile(filepath.Join(dir, file.Name()))
			if err == nil && (status == "" || todo.Status == status) {
				todos = append(todos, todo)
			}
		}
	}

	return todos
}

func deleteTodo(filename string) {
	os.Remove(filename)
}

func saveTodoToFile(filename string, todo *Todo) error {
	// Create the directory if it doesn't exist
	dir := filepath.Dir(filename)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err := os.MkdirAll(dir, 0755)
		if err != nil {
			return err
		}
	}

	// Save the TODO item to the file
	data, err := json.Marshal(todo)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(filename, data, 0644)
	if err != nil {
		return err
	}

	return nil
}

func loadTodoFromFile(filename string) (*Todo, error) {
	// Load the TODO item from the file
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	// Unmarshal the JSON data into a Todo struct
	var todo Todo
	err = json.Unmarshal(data, &todo)
	if err != nil {
		return nil, err
	}

	return &todo, nil
}