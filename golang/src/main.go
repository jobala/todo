package main

import (
	"core"
	"time"
)

func main() {

	todoItem := core.ToDo{
		ID:          1,
		Description: "A todo item",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	core.Save(todoItem)
}
