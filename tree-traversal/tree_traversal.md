
# Tree Traversal Algorithms: BFS and DFS Variants

Tree traversal is the process of visiting each node in a tree data structure, exactly once. Traversal is critical in various operations like searching, inserting, updating, or deleting nodes. There are two primary ways to traverse trees:

- **Breadth-First Search (BFS)**: Also known as level-order traversal, this method explores all the nodes at each depth before moving to the next level.
- **Depth-First Search (DFS)**: This method explores as far as possible along each branch before backtracking. DFS has three main variants: **Pre-order**, **Post-order**, and **In-order**.

## 1. Breadth-First Search (BFS)

### Overview

Breadth-First Search (BFS) explores the tree level by level, starting at the root node and progressing to nodes at each level, left to right. It is implemented using a queue to keep track of nodes that are yet to be visited at each level.

### Time and Space Complexity

- **Time Complexity**: O(n), where `n` is the number of nodes in the tree.
- **Space Complexity**: O(n), due to the use of a queue for storing nodes at each level.

### Flowchart

```plaintext
+-----------------+
| Start at root   |
| node            |
+-----------------+
       |
       v
+-----------------+
| Add node to     |
| queue           |
+-----------------+
       |
       v
+-----------------+
| Dequeue node,   |
| visit it        |
+-----------------+
       |
       v
+-----------------+
| Enqueue left    |
| and right child |
| if present      |
+-----------------+
       |
       v
+-----------------+
| Repeat until    |
| queue is empty  |
+-----------------+
```

---

## 2. Depth-First Search (DFS) Variants

### Overview

In DFS, we start from the root and explore as far as possible down one branch of the tree before backtracking. DFS can be performed in three ways:
- **Pre-order**: Visit the current node before its subtrees.
- **Post-order**: Visit the current node after its subtrees.
- **In-order**: Visit the left subtree, then the current node, then the right subtree. This variant is useful for **binary search trees** since it processes nodes in sorted order.

### 2.1. DFS Pre-order

#### Description

In Pre-order DFS, the node itself is visited first, followed by its left subtree and then its right subtree. This is useful when you want to explore nodes in a top-down manner.

### 2.2. DFS Post-order

#### Description

In Post-order DFS, we visit the left subtree, then the right subtree, and finally the current node. This order is useful when you need to delete or process a node after all its children have been processed.

### 2.3. DFS In-order

#### Description

In In-order DFS, the left subtree is visited first, then the current node, and finally the right subtree. This traversal is particularly useful for **Binary Search Trees (BST)**, as it returns nodes in sorted order.

### Time and Space Complexity of DFS

- **Time Complexity**: O(n) for all three variants, since we visit every node once.
- **Space Complexity**: O(h), where `h` is the height of the tree. For balanced trees, `h = log(n)`, but for skewed trees, `h = n`.

### Flowchart of DFS Pre-order

```plaintext
+-------------------+
| Start at root     |
| node              |
+-------------------+
        |
        v
+-------------------+
| Visit node        |
+-------------------+
        |
        v
+-------------------+
| Traverse left     |
| subtree           |
+-------------------+
        |
        v
+-------------------+
| Traverse right    |
| subtree           |
+-------------------+
        |
        v
+-------------------+
| Repeat for each   |
| subtree node      |
+-------------------+
```

---

## Conclusion

- **BFS** is better suited for exploring nodes level by level, making it useful for shortest path algorithms.
- **DFS Pre-order** is useful for copying the tree or when you need to process a node before its children.
- **DFS Post-order** is typically used in applications like freeing resources or deleting nodes.
- **DFS In-order** is ideal for binary search trees, as it outputs the nodes in sorted order.

