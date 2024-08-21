package main

import "fmt"

type Node struct {
    Value int
    Next  *Node
    Prev  *Node
}

type DoublyLinkedList struct {
    Head   *Node
    Tail   *Node
    Length int
}

func NewDoublyLinkedList(value int) *DoublyLinkedList {
    newNode := &Node{Value: value}
    return &DoublyLinkedList{
        Head:   newNode,
        Tail:   newNode,
        Length: 1,
    }
}

func (dll *DoublyLinkedList) PrintList() {
    temp := dll.Head
    for temp != nil {
        fmt.Println(temp.Value)
        temp = temp.Next
    }
}

func (dll *DoublyLinkedList) Append(value int) bool {
    newNode := &Node{Value: value}
    if dll.Length == 0 {
        dll.Head = newNode
        dll.Tail = newNode
    } else {
        dll.Tail.Next = newNode
        newNode.Prev = dll.Tail
        dll.Tail = newNode
    }
    dll.Length++
    return true
}

func (dll *DoublyLinkedList) Pop() *Node {
    if dll.Length == 0 {
        return nil
    }
    temp := dll.Tail
    if dll.Length == 1 {
        dll.Head = nil
        dll.Tail = nil
    } else {
        dll.Tail = dll.Tail.Prev
        dll.Tail.Next = nil
        temp.Prev = nil
    }
    dll.Length--
    return temp
}

func (dll *DoublyLinkedList) Prepend(value int) {
    newNode := &Node{Value: value}
    if dll.Length == 0 {
        dll.Head = newNode
        dll.Tail = newNode
    } else {
        newNode.Next = dll.Head
        dll.Head.Prev = newNode
        dll.Head = newNode
    }
    dll.Length++
}

func (dll *DoublyLinkedList) PopFirst() *Node {
    if dll.Length == 0 {
        return nil
    }
    temp := dll.Head
    if dll.Length == 1 {
        dll.Head = nil
        dll.Tail = nil
    } else {
        dll.Head = dll.Head.Next
        dll.Head.Prev = nil
        temp.Next = nil
    }
    dll.Length--
    return temp
}

func (dll *DoublyLinkedList) Get(index int) *Node {
    if index < 0 || index >= dll.Length {
        return nil
    }
    var temp *Node
    if index < dll.Length/2 {
        temp = dll.Head
        for i := 0; i < index; i++ {
            temp = temp.Next
        }
    } else {
        temp = dll.Tail
        for i := dll.Length - 1; i > index; i-- {
            temp = temp.Prev
        }
    }
    return temp
}

func (dll *DoublyLinkedList) SetValue(index int, value int) bool {
    temp := dll.Get(index)
    if temp != nil {
        temp.Value = value
        return true
    }
    return false
}

func (dll *DoublyLinkedList) Insert(index int, value int) bool {
    if index < 0 || index > dll.Length {
        return false
    }
    if index == 0 {
        dll.Prepend(value)
        return true
    }
    if index == dll.Length {
        dll.Append(value)
        return true
    }
    newNode := &Node{Value: value}
    before := dll.Get(index - 1)
    after := before.Next
    newNode.Prev = before
    newNode.Next = after
    before.Next = newNode
    if after != nil {
        after.Prev = newNode
    }
    dll.Length++
    return true
}

func (dll *DoublyLinkedList) Remove(index int) *Node {
    if index < 0 || index >= dll.Length {
        return nil
    }
    if index == 0 {
        return dll.PopFirst()
    }
    if index == dll.Length-1 {
        return dll.Pop()
    }
    temp := dll.Get(index)
    temp.Prev.Next = temp.Next
    if temp.Next != nil {
        temp.Next.Prev = temp.Prev
    }
    temp.Next = nil
    temp.Prev = nil
    dll.Length--
    return temp
}

func main() {
    doubleLinkList := NewDoublyLinkedList(7)
    doubleLinkList.Append(5)
    doubleLinkList.Append(6)
    doubleLinkList.Append(4)
    doubleLinkList.Append(3)
    doubleLinkList.Pop()
    doubleLinkList.Prepend(1)
    doubleLinkList.PopFirst()
    doubleLinkList.PrintList()
    fmt.Println(doubleLinkList.Get(1))
    fmt.Println(doubleLinkList.SetValue(1, 22))
    fmt.Println(doubleLinkList.Insert(4, 33))
    fmt.Println(doubleLinkList.Remove(2))
}
