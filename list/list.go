package list

import "fmt"

type Node struct {
	value int
	next  *Node
}

type List struct {
	head  *Node
	count int
}

func NewNode(v int) *Node {
	return &Node{value: v}
}

func NewList(values ...int) *List {
	l := List{}
	if len(values) == 0 {
		return &l
	}

	for i := len(values) - 1; i >= 0; i-- {
		l.AddHead(values[i])
	}
	return &l
}

func (l *List) Size() int {
	return l.count
}

func (l *List) IsEmpty() bool {
	return l.count == 0
}

func (l *List) AddHead(v int) {
	newNode := NewNode(v)
	if l.IsEmpty() {
		l.head = newNode
		l.count++
		return
	}

	newNode.next = l.head
	l.head = newNode
	l.count++
}

func (l *List) AddTail(v int) {
	newNode := NewNode(v)
	if l.IsEmpty() {
		l.head = newNode
		l.count++
		return
	}

	if l.head == nil {
		l.head.next = newNode
		l.count++
		return
	}

	n := l.head
	for {
		if n.next == nil {
			n.next = newNode
			l.count++
			return
		}
		n = n.next
	}
}

func (l *List) Print() {
	curr := l.head
	for curr != nil {
		fmt.Printf("%d ", curr.value)
		curr = curr.next
	}
}
