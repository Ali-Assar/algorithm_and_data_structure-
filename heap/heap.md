# MaxHeap Data Structure

## Introduction

A **heap** is a specialized tree-based data structure that satisfies the heap property:

- **MaxHeap**: For every parent node, its value is greater than or equal to the values of its children.
- **MinHeap**: For every parent node, its value is less than or equal to the values of its children.

In a **MaxHeap**, the largest element is always at the root, which makes it useful for implementing priority queues where the highest priority element is dequeued first.

### Why Use a Heap Instead of Other Data Structures?

Heaps are especially useful for scenarios that involve **efficient access to the maximum (or minimum) element**, particularly when we need frequent insertions and deletions. Here’s a comparison of heaps with other common data structures:

1. **Binary Search Trees (BST)**:
   - In a balanced BST, both search and insertion take **O(log n)** time. However, finding the maximum element takes **O(n)** in the worst case for an unbalanced BST.
   - In a MaxHeap, the maximum element is always at the root, so **accessing the maximum element** is **O(1)**, and insertions/removals are **O(log n)** due to the tree being balanced by default.

2. **Linked Lists**:
   - If we use an **unsorted linked list**, insertion is **O(1)**, but finding the maximum element requires a complete traversal, making it **O(n)**.
   - If we use a **sorted linked list**, finding the maximum is **O(1)**, but insertions take **O(n)** due to the need to find the correct position for the new element.
   - A heap offers a balance between both: accessing the maximum is **O(1)**, and insertions are more efficient than a sorted list, at **O(log n)**.

3. **Arrays**:
   - Finding the maximum element in an **unsorted array** takes **O(n)** time.
   - In a **sorted array**, finding the maximum is **O(1)**, but inserting a new element takes **O(n)** to maintain the order.
   - Heaps allow for **O(log n)** insertions and **O(1)** access to the maximum element, which is more efficient for certain use cases.

### Use Cases for Heaps

1. **Priority Queues**: Heaps are ideal for implementing priority queues, where you need to efficiently extract the element with the highest priority. This is commonly used in scheduling algorithms, such as job scheduling in operating systems.

2. **Dijkstra's Algorithm**: Heaps are used to efficiently select the next vertex with the smallest known distance in the shortest-path algorithm for graphs.

3. **Heap Sort**: A heap can be used to implement **heap sort**, which has a time complexity of **O(n log n)**, making it a competitive sorting algorithm.

4. **Merging Sorted Files**: Heaps can be used to merge multiple sorted input files or streams efficiently. This is used in external sorting algorithms, where you work with data too large to fit into memory.

### MaxHeap Operations

1. **Insertion**: When inserting a new value, add it to the end of the heap and then "bubble up" the value by comparing it with its parent until the heap property is restored.

2. **Removal (Extract Max)**: Remove the maximum value (the root), replace it with the last value in the heap, and "sink down" the root to maintain the heap property.

### Heap Visual Representation

For example, if you insert the following values into the MaxHeap: `[99, 1, 22, 23, 50, 100, 125]`, the heap structure might look like this:

```
        125
       /   \
     100    99
    /  \    / \
  23   50  22  1
```

After removing the max (125), the heap will adjust and look like this:

```
        100
       /   \
     50     99
    /  \    /
  23   1  22
```

### Time Complexity Summary

| Operation  | Time Complexity |
|------------|-----------------|
| Insert     | O(log n)        |
| Remove     | O(log n)        |
| Peek (Max) | O(1)            |

### Space Complexity
- **O(n)**: The heap requires space proportional to the number of elements stored in it.

