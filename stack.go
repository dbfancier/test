package main

import (
	"errors"
	"fmt"
)

func main() {
	s := NewStack()
	s.Push("Carrot")
	s.Push(100)
	s.Push(3.141592653)

	if data, err := s.Top(); err == nil {
		fmt.Println(data)
	} else {
		fmt.Println(err)
	}

	if data, err := s.Pop(); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(data)
	}

	fmt.Println(s)

}

type Stack struct {
	arr       []interface{}
	stackSize int
}

func NewStack() *Stack {
	var arr = make([]interface{}, 0)
	return &Stack{
		arr,
		0,
	}
}

func (s *Stack) Size() int {
	return s.stackSize
}

func (s *Stack) isEmpty() bool {
	if s.Size() == 0 {
		return true
	}
	return false
}

func (s *Stack) Push(data interface{}) {
	s.arr = append(s.arr, data)
	s.stackSize++
}

func (s *Stack) Pop() (interface{}, error) {
	if s.isEmpty() {
		return nil, errors.New("Stack size is 0, No element to pop.")
	}
	ele := s.arr[s.Size()-1]
	s.stackSize--
	s.arr = s.arr[0:s.Size()]
	return ele, nil
}

func (s *Stack) Top() (interface{}, error) { // 取栈顶元素，但不弹出
	if s.isEmpty() {
		return nil, errors.New("Stack size is 0, No element to fetch.")
	}
	return s.arr[s.Size()-1], nil
}
