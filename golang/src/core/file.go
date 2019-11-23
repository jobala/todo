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
	database := openDatabase()
	entry := createDbEntryFor(todoItem)
	fmt.Println(entry)

	defer close(database)
}

func createDbEntryFor(todoItem ToDo) string {
	return fmt.Sprintf("%d,%s,%s,%s",
		todoItem.ID, todoItem.Description, todoItem.CreatedAt, todoItem.UpdatedAt)
}

func openDatabase() *os.File {
	openedFile, err := os.Create("database.csv")
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
