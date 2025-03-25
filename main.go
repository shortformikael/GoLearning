package main

import (
	"fmt"

	"github.com/shortformikael/GoLearning/container"
)

var list container.LinkedList
var status string = "START"
var drawCh chan string = make(chan string)
var main_menu container.Menu

func main() {
	initFillList()
	start()
}

func start() {
	// Display
	go drawMenu()
	for {
		// Get Input
		var choice string
		fmt.Scan(&choice)
		// Process
		if choice == "exit" {
			drawCh <- "exit"
			break
		} else {
			drawCh <- "1"
		}
		// Repeat
	}
}

func drawMenu() {
	var current *container.Menu = &main_menu
	for {
		//Clear
		clearConsole()
		//Print Interface
		fmt.Println(">>> Shopping List <<<")
		current.Print()
		fmt.Print("> ")

		//Wait and Listen
		msg := <-drawCh
		if msg == "exit" { //Close on "exit" command
			break
		}
	}
}

func clearConsole() {
	fmt.Print("\033[H\033[2J")
}

func init() {
	initFillList()
	var mainList *container.LinkedList = &container.LinkedList{}
	mainList.Append("View")
	mainList.Append("Add")
	mainList.Append("Delete")
	mainList.Append("Exit")
	//arr := [4]string{"View", "Add", "Delete", "Exit"}
	//mainList.AddArray(arr)
	main_menu = *container.NewMenu("Main Menu", *mainList)

}

func initFillList() {
	list.Append("First")
	list.Append("Second")
	list.Append("Third")
	list.Append("Fourth")
	list.Append("Fifth")
	list.Append("Last")
}
