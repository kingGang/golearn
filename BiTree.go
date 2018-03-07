package main

import (
	// "encoding/json"
	"fmt"
	"golearn/mystack"
)

type Node struct {
	Data  interface{}
	Left  *Node
	Right *Node
}
type BiTree Node

//先序遍历
func PreOrder(root *Node) {
	// fmt.Printf("root=%+v\n", root)
	stack := mystack.NewStack()
	stack.Push(root)
	for !stack.Empty() {
		val, _ := stack.Pop()

		result, ok := val.(interface{})
		if ok {
			node, _ := result.(*Node)
			fmt.Println(node.Data)
			if node.Right != nil {
				stack.Push(node.Right)
			}
			if node.Left != nil {
				stack.Push(node.Left)
			}
		}
	}
}
func InOrder(root *Node) {
	stack := mystack.NewStack()
	for !stack.Empty() || root != nil {
		for root != nil {
			stack.Push(root)
			root = root.Left
		}
		val, _ := stack.Pop()
		interVal, ok := val.(interface{})
		if ok {
			node, _ := interVal.(*Node)
			fmt.Println(node.Data)
		}
		root = root.Right
	}
}

func main() {
	a := Node{}
	a.Data = "a"
	b := Node{}
	b.Data = "b"
	c := Node{}
	c.Data = "c"
	d := Node{}
	d.Data = "d"
	e := Node{}
	e.Data = "e"
	f := Node{}
	f.Data = "f"
	a.Left = &b
	a.Right = &c
	b.Left = &d
	c.Left = &e
	c.Right = &f
	PreOrder(&a)
	InOrder(&a)
}
