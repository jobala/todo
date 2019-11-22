package core

import (
	"os"
)

// Read todo items from file
func Read(todoItem ToDo) {

}

// Write todo items to files
func Write(todoItem ToDo) {
	fileDb := open("database.csv")

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
