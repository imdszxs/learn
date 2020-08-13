package main

import (
	"fmt"
	"learn/tree"
)

type myNode struct {
	node *tree.Node
}

func (node *myNode) PostOrder(){
	if node == nil || node.node == nil {
		return
	}
	left := myNode{node.node.Left}
	left.PostOrder()
	right := myNode{node.node.Right}
	right.PostOrder()
	fmt.Println(node.node.Value)
}

func main() {
	var root tree.Node
	root = tree.Node{Value:3}
	root.Left = &tree.Node{}
	root.Right = &tree.Node{5, nil, nil}
	root.Left.Right = &tree.Node{Value:2}
	root.Right.Left = &tree.Node{Value:4}
	root.Traverse()
	//fmt.Println()
	//myroot := myNode{&root}
	//myroot.PostOrder()

	nodeCount := 0

	root.TraverseFunc(func(node *tree.Node) {
		nodeCount++
	})

	fmt.Println(nodeCount)
}