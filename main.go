package main

import (
	"fmt"

	"github.com/shortformikael/GoLearning/app"
	"github.com/shortformikael/GoLearning/container"
)

var App *app.Engine = &app.Engine{}

func main() {
	fmt.Println("Soft Reset Complete")
	varInit()
	start()
}

func varInit() {
	fmt.Println("Started Program Initialization...")
	App.Menu = &container.TreeGraph{Head: &container.TreeNode{}}
	App.List = &container.LinkedList{}

	App.Menu.Head.Value = "Main Menu"
	App.Menu.Head.AddChild(&container.TreeNode{Value: "View"})
	App.Menu.Head.AddChild(&container.TreeNode{Value: "Settings"})
	App.Menu.Head.AddChild(&container.TreeNode{Value: "Exit"})

	App.Init()
}

func start() {
	App.Start()
}
