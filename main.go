package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type TodoItem struct {
	ID          int
	Todo        string
	IsCompleted bool
}

func main() {
	jsonData, err := os.ReadFile("data.json")
	if err != nil {
		fmt.Println("Failed to read file")
	}
	todo := []TodoItem{}

	err = json.Unmarshal(jsonData, &todo)
	if err != nil {
		fmt.Println("Failed to convert json to slice")
	}

	var choice string
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print(`
		░▀█▀░█▀█░█▀▄░█▀█░░░█▀▀░█░░░▀█▀
		░░█░░█░█░█░█░█░█░░░█░░░█░░░░█░
		░░▀░░▀▀▀░▀▀░░▀▀▀░░░▀▀▀░▀▀▀░▀▀▀
		`)

	for {
		fmt.Print("\n Enter a command: 'add' , 'list', 'complete' or 'q' / 'quit' : ")
		scanner.Scan()
		choice = scanner.Text()
		if err := scanner.Err(); err != nil {
			fmt.Println("Err reading input:", err)
		}

		choice = strings.TrimSpace(choice)

		if choice == "q" || choice == "quit" {
			saveData(todo)
			break
		} else if choice == "list" {
			// Calling list function
			listTasks(todo)
		} else if choice == "complete" {
			// Calling mark function
			todo = markTaskComplete(todo)
		} else if choice == "add" {
			// Calling add function
			todo = addTask(scanner, todo)
		} else {
			fmt.Println("Not a command")
		}

	}
}

// / ================ HELPER FUNCTIONS ======================== ///

func listTasks(todo []TodoItem) {
	for _, item := range todo {
		if item.IsCompleted {
			fmt.Println("\n  [X] Id: ", item.ID, "Item: ", item.Todo)
		} else {
			fmt.Println("\n  [ ] Id: ", item.ID, "Item: ", item.Todo)
		}
	}
}

func addTask(scanner *bufio.Scanner, todo []TodoItem) []TodoItem {
	fmt.Print("\n Add task : ")
	scanner.Scan()
	input := scanner.Text()
	if err := scanner.Err(); err != nil {
		fmt.Println("Err reading input:", err)
	}

	newTodo := TodoItem{
		ID:          len(todo) + 1,
		Todo:        input,
		IsCompleted: false,
	}
	todo = append(todo, newTodo)
	return todo
}

func markTaskComplete(todo []TodoItem) []TodoItem {
	if len(todo) == 0 {
		fmt.Println("To items to complete yet!")
		return todo
	}

	var targetID int

	fmt.Print("\n Enter the id: ")
	_, err := fmt.Scanln(&targetID)
	if err != nil {
		fmt.Print("Please enter a valid id")
	}

	for index, item := range todo {
		if item.ID == targetID {
			todo[index].IsCompleted = true
		}
	}

	return todo
}

func saveData(todo []TodoItem) {
	jsonData, err := json.Marshal(todo)
	if err != nil {
		fmt.Println("Failed to save in database", err)
	}

	err = os.WriteFile("data.json", jsonData, 0644)
	if err != nil {
		fmt.Println("Failed to save in file", err)
	}
}
