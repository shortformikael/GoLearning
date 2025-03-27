package app

import (
	"fmt"

	"github.com/shortformikael/GoLearning/container"
)

type Engine struct {
	Menu         *container.TreeGraph
	List         *container.LinkedList
	Running      bool
	current_menu *container.TreeNode
	commandCh    *chan string
}

func (e *Engine) Start() {
	fmt.Println("Starting Engine...")
	e.Running = true
	for e.Running {

		e.printTreeNode(e.current_menu)
		e.Running = false
	}
}

func (e *Engine) Init() {
	e.Running = false
	e.current_menu = e.Menu.Head
	e.commandCh = &chan //Command Channel, 
}

func (e *Engine) printTreeNode(node *container.TreeNode) {
	fmt.Println("==", node, "==")
	for i := 0; i < len(node.Children); i++ {
		fmt.Println("  [*]", node.Children[i])
	}
}

func keyboardListener() {

}
