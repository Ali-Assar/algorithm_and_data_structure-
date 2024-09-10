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

func (bst *BinarySearchTree) rInsert(value int) {
	bst.root = bst.rInsertRecursive(bst.root, value)
}

func (bst *BinarySearchTree) rInsertRecursive(currentNode *Node, value int) *Node {
	if currentNode == nil {
		return &Node{value: value}
	}
	if value < currentNode.value {
		currentNode.left = bst.rInsertRecursive(currentNode.left, value)
	} else if value > currentNode.value {
		currentNode.right = bst.rInsertRecursive(currentNode.right, value)
	}
	return currentNode
}

func (bst *BinarySearchTree) rContains(value int) bool {
	return bst.rContainsRecursive(bst.root, value)
}

func (bst *BinarySearchTree) rContainsRecursive(currentNode *Node, value int) bool {
	if currentNode == nil {
		return false
	}
	if value == currentNode.value {
		return true
	}
	if value < currentNode.value {
		return bst.rContainsRecursive(currentNode.left, value)
	}
	return bst.rContainsRecursive(currentNode.right, value)
}

func (bst *BinarySearchTree) deleteNode(value int) {
	bst.root = bst.deleteNodeRecursive(bst.root, value)
}

func (bst *BinarySearchTree) deleteNodeRecursive(currentNode *Node, value int) *Node {
	if currentNode == nil {
		return nil
	}
	if value < currentNode.value {
		currentNode.left = bst.deleteNodeRecursive(currentNode.left, value)
	} else if value > currentNode.value {
		currentNode.right = bst.deleteNodeRecursive(currentNode.right, value)
	} else {
		// Case 1: No children
		if currentNode.left == nil && currentNode.right == nil {
			return nil
		}
		// Case 2: One child
		if currentNode.left == nil {
			return currentNode.right
		}
		if currentNode.right == nil {
			return currentNode.left
		}
		// Case 3: Two children
		subTreeMin := bst.minValue(currentNode.right)
		currentNode.value = subTreeMin
		currentNode.right = bst.deleteNodeRecursive(currentNode.right, subTreeMin)
	}
	return currentNode
}

func (bst *BinarySearchTree) minValue(node *Node) int {
	currentNode := node
	for currentNode.left != nil {
		currentNode = currentNode.left
	}
	return currentNode.value
}

func main() {
	bst := BinarySearchTree{}
	bst.rInsert(2)
	bst.rInsert(1)
	bst.rInsert(3)

	fmt.Println("Root:", bst.root.value)
	fmt.Println("Root -> left:", bst.root.left.value)
	fmt.Println("Root -> right:", bst.root.right.value)

	fmt.Println("BST contains 3?", bst.rContains(3))
	fmt.Println("BST contains 5?", bst.rContains(5))

	bst.deleteNode(2)
	fmt.Println("Root after deletion:", bst.root.value)
	fmt.Println("Root -> left:", bst.root.left.value)
	fmt.Println("Root -> right:", bst.root.right)
}
