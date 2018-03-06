package mystack

import (
	"errors"
	"fmt"
)

type Stack struct {
	Element []interface{}
}

func NewStack() *Stack {
	return &Stack{}
}
func (stack *Stack) Push(value ...interface{}) {
	stack.Element = append(stack.Element, value...)
}
func (stack *Stack) Top() (value interface{}) {
	if stack.Size() > 0 {
		return stack.Element[stack.Size()-1]
	}
	return nil //read empty stack
}
func (stack *Stack) Pop() (value interface{}, err error) {
	if stack.Size() > 0 {
		value := stack.Element[stack.Size()-1]
		stack.Element = stack.Element[:stack.Size()-1]

		return value, nil
	}
	err = errors.New("Stack为空.")
	return nil, err

}
func (stack *Stack) Swap(other *Stack) {
	switch {
	case stack.Size() == 0 && other.Size() == 0:
		return
	case other.Size() == 0:
		other.Element = stack.Element[:stack.Size()]
		stack.Element = nil
	case stack.Size() == 0:
		stack.Element = other.Element
		other.Element = nil
	default:
		stack.Element, other.Element = other.Element, stack.Element
	}
	return
}
func (stack *Stack) Set(idx int, value interface{}) (err error) {
	if idx >= 0 && stack.Size() > 0 && stack.Size() > idx {
		stack.Element[idx] = value
		return nil
	}
	return errors.New("Set失败")
}
func (stack *Stack) Get(idx int) (value interface{}) {
	if idx > 0 && stack.Size() > 0 && stack.Size() > idx {
		return stack.Element[idx]
	}
	return nil
}
func (s *Stack) Size() int {
	return len(s.Element)
}
func (s *Stack) Empty() bool {
	if s.Element == nil || s.Size() == 0 {
		return true
	}
	return false
}
func (s *Stack) Print() {
	for i := 0; i <= len(s.Element)-1; i++ {
		fmt.Println(i, "=", s.Element[i])
	}
}

func main() {

	stack := NewStack()
	if stack.Empty() {
		fmt.Println("Stack为空! ")
	} else {
		fmt.Println("Stack不为空! ", stack.Size())
	}

	stack.Push(10)
	stack.Push(20)
	stack.Push(30)
	stack.Push(40)
	fmt.Println("当前Size() = ", stack.Size())
	stack.Print()

	fmt.Println("当前Top() = ", stack.Top())

	stack.Pop()
	fmt.Println("执行完Pop()后的Top() = ", stack.Top())
	stack.Print()

	stack.Set(2, 900)
	fmt.Println("\n执行完Set(2,900)后的Stack")
	stack.Print()

	fmt.Println("\nGet()查看指定的元素: ")
	fmt.Println("当前idx为1的元素 = ", stack.Get(1))
	fmt.Println("当前idx为2的元素 = ", stack.Get(2))

	stack2 := NewStack()
	stack2.Push("111")
	stack2.Push("222")
	fmt.Println("\nstack2的初始内容:")
	stack2.Print()

	stack.Swap(stack2)
	fmt.Println("Swap()后stack的内容:")
	stack.Print()
	fmt.Println("Swap()后stack2的内容:")
	stack2.Print()

	fmt.Println("\nstack增加字符串元素: ")
	stack.Push("中文元素")
	stack.Push("elem1")
	stack.Print()

}
