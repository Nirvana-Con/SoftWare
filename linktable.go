package main

import "fmt"

type LinkTableNode struct {
	Next *LinkTableNode
	Val  interface{}
}

type LinkTable struct {
	Head, Tail *LinkTableNode
	length     int
}

type list interface {
	AddHead(interface{}) bool
	AddTail(interface{}) bool
	Delete(interface{}) bool
	Len() int
	Search(interface{}) (int, bool)
	View() []interface{}
	GetHead() *LinkTableNode
	GetTail() *LinkTableNode
}

func NewLinkTable() list {
	return &LinkTable{nil, nil, 0}
}

func (l *LinkTable) Len() int {
	return l.length
}

func (l *LinkTable) GetHead() *LinkTableNode {
	return l.Head
}

func (l *LinkTable) GetTail() *LinkTableNode {
	return l.Tail
}

func (l *LinkTable) Delete(value interface{}) bool {
	if l.Head == nil {
		return false
	}
	if l.Head.Val == value {
		l.Head = l.Head.Next
		l.length = 0
		return true
	}
	p := l.Head
	q := p.Next
	for q != nil {
		if q.Val == value {
			if q == l.Tail {
				l.Tail = p
			}
			p.Next = q.Next
			l.length--
			return true
		}
		p = q
		q = q.Next
	}
	return false
}

//头插法
func (l *LinkTable) AddHead(value interface{}) bool {
	node := &LinkTableNode{nil, value}
	if l.Head == nil {
		l.Head = node
		l.Tail = node
	}
	node.Next = l.Head
	l.Head = node
	l.length++
	return true
}

//尾插法
func (l *LinkTable) AddTail(value interface{}) bool {
	node := &LinkTableNode{nil, value}
	if l.Head == nil {
		l.Head = node
		l.Tail = node
	}
	l.Tail.Next = node
	l.Tail = node
	l.length++
	return true
}

func (l *LinkTable) Search(value interface{}) (int, bool) {
	p := l.Head
	i := 0
	for p != nil {
		if p.Val == value {
			return i, true
		}
		p = p.Next
		i++
	}
	return -1, false
}

func (l *LinkTable) View() []interface{} {
	if l.Head == nil {
		return []interface{}{}
	}
	ans := make([]interface{}, 0)
	p := l.Head
	for p != nil {
		ans = append(ans, p.Val)
		p = p.Next
	}
	return ans
}

func main() {
	table := NewLinkTable()
	table.AddHead(1)
	table.AddHead(2)
	table.AddTail(3)
	fmt.Println(table.View())
	fmt.Println(table.Search(1))
	fmt.Println(table.Delete(3))
	fmt.Println(*table.GetTail())
	fmt.Println(table.View())
}
