package main

import (
	"fmt"
)

func PrintOptions() {
	fmt.Printf("\n 1- for creating note... \n")
	fmt.Printf(" 2- for update note... \n")
	fmt.Printf(" 3- for delete note... \n")
	fmt.Printf(" 4- for check|complate note... \n")
	fmt.Printf(" 5- for show notes... \n")
	fmt.Printf(" 6- for show complate notes... \n")
	fmt.Printf(" q- for exists the program... \n")
}

func PrintHelloAndHint() {
	asciiArtText := `

 ▄▄   ▄▄ ▄▄▄▄▄▄▄ ▄▄▄     ▄▄▄     ▄▄▄▄▄▄▄ 
 █  █ █  █       █   █   █   █   █       █
 █  █▄█  █   ▄   █   █   █   █   █   ▄   █
 █       █  █ █  █   █   █   █   █  █ █  █
 █   ▄   █  █▄█  █   █▄▄▄█   █▄▄▄█  █▄█  █
 █  █ █  █       █       █       █       █
 █▄▄█ █▄▄█▄▄▄▄▄▄▄█▄▄▄▄▄▄▄█▄▄▄▄▄▄▄█▄▄▄▄▄▄▄█
`

	fmt.Printf(" %v\n", asciiArtText)
	fmt.Printf("-- Hint in todo go todo --\n")
	fmt.Printf(" - you can adding note and update the note and check if complete or not...\n")
	fmt.Printf(" - and delete and can check for the old notes that complate and delete...\n")
	fmt.Printf(" - using words: \n")
	fmt.Printf(" - (show): show todos, (complate): adding complate todos, (add): add new todo\n")
	fmt.Printf(" - (show complate): show complate todo, (delete): delete todos, (update): update todo\n")
	fmt.Printf(" - (q, quit, exit): for exists the program\n")
}
