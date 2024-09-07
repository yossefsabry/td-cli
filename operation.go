// MAIN OPERATION FOR THE TODO APP LIKE ADD, DELETE, COMPLETE, SHOW,
// UPDATE AND SHOW COMPLETE TODO AND SHOW ALL TODO ALL THE MAIN OPERATION FOR THE
// TODO APP AND HELPER FUNCTION TO HELP THE MAIN OPERATION TO WORK CORRECTLY AND TO MAKE THE
// CODE MORE READABLE AND CLEAN USING THE HELPER FUNCTION TO MAKE THE CODE MORE READABLE AND
// CLEAN AND TO MAKE THE MAIN OPERATION MORE

/** WORK
- try adding better structure for the todo to make the code more readable
*/

package main

import (
	"encoding/csv"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"log"
	"os"
	"strconv"
	"strings"
	"text/tabwriter"
	"time"
	"unicode"
)

type todo struct {
	ID       string
	Message  string
	Complete bool
	Delete   bool
	CreateAt string
	UpdateAt string
}

// ##########################
// ### start helper func ####
// ##########################

// starting a func for extract 6 letters from the uuid.New()
func extractLetters(s string) string {
	letters := make([]rune, 0, 5)
	for _, ch := range s {
		if unicode.IsLetter(ch) {
			letters = append(letters, ch)
			if len(letters) > 5 {
				break
			}
		}
	}
	return string(letters)
}

// func for close a file
func closeFile(file *os.File) {
	err := file.Close()
	if err != nil {
		log.Fatal(err)
	}
}

// Helper function to find the index of a column by name param : headers, columnName
func findIndex(headers []string, columnName string) int {
	for i, header := range headers {
		if header == columnName {
			return i
		}
	}
	return -1
}

// for overwrite a full of the new changes { param : fileName, records }
func overWriteFile(fileName string, records [][]string) error {
	// after deleting the record now write the changes in the file
	outFile, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer closeFile(outFile)
	// creating a new writer for the file
	writer := csv.NewWriter(outFile)
	err = writer.WriteAll(records)
	if err != nil {
		return err
	}
	return nil
}

// @description get the last word for message and check if words  is less than 20 or not to format the message for user
func getLastNWords(message string, n int) string {
	words := strings.Fields(message) // split the message into words
	if len(words) > n {
		return strings.Join(words[len(words)-n:], " ") // get the last N Words
	}
	return message // return the message as it is if it's shorter
}

// **************************
// **** end helper func *****
// **************************

// ##########################
// ### start main operation ##
// ##########################

// now show todo for the user if found any todo param : file, reader
func showTodo(reader *csv.Reader) error {
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Printf("error happend in reader for file: %v\n", err)
		return err
	}
	// use tab writer to print the table for info
	writer := tabwriter.NewWriter(os.Stdout, 0, 0, 6, ' ', tabwriter.Debug)
	// how to set the title in center of the table
	totalTodo := 0
	fmt.Printf("\t\t-------------------------------------Todo List-------------------------------------\n")
	_, err = fmt.Fprintln(writer, " - id\tMessage\tComplete\tDeleted\tCreatedAt\tUpdateAt\t")
	if err != nil {
		return err
	}

	for i, record := range records {
		// skip the header now
		if i == 0 {
			continue
		}
		if record[3] == "true" || record[2] == "true" {
			continue
		}
		message := getLastNWords(record[1], 20)
		_, err2 := fmt.Fprintf(writer, " - %s\t%s\t%s\t%s\t%s\t%s\t\n",
			record[0], message, record[2], record[3], record[4], record[5])
		if err2 != nil {
			return err2
		}
		totalTodo++
		println("")
	}
	fmt.Printf("Total Todos: %d\n", totalTodo)
	err = writer.Flush()
	if err != nil {
		return err
	}
	return nil
}

// @starting adding todo for user { param : filename }
func addTodo(fileName string) error {
	// open the file for read and write
	var todoMessage string
	fmt.Printf("Enter the todo: ")
	_, err2 := fmt.Scanln(&todoMessage)
	if err2 != nil {
		return err2
	}
	if todoMessage == "" {
		return errors.New("message can't be empty")
	}
	FormatTime := time.Now().Format("2006-01-02 15:04:05")
	NewId := extractLetters(uuid.New().String())
	var Todo = todo{
		ID:       NewId,
		Message:  todoMessage,
		Complete: false,
		Delete:   false,
		CreateAt: FormatTime,
		UpdateAt: FormatTime,
	}
	// Convert Todo to a slice of strings to adding to the file
	var record = []string{
		Todo.ID,
		Todo.Message,
		strconv.FormatBool(Todo.Complete),
		strconv.FormatBool(Todo.Delete),
		Todo.CreateAt,
		Todo.UpdateAt,
	}
	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer closeFile(file)

	writer := csv.NewWriter(file)
	defer writer.Flush() // Ensure all data is written
	// Write the new record to the file
	if err := writer.Write(record); err != nil {
		return err
	}
	fmt.Printf("-- Todo added successfully --\n")
	return nil
}

// starting detesting the todo param : reader
func todoDelete(reader *csv.Reader) error {
	fmt.Printf("Enter the id of the todo you want to delete: ")
	var id string
	_, err := fmt.Scanln(&id)
	if err != nil {
		return err
	}
	records, err2 := reader.ReadAll()
	if err2 != nil {
		return err2
	}
	// Find the index of the columns
	headers := records[0]
	idIndex := findIndex(headers, "id")           // get the index of id
	deletedIndex := findIndex(headers, "Deleted") // get the index of deleted column from headers
	FormatTime := time.Now().Format("2006-01-02 15:04:05")
	for i, record := range records {
		if i == 0 {
			continue
		}
		if record[idIndex] == id && record[deletedIndex] == "false" {
			record[deletedIndex] = "true" // update the delete column
			record[5] = FormatTime        // update the time
		}
	}

	// overwrite file for new data
	err = overWriteFile("data.csv", records)
	if err != nil {
		return err
	}
	fmt.Printf("-- Todo deleted successfully --\n")
	return nil
}

// CompleteTodo adding complete todo param : reader
func CompleteTodo(reader *csv.Reader) error {
	fmt.Printf("Enter the id of the todo you want to complate: ")
	var id string
	_, err2 := fmt.Scanln(&id)
	if err2 != nil {
		return err2
	}
	records, err := reader.ReadAll()
	if err != nil {
		return err
	}
	// Find the index of the columns
	headers := records[0]
	idIndex := findIndex(headers, "id")             // get the index of id
	CompleteIndex := findIndex(headers, "Complete") // get the index of deleted column from headers
	FormatTime := time.Now().Format("2006-01-02 15:04:05")
	for i, record := range records {
		if i == 0 {
			continue
		}
		if record[idIndex] == id && record[CompleteIndex] == "false" {
			record[CompleteIndex] = "true" // update the complete column
			record[5] = FormatTime         // update the time
		}
	}

	// overwrite the file with new data
	err = overWriteFile("data.csv", records)
	if err != nil {
		return err
	}
	fmt.Printf("-- Complete Todo successfully --\n")
	return nil
}

// show complete todo param : file, reader
func showCompleteTodo(reader *csv.Reader) error {
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Printf("error happend in reader for file: %v\n", err)
		return err
	}
	// use tab writer to print the table for info
	writer := tabwriter.NewWriter(os.Stdout, 0, 0, 6, ' ', tabwriter.Debug)
	// how to set the title in center of the table
	_, err = fmt.Fprintln(writer, " - id\tMessage\tComplete\tDeleted\tCreatedAt\tUpdateAt\t")
	if err != nil {
		return err
	}

	for i, record := range records {
		// skip the header now
		if i == 0 {
			continue
		}
		if record[3] == "false" && record[2] == "true" {
			message := getLastNWords(record[1], 20)
			_, err := fmt.Fprintf(writer, " - %s\t%s\t%s\t%s\t%s\t%s\t\n",
				record[0], message, record[2], record[3], record[4], record[5],
			)
			if err != nil {
				return err
			}
			println("")
		} else {
			continue
		}
	}
	err = writer.Flush()
	if err != nil {
		return err
	}
	return nil
}

// update todo { param : reader }
func updateTodo(reader *csv.Reader) error {
	var id string
	var newMessage string
	fmt.Printf("Enter the id of the todo you want to update: ")
	_, err2 := fmt.Scanln(&id)
	if err2 != nil {
		return err2
	}
	fmt.Printf("Enter the new message for the todo: ")
	_, err3 := fmt.Scanln(&newMessage)
	if err3 != nil {
		return err3
	}

	records, err := reader.ReadAll()
	if err != nil {
		return err
	}
	// Find the index of the columns
	headers := records[0]
	idIndex := findIndex(headers, "id")          // get the index of id
	targetIndex := findIndex(headers, "Message") // get the index of deleted column from headers
	FormatTime := time.Now().Format("2006-01-02 15:04:05")
	for i, record := range records {
		if i == 0 {
			continue
		}
		if record[idIndex] == id {
			record[targetIndex] = newMessage // update with the new massage
			record[5] = FormatTime           // update the time
		}
	}

	// overwrite the file with new data
	err = overWriteFile("data.csv", records)
	if err != nil {
		return err
	}
	fmt.Printf("-- Todo update successfully --\n")
	return nil
}

// **************************
// **************************
// *** end main operation ***
// **************************
// **************************
