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

func InitList() (l LinkList) {
	l = &Node{}
	l.pNext = nil
	return l
}

func CreateListHead(l LinkList, n int) int {
	for i := 0; i < n; i++ {
		p := Node{}
		p.data = strconv.Itoa(i)
		p.pNext = l.pNext
		l.pNext = &p
	}
	return 1
}
func CreateListTail(l LinkList, n int) int {
	temp := l
	for i := 0; i < n; i++ {
		p := Node{}
		p.data = strconv.Itoa(i)
		p.pNext = nil
		temp.pNext = &p
		temp = &p
	}
	return 1
}
func GetLength(l LinkList) int {
	var length int = 0
	p := l.pNext
	for p != nil {
		length++
		p = p.pNext
	}
	return length
}
func PrintLinkList(l LinkList) {
	p := l.pNext
	var i int = 0
	fmt.Println("-----------打印链表-----------")
	if p == nil {
		fmt.Println("链表为空")
	}
	for p != nil {
		i++
		fmt.Printf("第%d个节点的数据为%+v\n", i, p)
		p = p.pNext
	}
}
func GetElem(l LinkList, i int) (n *Node, err int) {
	p := l.pNext
	if p == nil {
		fmt.Println("链表为空")
		err = 0
		return
	}
	if i < 0 {
		fmt.Printf("你所查询的节点%d应该大于0，请重新输入", i)
		err = 0
		return
	}
	j := 1
	for p != nil && j < i {
		j++
		p = p.pNext
	}
	if p == nil {
		fmt.Println("你所输入的节点超出长度")
		err = 0
		return
	}
	n = p
	err = 1
	return

}
func InsertList(l LinkList, i int, data string) {
	p := l.pNext
	if p == nil {
		fmt.Println("链表为空，默认插入第一个节点")
		insNode := Node{}
		insNode.data = data
		insNode.pNext = nil
		l.pNext = &insNode
	}
	j := 1
	for p != nil && j < i {
		j++
		p = p.pNext
	}
	if p == nil {
		fmt.Println("长度超出，请重新操作")
		return
	}
	insNode := Node{}
	insNode.data = data
	insNode.pNext = p.pNext
	p.pNext = &insNode
}

func main() {
	linkList := InitList()
	CreateListTail(linkList, 10)
	//PrintLinkList(linkList);
	//fmt.Println(GetElem(linkList,9))
	InsertList(linkList, 0, "你好")
	PrintLinkList(linkList)
	//fmt.Printf("\nlinkList=%+v,长度=%d\r\n", linkList, GetLength(linkList))
}
