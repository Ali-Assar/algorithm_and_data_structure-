package main

import "fmt"

type Node struct {
	value int
	left  *Node
	right *Node
}

type BinarySearchTree struct {
	root *Node
}

func NewNode(value int) *Node {
	return &Node{value: value}
}

func (bst *BinarySearchTree) Insert(value int) bool {
	newNode := NewNode(value)
	if bst.root == nil {
		bst.root = newNode
		return true
	} else {
		temp := bst.root
		for {
			if newNode.value == temp.value {
				return false
			}
			if newNode.value < temp.value {
				if temp.left == nil {
					temp.left = newNode
					return true
				}
				temp = temp.left
			} else {
				if temp.right == nil {
					temp.right = newNode
					return true
				}
				temp = temp.right
			}
		}
	}
}

func (bst *BinarySearchTree) Contains(value int) bool {
	temp := bst.root
	for temp != nil {
		if value < temp.value {
			temp = temp.left
		} else if value > temp.value {
			temp = temp.right
		} else {
			return true
		}
	}
	return false
}

// Example Usage
func main() {
	myTree := &BinarySearchTree{}
	myTree.Insert(2)
	myTree.Insert(1)
	myTree.Insert(3)
	myTree.Insert(10)
	myTree.Insert(33)
	myTree.Insert(45)
	myTree.Insert(12)
	myTree.Insert(98)

	fmt.Println(myTree.root.value)
	fmt.Println(myTree.root.left.value)
	fmt.Println(myTree.root.right.value)
	fmt.Println(myTree.Contains(12))
	fmt.Println(myTree.Contains(67))
}
