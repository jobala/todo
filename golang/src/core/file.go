package core

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Read todo items from file
func Read() {
	file, err := os.Open("database.csv")
	handleError(err)

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		text := scanner.Text()
		withoutCommas := strings.Replace(text, ",", " : ", -1)
		fmt.Println(withoutCommas)
	}

	defer close(file)
}

// Save todo items to files
func Save(todoItem ToDo) {
	database := openDatabase()
	err := createDbEntryFor(todoItem, database)
	handleError(err)

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
