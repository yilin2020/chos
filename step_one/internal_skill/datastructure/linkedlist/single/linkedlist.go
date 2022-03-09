package single

import (
	"bytes"
	"fmt"
)

type Linkedlist struct {
	root *Node
	size int
}
type Node struct {
	data int
	next *Node
}

func New() *Linkedlist {
	return &Linkedlist{
		root: &Node{},
	}
}
func (l *Linkedlist) Size() int {
	return l.size
}
func (l *Linkedlist) Empty() bool {
	return l.size == 0
}

func (l *Linkedlist) ValueAt(index int) int {
	if index < 0 || index >= l.size {
		return -1
	}
	cur := l.root
	for i := 0; i < index; i, cur = i+1, cur.next {
		if cur == nil {
			return -1
		}
	}
	return cur.data
}

func (l *Linkedlist) PushFront(value int) {
	cur := l.root
	cur.next = &Node{next: cur.next, data: value}
	l.size++
}
func (l *Linkedlist) PushBack(value int) {
	cur := l.root
	for cur.next != nil {
		cur = cur.next
	}
	cur.next = &Node{data: value}
	l.size++
}

func (l *Linkedlist) PopFront() int {
	if l.size == 0 {
		return -1
	}
	cur := l.root
	deleteNode := cur.next
	cur.next = deleteNode.next
	l.size--
	deleteNode = nil
	return deleteNode.data
}
func (l *Linkedlist) PopBack() int {
	cur := l.root
	for cur.next != nil {
		cur = cur.next
	}
	perv := cur
	perv.next = nil
	cur = nil
	l.size--
	return perv.data
}
func (l *Linkedlist) String() string {
	cur := l.root.next

	var buf bytes.Buffer
	for cur != nil {
		buf.WriteString(fmt.Sprint(cur.data))
		cur = cur.next
	}
	return buf.String()
}
