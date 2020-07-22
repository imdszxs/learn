package tree

import "fmt"

type Node struct {
	Value int
	Left *Node
	Right *Node
}

func (node Node) Print() {
	fmt.Println(node.Value)
}

func (node *Node) Traverse() {
	if(node == nil){
		return
	}
	node.Left.Traverse()
	node.Print()
	node.Right.Traverse()
}
