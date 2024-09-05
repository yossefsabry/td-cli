package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strings"
	"text/tabwriter"
)


type todo struct {
	ID string
	message string
	complate bool
	delete bool
	CreateAt string
	UpdateAt string
}



// now show todo for the user if found any todo
func showTodo(file *os.File) {
	// starting creatting a csv reader
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Printf("error happend in reader for file: %v\n",err)
		return
	}

	// use tabwriter to print the table for info
	writer := tabwriter.NewWriter(os.Stdout, 0,0,6, ' ' , tabwriter.Debug)
	println("\n\n")
	fmt.Fprintln(writer, "id\tMessage\tComplate\tDeleted\tCreatedAt\tUpdateAt\t")

	for i, record :=range records {
		// skip the header now
		if i == 0 {
			continue
		}


		message := getLastNWords(record[1], 20)
		fmt.Fprintf(writer, "%s\t%s\t%s\t%s\t%s\t%s\t\n", record[0], message, record[2], record[3], record[4], record[5])
	println("")
	}
	writer.Flush()
}

/**
	@description get the last word for message and check if words  is less then 20 or not to format the message for user
*/
func getLastNWords(message string,n int)string {
	words := strings.Fields(message) // split the message into words
	if len(words) > n {
		return strings.Join(words[len(words)-n:], " ") // get the last N Words
	}
	return message // return the message as it is if it's shourter
}


// @starting adding todo for user
func addTodo(file *os.File) {
		
}



