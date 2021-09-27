/*

 */
package main

import (
	"fmt"
)

type Stack struct {
	size int64
	top  int64
	data []interface{}
}

func (s *Stack) IsFull() bool {
	return s.top == s.size
}
func (s *Stack) IsEmpty() bool {
	return s.top == 0
}
func (s *Stack) Push(e interface{}) bool {
	if s.IsFull() {
		fmt.Printf("栈满，无法入栈")
		return false
	}
	s.data[s.top] = e
	s.top++
	return true
}
func (s *Stack) Pop() (r interface{}) {
	if s.IsEmpty() {
		fmt.Printf("栈空，无法出栈")
		return
	}
	s.top--
	r = s.data[s.top]
	return
}
func (s *Stack) Mate(str string) {
	for _, v := range str {
		s.Push(string(v))
		if string(v) == "{" || string(v) == "}" {
			fmt.Println(s.Pop())
		}
	}

}

func MakeStack(size int64) Stack {
	s := Stack{}
	s.size = size
	s.data = make([]interface{}, size)
	return s
}

func main() {
	str := "fdasfsad{fdsaf}grew{fdsfd}{gfsgs}"
	s := MakeStack(50)
	s.Mate(str)

}
