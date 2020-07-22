package main

import (
	"fmt"
	"learn/queue"
)

func main() {
	list := queue.Queue{1}
	list.Push(2)
	list.Push(3)
	fmt.Println(list)
	fmt.Println(list.Pop())
	fmt.Println(list.Pop())
	fmt.Println(list)
}
