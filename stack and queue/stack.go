package main

import "fmt"

type Node struct {
    value int
    next  *Node
}

type Stack struct {
    top    *Node
    height int
}

func NewStack(value int) *Stack {
    newNode := &Node{value: value}
    return &Stack{
        top:    newNode,
        height: 1,
    }
}

func (s *Stack) PrintStack() {
    temp := s.top
    for temp != nil {
        fmt.Println(temp.value)
        temp = temp.next
    }
}

func (s *Stack) Push(value int) {
    newNode := &Node{value: value}
    if s.height == 0 {
        s.top = newNode
    } else {
        newNode.next = s.top
        s.top = newNode
    }
    s.height++
}

func (s *Stack) Pop() *Node {
    if s.height == 0 {
        return nil
    }
    temp := s.top
    s.top = temp.next
    temp.next = nil
    s.height--
    return temp
}

func main() {
    myStack := NewStack(4)
    myStack.Push(6)
    myStack.Push(8)
    myStack.PrintStack()

    fmt.Println("Popped:", myStack.Pop().value)

    myStack.PrintStack()
}
