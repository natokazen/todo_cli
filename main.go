package main

import (
	"fmt"
)

type TodoItem struct {
	ID          int
	Todo        string
	IsCompleted bool
}

func main() {
	todo := []TodoItem{}

	var choice, input string

	for {
		fmt.Println("Enter a command: 'add' , 'list', 'complete' or 'q' / 'quit' :")
		_, err := fmt.Scanln(&choice)
		if err != nil {
			fmt.Print("Please enter a valid choice")
		}

		if choice == "q" || choice == "quit" {
			break
		} else if choice == "list" {
			for _, item := range todo {
				if item.IsCompleted {
					fmt.Println("[X] Id: ", item.ID, "Item: ", item.Todo)
				} else {
					fmt.Println("[ ] Id: ", item.ID, "Item: ", item.Todo)
				}
			}
		} else if choice == "complete" {

			var targetID int

			fmt.Println("Enter the id: ")
			_, err := fmt.Scanln(&targetID)
			if err != nil {
				fmt.Print("Please enter a valid id")
			}

			for index, item := range todo {
				if len(todo) == 0 {
					fmt.Println("To items to complete yet!")
				}

				if item.ID == targetID {
					todo[index].IsCompleted = true
				}
			}

		} else if choice == "add" {

			_, err := fmt.Scanln(&input)
			if err != nil {
				fmt.Print("")
			}

			newTodo := TodoItem{
				ID:          len(todo) + 1,
				Todo:        input,
				IsCompleted: false,
			}
			todo = append(todo, newTodo)

		} else {
			fmt.Println("Not a command")
		}

	}

	fmt.Println(todo)
}
