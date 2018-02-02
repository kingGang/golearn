package main

import (
	"fmt"
	"strconv"
)

type Node struct {
	data  string
	pNext *Node
}

type LinkList *Node

func InitList() (linkList *Node) {
	linkList = &Node{}
	return linkList
}
func CreateListHead(l *Node, n int) int {
	for i := 0; i < n; i++ {
		p := Node{}
		p.data = strconv.Itoa(i)
		p.pNext = l.pNext
		l.pNext = &p
	}
	return 1
}
func CreateListTail(l *Node, n int) int {
	temp := l
	for i := 0; i < n; i++ {
		p := Node{}
		p.data = strconv.Itoa(i)
		fmt.Printf("%+v", p)
		fmt.Println(i)
		p.pNext = nil
		temp.pNext = &p
		temp = &p
		fmt.Printf("%+v", temp)
	}
	return 1
}
func GetLength(l *Node) int {
	var length int = 0
	p := l.pNext
	for p.pNext != nil {
		length++
		fmt.Println("p.data=", p.data)
		p = p.pNext
	}
	return length
}

func main() {
	linkList := InitList()
	CreateListHead(linkList, 1)
	fmt.Printf("linkList=%+v,长度=%d", linkList, GetLength(linkList))
}
