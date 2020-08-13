package tree

import "fmt"

type Node struct {
	Value int
	Left *Node
	Right *Node
}

func (node Node) Print() {
	fmt.Printf("%d", node.Value)
}

func (node *Node) Traverse() {
	node.TraverseFunc(func(n *Node) {
		n.Print()
	})
	fmt.Println()
}

func (node *Node) TraverseFunc(f func(*Node)) {
	if(node == nil){
		return
	}
	node.Left.TraverseFunc(f)
	f(node)
	node.Right.TraverseFunc(f)
}
