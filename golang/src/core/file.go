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
	err := createDbEntryFor(todoItem, database)

	if err != nil {
		handleError(err)
	}

	fmt.Println("TODO item saved successfully")
	defer close(database)
}

func createDbEntryFor(todoItem ToDo, database *os.File) error {
	entry := fmt.Sprintf("%d,%s,%s,%s\n",
		todoItem.ID, todoItem.Description, todoItem.CreatedAt, todoItem.UpdatedAt)

	_, err := database.WriteString(entry)

	return err
}

func openDatabase() *os.File {
	openedFile, err := os.OpenFile("database.csv", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0777)
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
