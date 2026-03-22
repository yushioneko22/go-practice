package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Todo struct {
	ID   int    `json:"id"`
	Text string `json:"text"`
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: todo [add|list] [text]")
		return
	}

	subcommand := os.Args[1]
	switch subcommand {
	case "add":
		if len(os.Args) < 3 {
			fmt.Println("Please provide task text")
			return
		}
		addTodo(os.Args[2])
	case "list":
		listTodos()
	default:
		fmt.Println("Unknown subcommand")
	}
}

func loadTodos() []Todo {
	var todos []Todo
	data, err := os.ReadFile("todos.json")
	if err == nil {
		json.Unmarshal(data, &todos)
	}
	return todos
}

func saveTodos(todos []Todo) {
	data, _ := json.MarshalIndent(todos, "", "  ")
	os.WriteFile("todos.json", data, 0644)
}

func addTodo(text string) {
	todos := loadTodos()
	newTodo := Todo{
		ID:   len(todos) + 1,
		Text: text,
	}
	todos = append(todos, newTodo)
	saveTodos(todos)
	fmt.Printf("Added: %s\n", text)
}

func listTodos() {
	todos := loadTodos()
	for _, t := range todos {
		fmt.Printf("%d: %s\n", t.ID, t.Text)
	}
}
