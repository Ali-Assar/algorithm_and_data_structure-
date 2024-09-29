package main

import (
	"fmt"
)

type Node struct {
	value int
	left  *Node
	right *Node
}

type BinarySearchTree struct {
	root *Node
}

func (bst *BinarySearchTree) Insert(value int) bool {
	newNode := &Node{value: value}
	if bst.root == nil {
		bst.root = newNode
		return true
	}
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

func (bst *BinarySearchTree) BFS() []int {
	results := []int{}
	if bst.root == nil {
		return results
	}
	queue := []*Node{bst.root}

	for len(queue) > 0 {
		currentNode := queue[0]
		queue = queue[1:]
		results = append(results, currentNode.value)
		if currentNode.left != nil {
			queue = append(queue, currentNode.left)
		}
		if currentNode.right != nil {
			queue = append(queue, currentNode.right)
		}
	}
	return results
}

func (bst *BinarySearchTree) DFSPreOrder() []int {
	var results []int
	var traverse func(node *Node)
	traverse = func(node *Node) {
		results = append(results, node.value)
		if node.left != nil {
			traverse(node.left)
		}
		if node.right != nil {
			traverse(node.right)
		}
	}
	traverse(bst.root)
	return results
}

func (bst *BinarySearchTree) DFSPostOrder() []int {
	var results []int
	var traverse func(node *Node)
	traverse = func(node *Node) {
		if node.left != nil {
			traverse(node.left)
		}
		if node.right != nil {
			traverse(node.right)
		}
		results = append(results, node.value)
	}
	traverse(bst.root)
	return results
}

func (bst *BinarySearchTree) DFSInOrder() []int {
	var results []int
	var traverse func(node *Node)
	traverse = func(node *Node) {
		if node.left != nil {
			traverse(node.left)
		}
		results = append(results, node.value)
		if node.right != nil {
			traverse(node.right)
		}
	}
	traverse(bst.root)
	return results
}

func main() {
	tree := BinarySearchTree{}
	values := []int{28, 45, 98, 11, 21, 86, 21}
	for _, value := range values {
		tree.Insert(value)
	}

	fmt.Println("BFS:", tree.BFS())
	fmt.Println("DFS Pre-order:", tree.DFSPreOrder())
	fmt.Println("DFS Post-order:", tree.DFSPostOrder())
	fmt.Println("DFS In-order:", tree.DFSInOrder())
}
