package main

import (
	"errors"
	"fmt"
	"strconv"
)

func main() {
	l := NewList()

	for i := 0; i < 3; i++ {
		err := l.Insert(i, i)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
	l.PrintList()
	l.Reverse()
	fmt.Println(l.isNull())
	l.PrintList()
}

type Node struct {
	data interface{}
	next *Node
}

type List struct {
	head   *Node
	length int
}

type chainError struct {
	pos    int
	length int
}

func NewChainError(pos, len int) *chainError {
	return &chainError{pos, len}
}

func (e *chainError) Error() string {
	return "Error: 您希望在第" + strconv.Itoa(e.pos+1) + "个元素处插入节点, 但是链表长度仅为" + strconv.Itoa(e.length) + "!\n"
}

type Noder interface {
	Insert(i int, v interface{})
	Delete(i int)
	Getlength() int
	Search(v interface{}) []int
	isNull() bool
}

type ExtNoder interface {
	Noder
	Head() Noder
	AllSeek()
}

func NewNode(v interface{}) *Node {
	return &Node{data: v}
}

func NewList() *List {
	return &List{NewNode(nil), 0}
}

func (l *List) GetLength() int {
	pre := l.head
	for count := 0; ; {
		if pre.next == nil {
			return count
		}
		count++
		pre = pre.next
	}
}

func (l *List) Insert(i int, v interface{}) error { // i从0开始计算
	n := NewNode(v)
	pre := l.head
	if l.GetLength() < i {
		return NewChainError(i, l.GetLength())
	}
	for positon := 0; positon <= i; positon++ {
		if positon == i {
			n.next = pre.next
			pre.next = n
			l.length++
		}
		pre = pre.next
	}
	return nil
}

func (l *List) Delete(i int) error {
	pre := l.head
	if l.GetLength() < i+1 {
		return errors.New("您要删除的元素超出链表长度范围!\n")
	}
	for pos := 0; pos <= i; pos++ {
		if pos == i {
			pre.next = pre.next.next
			l.length--
		}
		pre = pre.next
	}
	return nil
}

func (l *List) Search(v interface{}) int {
	pre := l.head
	for pos := 0; pos <= l.GetLength()-1; pos++ {
		if pre.next.data == v {
			return pos
		}
		pre = pre.next
	}
	return -1
}

func (l *List) isNull() bool {
	pre := l.head
	if pre.next == nil {
		return true
	}
	return false
}

func (l *List) PrintList() {
	pre := l.head
	for pos := 0; pos <= l.GetLength()-1; pos++ {
		fmt.Printf("%v\n", pre.next.data)
		pre = pre.next
	}
}

func (l *List) Reverse() {
	var pre *Node = l.head.next
	var cur *Node = pre.next
	var next *Node = cur.next
	for {
		fmt.Println(pre)
		if cur == nil {
			l.head.next.next = nil
			l.head.next = pre
			break
		}
		cur.next = pre
		pre = cur
		cur = next
		if cur != nil {
			next = next.next
		}
	}
}
