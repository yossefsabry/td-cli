
// MAIN OPERATION FOR THE TODO APP LIKE ADD, DELETE, COMPLATE, SHOW, UPDATEA AND SHOW COMPLATE TODO AND SHOW ALL TODO ALL THE MAIN OPERATION FOR THE TODO APP AND HELPER FUNCTION TO HELP THE MAIN OPERATION TO WORK CORRECTLY AND TO MAKE THE CODE MORE READABLE AND CLEAN USING THE HELPER FUNCTION TO MAKE THE CODE MORE READABLE AND CLEAN AND TO MAKE THE MAIN OPERATION MORE


/** WORK 
	- the id for the todo will be auto increment
	- try adding better structure for the todo to make the code more readable
*/

package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"
	"text/tabwriter"
	"time"
	"errors"
)


type todo struct {
	ID string
	Message string
	Complate bool
	Delete bool
	CreateAt string
	UpdateAt string
}

// ##########################
// ##########################
// ### start helper func ####
// ##########################
// ##########################

// Helper function to find the index of a column by name param : headers, columnName
func findIndex(headers []string, columnName string) int {
    for i, header := range headers {
        if header == columnName {
            return i
        }
    }
    return -1
}

// for over wirte a full with the new changes { param : fileName, records }
func overWriteFile(fileName string, records [][]string) error {
	// after deleting the record now write the changes in the file
	outFile, err := os.Create(fileName)
	if err != nil{
		return err
	}
	defer outFile.Close()
	// creattting a new writer for the file
	writer := csv.NewWriter(outFile)
	err = writer.WriteAll(records)
	if err != nil {
		return err
	}
	return nil
}

// @description get the last word for message and check if words  is less then 20 or not to format the message for user 
func getLastNWords(message string,n int)string {
	words := strings.Fields(message) // split the message into words
	if len(words) > n {
		return strings.Join(words[len(words)-n:], " ") // get the last N Words
	}
	return message // return the message as it is if it's shourter
}


// **************************
// **************************
// **** end helper func *****
// **************************
// **************************


// ##########################
// ##########################
// ### start main operation ##
// ##########################
// ##########################


// now show todo for the user if found any todo param : file, reader
func showTodo(file *os.File ,reader *csv.Reader)error {
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Printf("error happend in reader for file: %v\n",err)
		return err
	}
	// use tabwriter to print the table for info
	writer := tabwriter.NewWriter(os.Stdout, 0,0,6, ' ' , tabwriter.Debug)
	// how to setting the title in center of the table
	var totatTodo int = 0
	fmt.Printf("\t\t-------------------------------------Todo List-------------------------------------\n")
	fmt.Fprintln(writer, " - id\tMessage\tComplate\tDeleted\tCreatedAt\tUpdateAt\t")

	for i, record :=range records {
		// skip the header now
		if i == 0 {
			continue
		}
		if record[3] == "true" || record[2] == "true" {
			continue
		}
		message := getLastNWords(record[1], 20)
		fmt.Fprintf(writer, " - %s\t%s\t%s\t%s\t%s\t%s\t\n",
			record[0], message, record[2], record[3], record[4], record[5])
		totatTodo++
		println("")
	}
	fmt.Printf("Total Todos: %d\n",totatTodo)
	writer.Flush()
	return nil
}



// @starting adding todo for user { param : filename }
func addTodo(fileName string)error {
	// open the file for read and write
	var todoMessage string	
	fmt.Printf("Enter the todo: ")
	fmt.Scanln(&todoMessage)
	if (todoMessage == "") {
		return errors.New("Message can't be empty")
	}
	FormatTime := time.Now().Format("2006-01-02 15:04:05")
	Todo := todo{
		ID: "1", 
		Message: todoMessage,
		Complate: false,
		Delete: false,
		CreateAt: FormatTime,
		UpdateAt: FormatTime,
	}
	// Convert Todo to a slice of strings to adding to the file
		record := []string{
		Todo.ID,
		Todo.Message,
		strconv.FormatBool(Todo.Complate),
		strconv.FormatBool(Todo.Delete),
		Todo.CreateAt,
		Todo.UpdateAt,
	}
	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil{
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush() // Ensure all data is written
	// Write the new record to the file
	if err := writer.Write(record); err != nil {
		return err
	}
	fmt.Printf("-- Todo added successfully --")
	return nil
}

// starting deteting the todo param : reader
func todoDelete(reader *csv.Reader)error {
	fmt.Printf("Enter the id of the todo you want to delete: ")
	var id string
	fmt.Scanln(&id)
	records, err := reader.ReadAll()
	if err != nil {
		return err
	}
	// Find the index of the columns
	headers := records[0]
	idIndex := findIndex(headers, "id") // get the index of id
	deletedIndex := findIndex(headers, "Deleted") // get the index of deleted column from headers
	FormatTime := time.Now().Format("2006-01-02 15:04:05")
	for i, record := range records {
		if i == 0 {
			continue
		}
		if record[idIndex] == id && record[deletedIndex] == "false" {
			record[deletedIndex] = "true" // update the delete column
			record[5] = FormatTime // update the time
		}
	}

	// over write file for new data
	err = overWriteFile("data.csv", records)
	if err != nil{
		return err
	}
	fmt.Printf("-- Todo deleted successfully --\n")
	return nil
}

// adding complate todo param : reader
func ComplateTodo(reader *csv.Reader)error {
	fmt.Printf("Enter the id of the todo you want to complate: ")
	var id string
	fmt.Scanln(&id)
	records, err := reader.ReadAll()
	if err != nil {
		return err
	}
	// Find the index of the columns
	headers := records[0]
	idIndex := findIndex(headers, "id") // get the index of id
	ComplateIndex := findIndex(headers, "Complate") // get the index of deleted column from headers
	FormatTime := time.Now().Format("2006-01-02 15:04:05")
	for i, record := range records {
		if i == 0 {
			continue
		}
		if record[idIndex] == id && record[ComplateIndex] == "false" {
			record[ComplateIndex] = "true" // update the complate column
			record[5] = FormatTime // update the time
		}
	}
	// over write the file with new data
	err = overWriteFile("data.csv", records)
	if err != nil{
		return err
	}
	fmt.Printf("-- Complate Todo successfully --\n")
	return nil
}

// show complate todo param : file, reader
func showComplateTodo(file *os.File ,reader *csv.Reader)error {
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Printf("error happend in reader for file: %v\n",err)
		return err
	}
	// use tabwriter to print the table for info
	writer := tabwriter.NewWriter(os.Stdout, 0,0,6, ' ' , tabwriter.Debug)
	// how to setting the title in center of the table
	fmt.Fprintln(writer, " - id\tMessage\tComplate\tDeleted\tCreatedAt\tUpdateAt\t")

	for i, record :=range records {
		// skip the header now
		if i == 0 {
			continue
		}
		if record[3] == "false" && record[2] == "true" {
			message := getLastNWords(record[1], 20)
			fmt.Fprintf(writer, " - %s\t%s\t%s\t%s\t%s\t%s\t\n",
				record[0], message, record[2], record[3], record[4], record[5],
				)
			println("")
		}else {
			continue
		}
	}
	writer.Flush()
	return nil
}

// update todo { param : reader }
func updateTodo(reader *csv.Reader)error {
	var id string
	var newMessage string
	fmt.Printf("Enter the id of the todo you want to update: ")
	fmt.Scanln(&id)
	fmt.Printf("Enter the new message for the todo: ")
	fmt.Scanln(&newMessage)

	records, err := reader.ReadAll()
	if err != nil {
		return err
	}
	// Find the index of the columns
	headers := records[0]
	idIndex := findIndex(headers, "id") // get the index of id
	targetIndex := findIndex(headers, "Message") // get the index of deleted column from headers
	FormatTime := time.Now().Format("2006-01-02 15:04:05")
	for i, record := range records {
		if i == 0 {
			continue
		}
		if record[idIndex] == id {
			record[targetIndex] = newMessage // update with the new massage
			record[5] = FormatTime // update the time
		}
	}

	// over write the file with new data
	err = overWriteFile("data.csv", records)
	if err != nil{
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
