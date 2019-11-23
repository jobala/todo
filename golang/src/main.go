package main

import (
	"core"
	"fmt"
	"time"
)

func main() {
	fmt.Println("vim-go")

	todoItem := core.ToDo{
		ID:          1,
		Description: "A todo item",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	core.Write(todoItem)
}
