package main

import (
	"fmt"
	"mystack"
)

type Node struct {
	Data  interface{}
	Left  *Node
	Right *Node
}
type BiTree Node

func PreOrder(root *Node) {
	stack := mystack.NewStack()
	stack.Push(root)
	for !stack.Empty() {
		val, err := stack.Pop()
		fmt.Printf("%+v", val)
		fmt.Println("duandian", err)
		if err == nil {
			node, er := val.(Node)
			fmt.Printf("%+v", node)
			fmt.Println(er)
			fmt.Println("第一次：", node.Data)
			if node.Right != nil {
				stack.Push(node.Right)
			}
			if node.Left != nil {
				stack.Push(node.Left)
			}
		}
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
}
