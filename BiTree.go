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

//中序遍历
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
			interVal, _ := interVal.(*Node)
			fmt.Println(interVal.Data)
			root = interVal.Right
			// fmt.Printf("print node :%+v\n", root)
		}
	}
}

//后序遍历
func TailOrder(root *Node) {
	tempStack := mystack.NewStack()
	outStack := mystack.NewStack()
	tempStack.Push(root)
	for !tempStack.Empty() {
		temp, err := tempStack.Pop()
		if err == nil {
			outStack.Push(temp)
			val, ok := temp.(interface{})
			if ok {
				val, ok := val.(*Node)
				if ok {
					if val.Left != nil {
						tempStack.Push(val.Left)
					}
					if val.Right != nil {
						tempStack.Push(val.Right)
					}
				}

			}
		}
	}
	for !outStack.Empty() {
		value, _ := outStack.Pop()
		fmt.Printf("%+v\n", value)
	}
}

//递归求二叉树叶子节点数
func BtLeaf(root *Node) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 1
	}
	return BtLeaf(root.Left) + BtLeaf(root.Right)
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
	g := Node{}
	g.Data = "g"
	a.Left = &b
	a.Right = &c
	b.Left = &d
	b.Right = &g
	c.Left = &e
	c.Right = &f
	PreOrder(&a)
	fmt.Println("---------------------\n")
	InOrder(&a)
	fmt.Println("---------------------\n")
	TailOrder(&a)

	fmt.Println("叶子节点数为：", BtLeaf(&a))
}
