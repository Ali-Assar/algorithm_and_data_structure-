# LinkedList Implementation in Go and Python

This repository contains implementations of a singly linked list in both Go and Python. Linked lists are a fundamental data structure in computer science, and this repository provides examples of their implementation, along with explanations of the key operations and their algorithmic complexity.

## What is a Linked List?

A linked list is a linear data structure where elements are stored in nodes. Each node contains two parts:

1. **Value**: The data or value stored in the node.
2. **Next**: A reference (or pointer) to the next node in the sequence.

### Types of Linked Lists

- **Singly Linked List**: Each node points to the next node in the list. The last node points to `nil` (or `None` in Python), indicating the end of the list.
- **Doubly Linked List**: Each node contains an additional pointer to the previous node, allowing traversal in both directions.
- **Circular Linked List**: The last node points back to the first node, forming a loop.

This repository focuses on a **singly linked list**.

## Key Operations and Their Complexity

### 1. Append

**Operation**: Add an element to the end of the list.

**Algorithm**:
- Create a new node.
- If the list is empty, set both `head` and `tail` to the new node.
- Otherwise, set the `next` pointer of the current `tail` to the new node and update `tail` to the new node.
- Increment the length of the list.

**Complexity**: O(1)

Appending an element to the end of the list is a constant time operation because we maintain a pointer to the `tail`.

### 2. Prepend

**Operation**: Add an element to the beginning of the list.

**Algorithm**:
- Create a new node.
- If the list is empty, set both `head` and `tail` to the new node.
- Otherwise, set the `next` pointer of the new node to the current `head` and update `head` to the new node.
- Increment the length of the list.

**Complexity**: O(1)

Prepending an element is a constant time operation since we only need to update the `head` pointer.

### 3. Pop

**Operation**: Remove the last element from the list.

**Algorithm**:
- If the list is empty, return `nil`.
- Traverse the list to find the second-to-last node.
- Update the `tail` to point to the second-to-last node and set its `next` to `nil`.
- If the list becomes empty after this operation, set both `head` and `tail` to `nil`.
- Decrement the length of the list.

**Complexity**: O(n)

Popping the last element requires traversing the list to find the second-to-last element, which takes linear time.

### 4. Pop First

**Operation**: Remove the first element from the list.

**Algorithm**:
- If the list is empty, return `nil`.
- Update `head` to point to the second node.
- If the list becomes empty after this operation, set `tail` to `nil`.
- Decrement the length of the list.

**Complexity**: O(1)

Popping the first element is a constant time operation since we only update the `head` pointer.

### 5. Get

**Operation**: Retrieve the value of the node at a specific index.

**Algorithm**:
- If the index is out of bounds, return `nil`.
- Start from the `head` and traverse the list until the specified index is reached.
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
- Insert the new node between the found node and the next node.
- Increment the length of the list.

**Complexity**: O(n)

Inserting an element at a specific index may require traversing the list, resulting in linear time complexity.

### 8. Remove

**Operation**: Remove a node at a specified index.

**Algorithm**:
- If the index is out of bounds, return `nil`.
- If the index is 0, use `PopFirst` to remove the node.
- If the index is equal to the length of the list minus one, use `Pop` to remove the node.
- Otherwise, use the `Get` operation to find the node just before the specified index.
- Remove the node by adjusting the pointers.
- Decrement the length of the list.

**Complexity**: O(n)

Removing an element at a specific index may require traversing the list, resulting in linear time complexity.

### 9. Reverse

**Operation**: Reverse the order of the linked list.

**Algorithm**:
- Initialize three pointers: `before` (set to `nil`), `temp` (set to `head`), and `after` (set to `temp.next`).
- Iterate through the list and reverse the `next` pointers.
- Swap the `head` and `tail` pointers.

**Complexity**: O(n)

Reversing the linked list requires a single pass through the list, resulting in linear time complexity.

## Code Availability

This repository provides the linked list implementation in both Go and Python. The directory structure is as follows:

```
/go/
    linkedlist.go  // Go implementation of LinkedList
/python/
    linkedlist.py  // Python implementation of LinkedList
```

Both implementations follow the same logic and structure, allowing you to compare and contrast how linked lists are implemented in these two languages.

## Conclusion

Understanding the operations and their complexities is crucial for leveraging linked lists effectively. Linked lists offer dynamic memory allocation and efficient insertion and deletion at the beginning of the list. However, accessing and searching elements in a linked list may be less efficient compared to arrays due to the O(n) complexity for these operations.



