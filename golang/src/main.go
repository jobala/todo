package main

import (
	"bufio"
	"core"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("ToDo Manager")
	fmt.Println("--------------------------------")

	for {
		fmt.Print("todo> ")
		text, _ := reader.ReadString('\n')

		switch strings.Replace(text, "\n", "", -1) {
		case "read":
			core.Read()
		case "add":
			addTodo()
		case "exit":
			os.Exit(0)
		default:
			fmt.Println("Unknown command")
		}
	}

	// core.Read()
}

func addTodo() {
	todoItem := core.ToDo{
		ID:          rand.Int(),
		Description: "A todo item",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	core.Save(todoItem)
}
