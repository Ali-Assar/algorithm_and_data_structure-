# DoublyLinkedList Implementation in Go and Python

This folder of repository contains implementations of a doubly linked list in both Go and Python. Linked lists are a fundamental data structure in computer science, and this folder of repository provides examples of their implementation, along with explanations of the key operations and their algorithmic complexity.

## What is a Doubly Linked List?

A doubly linked list is a linear data structure where elements are stored in nodes. Each node contains three parts:

1. **Value**: The data or value stored in the node.
2. **Next**: A reference (or pointer) to the next node in the sequence.
3. **Prev**: A reference (or pointer) to the previous node in the sequence.

### Types of Linked Lists

- **Singly Linked List**: Each node points to the next node in the list. The last node points to `nil` (or `None` in Python), indicating the end of the list.
- **Doubly Linked List**: Each node contains an additional pointer to the previous node, allowing traversal in both directions.
- **Circular Linked List**: The last node points back to the first node, forming a loop.

This folder of repository focuses on a **doubly linked list**.

## Key Operations and Their Complexity

### 1. Append

**Operation**: Add an element to the end of the list.

**Algorithm**:
- Create a new node.
- If the list is empty, set both `head` and `tail` to the new node.
- Otherwise, set the `next` pointer of the current `tail` to the new node, set the `prev` pointer of the new node to the current `tail`, and update `tail` to the new node.
- Increment the length of the list.

**Complexity**: O(1)

Appending an element to the end of the list is a constant time operation because we maintain a pointer to the `tail`.

### 2. Prepend

**Operation**: Add an element to the beginning of the list.

**Algorithm**:
- Create a new node.
- If the list is empty, set both `head` and `tail` to the new node.
- Otherwise, set the `next` pointer of the new node to the current `head`, set the `prev` pointer of the current `head` to the new node, and update `head` to the new node.
- Increment the length of the list.

**Complexity**: O(1)

Prepending an element is a constant time operation since we only need to update the `head` pointer.

### 3. Pop

**Operation**: Remove the last element from the list.

**Algorithm**:
- If the list is empty, return `nil`.
- If the list has only one element, set both `head` and `tail` to `nil`.
- Otherwise, set `tail` to the previous node, update the `next` pointer of the new `tail` to `nil`, and clear the `prev` pointer of the removed node.
- Decrement the length of the list.

**Complexity**: O(1)

Popping the last element is a constant time operation because we maintain a pointer to the `tail`.

### 4. Pop First

**Operation**: Remove the first element from the list.

**Algorithm**:
- If the list is empty, return `nil`.
- If the list has only one element, set both `head` and `tail` to `nil`.
- Otherwise, set `head` to the next node, update the `prev` pointer of the new `head` to `nil`, and clear the `next` pointer of the removed node.
- Decrement the length of the list.

**Complexity**: O(1)

Popping the first element is a constant time operation since we only update the `head` pointer.

### 5. Get

**Operation**: Retrieve the value of the node at a specific index.

**Algorithm**:
- If the index is out of bounds, return `nil`.
- If the index is closer to the `head`, start from `head` and traverse the list until the specified index is reached.
- If the index is closer to the `tail`, start from `tail` and traverse the list backwards until the specified index is reached.
- Return the node at that index.

**Complexity**: O(n)

Getting the value at a specific index requires traversing the list, resulting in linear time complexity.

### 6. SetValue

**Operation**: Update the value of a node at a specific index.

**Algorithm**:
- Use the `Get` operation to retrieve the node at the specified index.
- If the node exists, update its value.

**Complexity**: O(n)

Setting a value requires finding the node at a specific index, which takes linear time.

### 7. Insert

**Operation**: Insert a new node at a specified index.

**Algorithm**:
- If the index is out of bounds, return `false`.
- If the index is 0, prepend the value.
- If the index is equal to the length of the list, append the value.
- Otherwise, use the `Get` operation to find the node just before the specified index.
- Insert the new node between the found node and the next node, adjusting the `next` and `prev` pointers accordingly.
- Increment the length of the list.

**Complexity**: O(n)

Inserting an element at a specific index may require traversing the list, resulting in linear time complexity.

### 8. Remove

**Operation**: Remove a node at a specified index.

**Algorithm**:
- If the index is out of bounds, return `nil`.
- If the index is 0, use `PopFirst` to remove the node.
- If the index is equal to the length of the list minus one, use `Pop` to remove the node.
- Otherwise, use the `Get` operation to find the node at the specified index.
- Remove the node by adjusting the `next` and `prev` pointers of the neighboring nodes, and clearing the pointers of the removed node.
- Decrement the length of the list.

**Complexity**: O(n)

Removing an element at a specific index may require traversing the list, resulting in linear time complexity.

## Code Availability

This repository provides the doubly linked list implementation in both Go and Python. The directory structure is as follows:

```
/go/
    doubly_linked_list /doubly_linked_list.go  // Go implementation of DoublyLinkedList
/python/
    doubly_linked_list/doubly_linked_list.py  // Python implementation of DoublyLinkedList
```

Both implementations follow the same logic and structure, allowing you to compare and contrast how doubly linked lists are implemented in these two languages.

## Conclusion

Understanding the operations and their complexities is crucial for leveraging doubly linked lists effectively. Doubly linked lists offer dynamic memory allocation and efficient insertion and deletion at both ends of the list. They provide additional flexibility over singly linked lists by allowing traversal in both directions, though at the cost of increased memory usage due to the additional pointer for the previous node.
```