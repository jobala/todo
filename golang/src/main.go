package main

import (
	"fmt"
	"core"
)

func main() {
	fmt.Println("vim-go")

	todoItem := core.ToDo{
		ID: 1,
		Description: "A todo item",
		CreatedAt: "Today",
		UpdatedAt: "Yesterday",
	}

	core.Write(todoItem)
}
