package main

import (
	"fmt"

	"github.com/shortformikael/GoLearning/container"
)

var list container.LinkedList

func main() {
	fmt.Println("START")
	initFillList()

}

func initFillList() {
	list.Append("First")
	list.Append("Second")
	list.Append("Third")
	list.Append("Fourth")
	list.Append("Fifth")
	list.Append("Last")
}
