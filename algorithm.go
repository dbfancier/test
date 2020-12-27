package main

import (
	"errors"
	"fmt"
	"strconv"
)

func main() {
	l := NewList()

	/*for i := 0; i < 20; i++ {
		err := l.Insert(i, i)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
	l.PrintList()
	l.InsertReverse()
	fmt.Println(l.isNull())
	l.PrintList() */
	/*for i := 0; i <= 20; i++ {
		for j := 0; j < i; j++ {
			l.Add(i)
		}
	}*/
	l.Add(8)
	l.Add(0)
	l.Add(0)
	l.Add(1)
	l.Add(3)
	l.Add(8)
	l.Add(1)
	l.Add(0)
	l.Add(1)
	l.Add(8)
	l.PrintList()
	l.RemoveDup()

	fmt.Println("-------")
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

func (l *List) Add(v interface{}) {
	n := NewNode(v)
	cur := l.head
	for pos := 0; ; pos++ {
		if cur.next == nil {
			cur.next = n
			break
		}
		cur = cur.next
	}
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

func (l *List) Search(v interface{}) []int {
	cur := l.head.next
	var result = []int{}
	for pos := 0; pos <= l.GetLength()-1; pos++ {
		if cur.data == v {
			result = append(result, pos)
		} else if cur == nil {
			return result
		}
		cur = cur.next
	}
	if len(result) == 0 {
		result = append(result, -1)
	}
	return result
}

func (l *List) SearchN(n int) *Node {
	cur := l.head.next
	for pos := 0; ; pos++ {
		if pos == n {
			return cur
		}
		cur = cur.next
	}
	return nil
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

func (l *List) Reverse() { // 就地翻转
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

func (l *List) InsertReverse() { // 从第二个元素开始，插到[0]位置，实现翻转
	cur := l.head.next.next
	for pos := 1; pos <= l.GetLength()-1; pos++ {
		l.Delete(pos)
		l.Insert(0, cur.data)
		// if cur == nil { break }
		cur = cur.next
	}
}

func (l *List) RemoveDupOrder() { // 用两个变量控制外循环和内循环中的Node，通过NestLoop的方式进行对比删除(顺序删除)
	cur := l.head.next
	var tmp *Node
	for pos := 0; pos <= l.GetLength()-1; pos++ {
		tmp = cur
		for cursor := pos + 1; cursor <= l.GetLength()-1; cursor++ {
			if tmp.data == tmp.next.data {
				l.Delete(cursor)
			}
			tmp = tmp.next
		}
		cur = cur.next
	}
}

func (l *List) RemoveDup() { // 用两个变量控制外循环和内循环中的Node，通过NestLoop的方式进行对比删除
	cur := l.head.next
	delCount := 0
	for cur != nil {
		cursors := l.Search(cur.data)
		cursors = cursors[1:]
		for _, v := range cursors {
			l.Delete(v - delCount)
			delCount++
		}
		delCount = 0
		cur = cur.next
	}
}
