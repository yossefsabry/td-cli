// Author: Yossef sabry -> github.com/yossefsabry

// THIS IS  GO CMD TODO AND TASK MANAGER APP THAT HELP YOU TO MANAGE YOUR TASK AND TODO IN THE CSV FILE NOW YOU CAN ADD, UPDATE, DELETE, COMPLATE, SHOW, SHOW COMPLATE AND EXIT THE APP AND THE DATA SAVE IN THE CSV FILE

package main

import (
	"encoding/csv"
	"fmt"
	_ "io"
	"log"
	"os"
)

func main() {
	fileName := "data.csv"
	header := []string{"id", "Message", "Complate", "Deleted", "CreatedAt", "UpdateAt"}
	checkFile(fileName, header) // func for check for the file is exists or not

	PrintHelloAndHint() // for printing Hello and Hint 

	// starting for loop for app
	run := true
	for run {
// HINT:  ADDING THE OPEN FILE IN THE WHILE LOOP BECAUSE WHEN YOU A OPEN A FILE AND MAKE A READER FROM THE FILE WHEN YOU DO OPERATION ON THE FILE THE OPERATION SAVA AND THE FILE START FROM THE NEW OPERATION FOR EXAMPLE WHEN READING A CONTENT OF FILE THEN NEXT TIME MAKE ANOTHER OPERATION THE OPERATION START FROM THE END OF THE LAST OPERATION FOR THAT WE ADDING THE OPEN FILE IN THE FOR LOOP TO OPEN EVERY THE FILE EVERY TIME
		file, err := os.Open(fileName) // open a file
		reader := csv.NewReader(file)
		if err != nil {
			log.Fatal("error happend when open file: %v", err)
		}

		defer file.Close() // closing the file
		PrintOptions() // for printing the options for user
		fmt.Printf("-- Enter An Option: ")
		var userOption string
		_, err = fmt.Scanln(&userOption)
		if err != nil {
			fmt.Println("-- Error reading input:", err)
		}
		switch userOption {
		case "1", "add":
			err := addTodo(fileName)
			if err != nil {
				log.Fatal("error happend in deleting the todo: %v", err)
			}
		case "2", "update":
			updateTodo(reader)
		case "3", "delete":
			err := todoDelete(reader) // for deleting the todo
			if err != nil {
				log.Fatal("error happend in deleting the todo: %v", err)
			}
		case "4", "complate":
			err := ComplateTodo(reader)	
			if err != nil {
				log.Fatal("error happend in complate the todo: %v", err)
			}
		case "5", "show":
			err := showTodo(file, reader)
			if err != nil {
				log.Fatal("error happend in show the todo: %v", err)
			}
			// testReader(reader) // for testing porpouse
		case "6", "show complate":
			err := showComplateTodo(file, reader)
			if err != nil {
				log.Fatal("error happend in show the todo: %v", err)
			}
		case "q", "exit", "quit", "Q":
			run = false
			fmt.Printf("-- Exiting...")
			return
		default: 
			fmt.Printf("\n-- not vaild option enter [1..6], q --\n")
		}
	}

	return
}



// func testReader(reader *csv.Reader) {
// 	for {
// 		rec, err := reader.Read()
// 		if err == io.EOF {
// 			break
// 		}
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		// do something with read line
// 		fmt.Printf("%+v\n", rec)
// 	}
// }


