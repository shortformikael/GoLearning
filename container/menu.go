package container

import "fmt"

type Menu struct {
	Items       *TreeGraph
	Cursor      *int
	Current     *TreeNode
	Breadcrumbs *LinkedList
}

func NewMenu(pItems *TreeGraph) Menu {
	var l *LinkedList = &LinkedList{}
	l.Append(pItems.Head.Value)
	return &Menu{
		Items:       pItems,
		Cursor:      0,
		Current:     pItems.Head,
		Breadcrumbs: l,
	}
}

func (m *Menu) PrintCli() {
	fmt.Println("")
	// Breadcrumbs
	for _, crumb := range m.Breadcrumbs.GetArray {
		fmt.Print("> ", crumb, " ")
	}
	fmt.Print("\n")

	for i := 0; i < len(m.current.Children); i++ {
		if i == m.Pointer {
			fmt.Println(" >[*]", m.current.Children[i])
		} else {
			fmt.Println("  [*]", m.current.Children[i])
		}
	}
}
