package main

import "fmt"

type Node struct {
	value int
	next  *Node
}

type LinkedList struct {
	head   *Node
	tail   *Node
	length int
}

func NewLinkedList(value int) *LinkedList {
	newNode := &Node{value: value}
	return &LinkedList{
		head:   newNode,
		tail:   newNode,
		length: 1,
	}
}

func (l *LinkedList) PrintList() {
	temp := l.head
	for temp != nil {
		fmt.Println(temp.value)
		temp = temp.next
	}
}

func (l *LinkedList) Append(value int) bool {
	newNode := &Node{value: value}
	if l.head == nil {
		l.head = newNode
		l.tail = newNode
	} else {
		l.tail.next = newNode
		l.tail = newNode
	}
	l.length++
	return true
}

func (l *LinkedList) Pop() *Node {
	if l.length == 0 {
		return nil
	}
	temp := l.head
	var pre *Node
	for temp.next != nil {
		pre = temp
		temp = temp.next
	}
	l.tail = pre
	if l.tail != nil {
		l.tail.next = nil
	}
	l.length--
	if l.length == 0 {
		l.head = nil
		l.tail = nil
	}
	return temp
}

func (l *LinkedList) Prepend(value int) bool {
	newNode := &Node{value: value}
	if l.length == 0 {
		l.head = newNode
		l.tail = newNode
	} else {
		newNode.next = l.head
		l.head = newNode
	}
	l.length++
	return true
}

func (l *LinkedList) PopFirst() *Node {
	if l.length == 0 {
		return nil
	}
	temp := l.head
	l.head = l.head.next
	temp.next = nil
	l.length--
	if l.length == 0 {
		l.tail = nil
	}
	return temp
}

func (l *LinkedList) Get(index int) *Node {
	if index < 0 || index >= l.length {
		return nil
	}
	temp := l.head
	for i := 0; i < index; i++ {
		temp = temp.next
	}
	return temp
}

func (l *LinkedList) SetValue(index int, value int) bool {
	temp := l.Get(index)
	if temp != nil {
		temp.value = value
		return true
	}
	return false
}

func (l *LinkedList) Insert(index int, value int) bool {
	if index < 0 || index > l.length {
		return false
	}
	if index == 0 {
		return l.Prepend(value)
	}
	if index == l.length {
		return l.Append(value)
	}
	newNode := &Node{value: value}
	prev := l.Get(index - 1)
	newNode.next = prev.next
	prev.next = newNode
	l.length++
	return true
}

func (l *LinkedList) Remove(index int) *Node {
	if index < 0 || index >= l.length {
		return nil
	}
	if index == 0 {
		return l.PopFirst()
	}
	if index == l.length-1 {
		return l.Pop()
	}
	prev := l.Get(index - 1)
	temp := prev.next
	prev.next = temp.next
	temp.next = nil
	l.length--
	return temp
}

func (l *LinkedList) Reverse() {
	temp := l.head
	l.head = l.tail
	l.tail = temp
	var after *Node
	var before *Node
	for i := 0; i < l.length; i++ {
		after = temp.next
		temp.next = before
		before = temp
		temp = after
	}
}

func main() {
	ll := NewLinkedList(10)
	ll.Append(20)
	ll.Append(30)
	ll.PrintList()

	ll.Reverse()
	ll.PrintList()
}
