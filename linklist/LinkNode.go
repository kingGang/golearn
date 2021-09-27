package main

import (
	"fmt"
	"strconv"
)

//节点
type Node struct {
	data  string
	pNext *Node
}

//指向节点的指针
type LinkList *Node

func InitList() (l LinkList) {
	l = &Node{}
	l.pNext = nil
	return l
}

//创建一个指向
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
	p := l
	if p == nil {
		fmt.Println("链表为空，默认插入第一个节点")
		insNode := Node{}
		insNode.data = data
		insNode.pNext = nil
		l.pNext = &insNode
	}
	j := 0
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
func InsertListTail(l LinkList, data string) {
	p := l
	for p.pNext != nil {
		p = p.pNext
	}
	insNode := Node{}
	insNode.data = data
	insNode.pNext = nil
	p.pNext = &insNode

}
func DelList(l LinkList, i int) int {
	p := l
	if p.pNext == nil {
		fmt.Println("链表为空，无法删除")
		return 0
	}
	var j int = 0
	for p.pNext != nil && j < i-1 {
		j++
		p = p.pNext
	}
	if p.pNext == nil {
		fmt.Println("你要删除的节点，超过链表长度\n")
	}
	p.pNext = p.pNext.pNext
	return 1
}
func ClearList(linklist LinkList) int {
	linklist.pNext = nil
	// pre := linkli
	// for pre != nil {
	// 	temp = pre
	// 	pre = pre.pNext
	// 	// free(temp)
	// }
	// linklist.pNext = nil
	return 1
}

func main() {
	linkList := InitList()
	fmt.Printf("%+v\n", linkList)
	CreateListTail(linkList, 10)
	fmt.Printf("插入后：%+v\n", linkList)
	//PrintLinkList(linkList);
	//fmt.Println(GetElem(linkList,9))
	InsertListTail(linkList, "你好")
	PrintLinkList(linkList)
	DelList(linkList, 10)
	PrintLinkList(linkList)
	ClearList(linkList)
	PrintLinkList(linkList)
	fmt.Printf("\nlinkList=%+v,长度=%d\r\n", linkList, GetLength(linkList))
}
