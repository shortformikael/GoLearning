package container

import (
	"errors"
	"fmt"
)

type TreeGraph struct {
	head *TreeNode
}

func (g *TreeGraph) Search(s any) (*TreeNode, error) {
	if g.head == nil {
		return nil, errors.New("EMPTY TREE")
	} else if g.head.children != nil {
		if g.head.value == s {
			return g.head, nil
		} else {
			return nil, errors.New("VALUE DOES NOT EXIST IN TREE")
		}
	}

	current := &g.head.children[0]
	result := g.dfsSearch(*current, &s)

	if result != nil {
		return result, nil
	} else {
		return nil, errors.New("VALUE DOES NOT EXIST IN TREE")
	}

}

func (g *TreeGraph) dfsSearch(current *TreeNode, s *any) *TreeNode {
	for _, child := range current.children {
		if child.value == s {
			return child
		} else if child.children != nil {
			return g.dfsSearch(child, s)
		}
	}
	return nil
}

func (g *TreeNode) DfsPrintTraversal() {
	fmt.Println("Node Value: ", g.value)
	for i := 0; i < len(g.children); i++ {
		g.DfsPrintTraversal()
	}
}

type TreeNode struct {
	value    any
	children []*TreeNode
}

func (n *TreeNode) AddChild(child *TreeNode) {
	n.children = append(n.children, child)
}

func (n *TreeNode) Get(i int) (*TreeNode, error) {
	if i < 0 || len(n.children) < i {
		return nil, errors.New("OUT OF BOUNDS")
	}
	return n.children[i], nil
}

func (n *TreeNode) SearchChildren(child any) (*TreeNode, error) {
	for i := 0; i < len(n.children); i++ {
		if child == n.children[i].value {
			return n.children[i], nil
		}
	}
	return nil, errors.New("DOES NOT EXIST IN ARRAY")
}
