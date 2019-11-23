package core

import (
	"fmt"
	"os"
)

// Read todo items from file
func Read(todoItem ToDo) {

}

// Write todo items to files
func Write(todoItem ToDo) {
	fileDb := open("database.csv")
	entry := fmt.Sprintf("%d,%s,%s,%s",
		todoItem.ID, todoItem.Description, todoItem.CreatedAt, todoItem.UpdatedAt)
	fmt.Println(entry)

	defer close(fileDb)
}

func open(file string) *os.File {
	openedFile, err := os.Create(file)
	handleError(err)

	return openedFile
}

func close(file *os.File) {
	err := file.Close()
	handleError(err)
}

func handleError(err error) {
	if err != nil {
		panic(err)
	}
}
