package main

import "fmt"

type Node struct {
	value int
	next  *Node
}

type Queue struct {
	first  *Node
	last   *Node
	length int
}

func NewQueue(value int) *Queue {
	newNode := &Node{value: value}
	return &Queue{
		first:  newNode,
		last:   newNode,
		length: 1,
	}
}

func (q *Queue) PrintQueue() {
	temp := q.first
	for temp != nil {
		fmt.Println(temp.value)
		temp = temp.next
	}
}

func (q *Queue) Enqueue(value int) {
	newNode := &Node{value: value}
	if q.length == 0 {
		q.first = newNode
		q.last = newNode
	} else {
		q.last.next = newNode
		q.last = newNode
	}
	q.length++
}

func (q *Queue) Dequeue() *Node {
	if q.length == 0 {
		return nil
	}
	temp := q.first
	if q.length == 1 {
		q.first = nil
		q.last = nil
	} else {
		q.first = q.first.next
		temp.next = nil
	}
	q.length--
	return temp
}

func main() {
	myQueue := NewQueue(4)
	myQueue.Enqueue(5)
	myQueue.Enqueue(8)
	myQueue.PrintQueue()

	fmt.Println("Dequeued:", myQueue.Dequeue().value)

	myQueue.PrintQueue()
}
