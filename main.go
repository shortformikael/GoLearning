package main

import (
	"fmt"

	"github.com/shortformikael/GoLearning/container"
)

var list container.LinkedList
var status string = "START"

func main() {
	initFillList()
	start()
}

func start() {
	for {
		// Display
		fmt.Println("Hello, program start")
		// Get Input
		var choice string
		fmt.Print("> ")
		fmt.Scan(&choice)
		// Process
		if choice == "exit" {
			break
		}
		// Repeat
	}
}

func initFillList() {
	list.Append("First")
	list.Append("Second")
	list.Append("Third")
	list.Append("Fourth")
	list.Append("Fifth")
	list.Append("Last")
}
