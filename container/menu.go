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

	for i := 0; i < len(m.Current.Children); i++ {
		if i == m.Pointer {
			fmt.Println(" >[*]", m.Current.Children[i])
		} else {
			fmt.Println("  [*]", m.Current.Children[i])
		}
	}
}

func (m *Menu) Next() {
	if m.Pointer+1 >= len(m.Current.Children) {
		m.Pointer = 0
	} else {
		m.Pointer = m.Pointer + 1
	}
}

func (m *Menu) Previous() {
	if m.Pointer-1 < 0 {
		m.Pointer = len(m.Current.Children)
	} else {
		m.Pointer = m.Pointer - 1
	}
}

func (m *Menu) Select() {
	m.Current = m.Current.Children[m.Pointer]
	m.Breadcrumbs.Append(m.Current.Value) //????
	m.Pointer = 0
}
