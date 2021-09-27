package main

import (
	"fmt"
)

type SeqList struct {
	data    []interface{}
	Length  int
	MaxSize int
}

func (s *SeqList) Locate(e interface{}) int {
	for i := 0; i < len(s.data); i++ {
		if s.data[i] == e {
			return i
		}
	}
	return -1
}

//插入
func (s *SeqList) InsertList(index int, e interface{}) bool {
	if index < 0 || index > s.MaxSize {
		fmt.Println("位置非法")
		return false
	}
	if s.Length >= s.MaxSize {
		fmt.Println("Overflow")
		return false
	}
	for i := s.Length; i >= index; i-- {
		s.data[i] = s.data[i-1]
	}
	s.data[index] = e
	s.Length++
	return true
}
func (s *SeqList) DelList(index int) bool {
	if index < 0 || index >= s.MaxSize {
		fmt.Println("删除位置不合法")
		return false
	}
	for i := index + 1; i <= len(s.data)-1; i++ {
		s.data[i-1] = s.data[i]
	}
	s.Length--
	return true
}

//noinspection GoBinaryAndUnaryExpressionTypesCompatibility
func main() {
	s := SeqList{}
	s.MaxSize = 10
	data := []int{1, 2, 3, 4, 5, 6, 7, 8}
	s.Length = len(data)
	s.data = make([]interface{}, s.MaxSize, s.MaxSize)
	//s.data=append(s.data,data[0],data[1:])
	//fmt.Printf("长度%d--%v",len(s.data),s.data)
	//s.data=data
	for i, d := range data {
		s.data[i] = d
	}
	fmt.Printf("%v", s)

}
