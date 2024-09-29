# Binary Search Tree (BST) Implementation in Go and Python

This Folder of repository contains implementations of a Binary Search Tree (BST) in both Go and Python. A BST is a widely used data structure that allows for efficient search, insertion, and deletion operations.

## What is a Binary Search Tree (BST)?

A Binary Search Tree is a binary tree in which each node has at most two children, referred to as the left child and the right child. The key property of a BST is that for any given node:

- All values in the left subtree are less than the node's value.
- All values in the right subtree are greater than the node's value.

This property allows for efficient searching, as you can eliminate half of the tree at each step.

### Why Use a Binary Search Tree?

- **Efficient Search**: The BST allows for average-case time complexity of O(log n) for search operations. This efficiency is due to the binary property, which halves the search space with each comparison.
- **Order**: BSTs maintain elements in a sorted order, which makes them useful for dynamic sets of ordered data.

## Key Operations and Their Complexity

### 1. Insert

**Operation**: Add a new element to the BST.

**Explanation**: Inserting an element into a BST involves finding the correct position based on the BST property. Starting from the root, if the new value is less than the current node's value, move to the left child; if it's greater, move to the right child. Repeat this process until an appropriate empty spot is found.

**Algorithm**:
- Start at the root.
- Compare the new value with the current node's value.
- If the new value is less, go to the left child; if greater, go to the right child.
- Insert the new value when an empty spot is found.

**Complexity**: O(log n) on average. However, in the worst case (a completely unbalanced tree), the complexity can degrade to O(n).

### 2. Contains

**Operation**: Check if a specific value exists in the BST.

**Explanation**: The contains operation leverages the BST's ordered structure to efficiently find a value. Starting from the root, the algorithm compares the target value with the current node's value, moving left or right depending on whether the target is smaller or larger.

**Algorithm**:
- Start at the root.
- Compare the target value with the current node's value.
- If the target value is equal to the current node's value, return `True`.
- If the target value is less, go to the left child; if greater, go to the right child.
- If the value is found, return `True`; otherwise, return `False`.

**Complexity**: O(log n) on average, but O(n) in the worst case.

## Conclusion

Binary Search Trees are a fundamental data structure in computer science, enabling efficient search, insertion, and deletion operations. By maintaining a sorted structure, BSTs allow you to quickly locate elements and maintain ordered data dynamically. Understanding how BSTs work and their associated complexities is essential for optimizing search-related tasks in various applications.
