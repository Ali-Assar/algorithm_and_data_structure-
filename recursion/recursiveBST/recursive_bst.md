
# Recursive Binary Search Tree (BST)

## Introduction

A **Binary Search Tree (BST)** is a binary tree where each node contains a value, and it follows these rules:
- The left subtree of a node contains only values less than the node's value.
- The right subtree of a node contains only values greater than the node's value.
- Both subtrees are also binary search trees.

In this explanation, we will cover the recursive approach for:
1. Inserting values into a BST.
2. Searching for values in a BST.
3. Deleting a node from a BST.

## Recursive Algorithm

### Recursive Insert

The recursive insert function follows a divide-and-conquer approach. We compare the value to be inserted with the current node’s value and recurse accordingly:

- If the current node is `None` (i.e., an empty spot in the tree), we create a new node with the value.
- If the value is smaller than the current node's value, we move to the left subtree and try to insert the value recursively.
- If the value is larger than the current node's value, we move to the right subtree and insert the value recursively.

This process continues until the appropriate spot for the new node is found.

#### Time Complexity

- **Best and Average Case**: O(log n) – If the tree is balanced, each recursive call moves us closer to the leaf node, reducing the problem size by half.
- **Worst Case**: O(n) – If the tree is unbalanced and resembles a linked list, insertion could take linear time.

### Recursive Search (Contains)

The recursive search function follows a similar strategy to the insert function:

- If the current node is `None`, the value is not in the tree.
- If the current node's value matches the value we're searching for, we return `True`.
- If the value is smaller than the current node’s value, we recurse to the left subtree.
- If the value is larger than the current node’s value, we recurse to the right subtree.

#### Time Complexity

- **Best and Average Case**: O(log n) – Searching the tree cuts the problem size by half with each recursive step.
- **Worst Case**: O(n) – If the tree is unbalanced, the search could take linear time.

### Recursive Delete

Deletion in a BST is more complex because we need to consider three cases:
1. **Node to be deleted has no children**: Simply remove the node.
2. **Node to be deleted has one child**: Replace the node with its child.
3. **Node to be deleted has two children**: Find the smallest value in the right subtree (in-order successor) and replace the node's value with that. Then, delete the in-order successor.

#### Time Complexity

- **Best and Average Case**: O(log n) – Deletion can be efficiently performed if the tree is balanced.
- **Worst Case**: O(n) – In an unbalanced tree, deletion can take linear time.

### Visualization

Below is a simple visualization of the tree operations:

#### Initial Tree:
```
      2
     / \
    1   3
```

#### After Deleting `2`:
```
      3
     /
    1
```