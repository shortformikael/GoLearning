package models

import (
	"fmt"
)

type menu struct {
	Status menuOption
	Items  [5]menuOption
}

func (m menu) Print() {
	for i := 0; i < len(m.Items); i++ {
		fmt.Printf("%d: %s \n", i, m.Items[i])
	}
}

func NewMenu() *menu {
	return &menu{
		Status: Start,
		Items: [5]menuOption{
			Start,
			Add,
			Delete,
			View,
			"EXIT",
		},
	}
}
