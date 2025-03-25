package container

import "fmt"

type Menu struct {
	Name   string
	Items  LinkedList
	Parent *Menu
}

func NewMenu(s string, items LinkedList) *Menu {
	return &Menu{
		Name:  s,
		Items: items,
	}
}

func (m *Menu) Print() {
	fmt.Println("==", m.Name, "==")
	m.Items.PrintList()
}
