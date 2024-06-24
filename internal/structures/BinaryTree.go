package structures

import (
	"fmt"
)

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

type Tree interface {
	Insert(value int)
	Search(value int) bool
	InOrderTraversal(f func(int))
	PrintTree()
}

func (n *Node) Insert(value int) {
	if value <= n.Value {
		if n.Left == nil {
			n.Left = &Node{Value: value}
		} else {
			n.Left.Insert(value)
		}
	} else {
		if n.Right == nil {
			n.Right = &Node{Value: value}
		} else {
			n.Right.Insert(value)
		}
	}
}

func (n *Node) Search(value int) bool {
	if n == nil {
		return false
	}
	if value == n.Value {
		return true
	} else if value < n.Value {
		return n.Left.Search(value)
	} else {
		return n.Right.Search(value)
	}
}

func (n *Node) InOrderTraversal(f func(int)) {
	if n != nil {
		n.Left.InOrderTraversal(f)
		f(n.Value)
		n.Right.InOrderTraversal(f)
	}
}

func (n *Node) PrintTree(level int) {
	if n == nil {
		return
	}
	if n.Right != nil {
		n.Right.PrintTree(level + 1)
	}
	fmt.Printf("%*s%d\n", level*4, "", n.Value)
	if n.Left != nil {
		n.Left.PrintTree(level + 1)
	}
}
