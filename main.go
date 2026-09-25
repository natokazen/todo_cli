package main

import (
	"bufio"
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
	todo := []TodoItem{}

	var choice string
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("Enter a command: 'add' , 'list', 'complete' or 'q' / 'quit' :")
		scanner.Scan()
		choice = scanner.Text()
		if err := scanner.Err(); err != nil {
			fmt.Println("Err reading input:", err)
		}

		choice = strings.TrimSpace(choice)

		if choice == "q" || choice == "quit" {
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
			fmt.Println("[X] Id: ", item.ID, "Item: ", item.Todo)
		} else {
			fmt.Println("[ ] Id: ", item.ID, "Item: ", item.Todo)
		}
	}
}

func addTask(scanner *bufio.Scanner, todo []TodoItem) []TodoItem {
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

	fmt.Println("Enter the id: ")
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
