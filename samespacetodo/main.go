package main

import (
	"fmt"
)

func main() {
	// Create a new TODO item
	createTodo("todo-1.json", &Todo{
		ID:          "1",
		UserID:      "user1",
		Title:       "My First Todo",
		Description: "This is my first todo item",
		Status:      TodoStatusPending,
	})

	// Retrieve a specific TODO item
	todo, _ := getTodo("todo-1.json")
	fmt.Printf("Retrieved TODO item: %+v\n", todo)

	// Update the TODO item
	todo.Status = TodoStatusCompleted
	updateTodo("todo-1.json", todo)

	// List TODO items, filtering by status
	todos := listTodos("", TodoStatusPending)
	fmt.Printf("Pending TODO items: %+v\n", todos)

	// Delete the TODO item
	deleteTodo("todo-1.json")
}