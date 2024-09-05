package main

import (
	"fmt"
	"log"
	"os"
)


func main() {


	fileName := "data.csv"
	header := []string{"id", "Message", "Complate", "Deleted", "CreatedAt", "UpdateAt"}
	checkFile(fileName, header) // func for check for the file is exists or not

	PrintHelloAndHint() // for printing Hello and Hint 

	file, err := os.Open(fileName) // open a file
	if err != nil {
		log.Fatal("error happend when open file: %v", err)
	}
	defer file.Close() // closing the file

	// starting for loop for app
	run := true
	for run {
		PrintOptions() // for printing the options for user
		fmt.Printf("-- Enter An Option: ")
		var userOption string
		_, err := fmt.Scanln(&userOption)
		if err != nil {
			fmt.Println("-- Error reading input:", err)
			return
		}
		switch userOption {
		case "1":
			continue
		case "2":
			continue
		case "3":
			continue
		case "4":
			continue
		case "5":
			showTodo(file)
		case "6":
			continue
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
