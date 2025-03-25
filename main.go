package main

import (
	"fmt"
	"time"

	"github.com/shortformikael/GoLearning/container"
)

var list *container.LinkedList = &container.LinkedList{}
var status string = "START"
var drawCh chan string = make(chan string)
var runCh chan bool = make(chan bool)
var main_menu container.Menu
var current_menu *container.Menu = &main_menu
var running bool

func main() {
	start()
}

func start() {
	// Display
	running = true
	//go runListener()
	go drawMenu()
	var choice string
	for running {
		// Get Input
		fmt.Scan(&choice)
		// Process
		ProcessChoice(choice)
	}

}

func ProcessChoice(choice string) {

	var cond interface{}
	var err error
	cond, err = current_menu.Items.Get(choice)

	if err != nil {
		if current_menu.Name != "Add" {
			fmt.Println("Error:", err)
		} else {
			list.Append(choice)
		}
	} else {
		switch c := cond.(type) {
		case string:
			if current_menu.Name == "Delete" {
				//list.Delete(c)
			}
		case *container.Menu:
			current_menu = c
		case *container.MenuItem:
			if c.Name == "Exit" {
				if current_menu.Parent != nil {
					current_menu = current_menu.Parent
				} else {
					running = false
				}
			}
		}
	}

	drawCh <- "draw"
}

// NOT FINISHED
func filterChoice(choice interface{}) (interface{}, error) {
	r, _ := choice.(string)
	return r, nil
}

func drawMenu() {
	for running {
		//Clear
		clearConsole()
		//Print Interface
		fmt.Println(">>> Shopping List <<<")
		current_menu.Print()
		fmt.Print("> ")

		//Wait and Listen
		command := <-drawCh
		fmt.Println(command)
		time.Sleep(500 * time.Millisecond)
	}
}

func clearConsole() {
	fmt.Print("\033[H\033[2J")
}

func init() {
	initFillList(list)

	//Menu Init
	var viewList *container.LinkedList = &container.LinkedList{}
	viewList.Append(&container.MenuItem{Name: "Exit"})
	viewList.Append(list)
	view_menu := container.NewMenu("View", *viewList)
	view_menu.Parent = &main_menu

	var addList *container.LinkedList = &container.LinkedList{}
	addList.Append(&container.MenuItem{Name: "Exit"})
	addList.Append(list)
	add_menu := container.NewMenu("Add", *addList)
	add_menu.Parent = &main_menu

	var deleteList *container.LinkedList = &container.LinkedList{}
	deleteList.Append(&container.MenuItem{Name: "Exit"})
	deleteList.Append(list)
	delete_menu := container.NewMenu("Delete", *deleteList)
	delete_menu.Parent = &main_menu

	var mainList *container.LinkedList = &container.LinkedList{}
	mainList.Append(view_menu)
	mainList.Append(add_menu)
	mainList.Append(delete_menu)
	mainList.Append(&container.MenuItem{Name: "Exit"})
	//arr := [4]string{"View", "Add", "Delete", "Exit"}
	//mainList.AddArray(arr)
	main_menu = *container.NewMenu("Main Menu", *mainList)
}

func initFillList(l *container.LinkedList) {
	l.Append("First")
	l.Append("Second")
	l.Append("Third")
	l.Append("Fourth")
	l.Append("Fifth")
	l.Append("Last")
}
