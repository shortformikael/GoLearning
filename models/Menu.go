package models

import (
	"fmt"
)

type Menu struct {
	Status MenuOption
	Items  [5]MenuOption
}

func (m Menu) Print() {
	for i := 0; i < len(m.Items); i++ {
		fmt.Printf("%d: %s \n", i, m.Items[i])
	}
}

func NewMenu() *Menu {
	return &Menu{
		Status: Start,
		Items: [5]MenuOption{
			Start,
			Add,
			Delete,
			View,
			"EXIT",
		},
	}
}
