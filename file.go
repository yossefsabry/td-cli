package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

// func for check for the file is exsits or not
func checkFile(fileName string, header []string) {
	_, err := os.Stat(fileName)
	if err != nil  {
		fmt.Printf("the file not found starting creating file: %v\n", fileName)

		// starting creating the file
		file, err := os.Create(fileName)
		defer file.Close() // closing the file
		if err != nil {
			log.Fatal("error happend: %v\n", err)
		}

		// starting adding the header for the file
		writer := csv.NewWriter(file)
		err = writer.Write(header)
		if err != nil {
			log.Fatal("error in writing the header for the file: %v\n", err)
		}
		writer.Flush()
        if err := writer.Error(); err != nil {
            fmt.Println("Error flushing writer:", err)
            return
        }
        fmt.Println("File created and header written.")
	}

}
