# Stack and Queue Implementations in Go and Python

This folder of repository contains implementations of the `Stack` and `Queue` data structures in both Go and Python. Stacks and queues are fundamental data structures used in various algorithms and applications.

## What is a Stack?

A stack is a linear data structure that follows the Last In, First Out (LIFO) principle. This means that the most recently added element is the first one to be removed. Think of a stack like a pile of plates: you add (push) new plates on top of the stack, and you remove (pop) plates from the top. The last plate you added is the first one you'll take off.

### Key Operations and Their Complexity

#### 1. Push

**Operation**: Add an element to the top of the stack.

**Explanation**: In a stack, elements are added from the top because it allows for quick access to the most recent item. This structure is useful in scenarios where you need to keep track of operations or elements in the order they were added, but remove them in reverse order (e.g., undo operations).

**Algorithm**:
- Create a new node.
- Set the `next` pointer of the new node to the current top node.
- Update the top to the new node.
- Increment the height of the stack.

**Complexity**: O(1)

#### 2. Pop

**Operation**: Remove and return the top element from the stack.

**Explanation**: The pop operation removes the element from the top of the stack because it follows the LIFO principle. This is useful when you need to backtrack or reverse operations, as the most recent item is easily accessible and can be quickly removed.

**Algorithm**:
- Check if the stack is empty. If so, return `nil`.
- Store the current top node in a temporary variable.
- Update the top to the next node.
- Clear the `next` pointer of the popped node.
- Decrement the height of the stack.

**Complexity**: O(1)

### Example Usage

```python
my_stack = Stack(4)
my_stack.push(6)
my_stack.push(8)
my_stack.print_stack()

print(my_stack.pop())
my_stack.print_stack()
```

```go
myStack := NewStack(4)
myStack.Push(6)
myStack.Push(8)
myStack.PrintStack()

fmt.Println("Popped:", myStack.Pop().value)
myStack.PrintStack()
```

## What is a Queue?

A queue is a linear data structure that follows the First In, First Out (FIFO) principle. This means that the first element added to the queue will be the first one to be removed. Think of a queue like a line of people waiting to buy tickets: the first person in line is the first one to get their ticket and leave the line.

### Key Operations and Their Complexity

#### 1. Enqueue

**Operation**: Add an element to the end of the queue.

**Explanation**: In a queue, elements are added to the end because it follows the FIFO principle. This structure is useful in scenarios where you need to process elements in the order they arrive, such as tasks in a printer queue or customer service requests.

**Algorithm**:
- Create a new node.
- If the queue is empty, set both `first` and `last` to the new node.
- Otherwise, set the `next` pointer of the current last node to the new node and update `last` to the new node.
- Increment the length of the queue.

**Complexity**: O(1)

#### 2. Dequeue

**Operation**: Remove and return the first element from the queue.

**Explanation**: The dequeue operation removes the element from the front of the queue because it follows the FIFO principle. This ensures that the elements are processed in the exact order they were added, which is essential for fair and sequential processing.

**Algorithm**:
- Check if the queue is empty. If so, return `nil`.
- Store the current first node in a temporary variable.
- Update the first to the next node.
- Clear the `next` pointer of the dequeued node.
- Decrement the length of the queue.
- If the queue is empty after the operation, set `last` to `nil`.

**Complexity**: O(1)

### Example Usage

```python
my_queue = Queue(4)
my_queue.enqueue(5)
my_queue.enqueue(8)
my_queue.print_queue()

print(my_queue.dequeue())
my_queue.print_queue()
```

```go
myQueue := NewQueue(4)
myQueue.Enqueue(5)
myQueue.Enqueue(8)
myQueue.PrintQueue()

fmt.Println("Dequeued:", myQueue.Dequeue().value)
myQueue.PrintQueue()
```

## Code Availability

This repository provides the Stack and Queue implementations in both Go and Python. The directory structure is as follows:

```
/go/
    stack.go   // Go implementation of Stack
    queue/queue.go   // Go implementation of Queue
/python/
    stack.py   // Python implementation of Stack
   queue/queue.py   // Python implementation of Queue
```

Both implementations follow the same logic and structure, allowing you to compare and contrast how stacks and queues are implemented in these two languages.

## Conclusion

Stacks and queues are essential data structures widely used in algorithms, such as depth-first search (DFS) for stacks and breadth-first search (BFS) for queues. Stacks are ideal when you need to reverse the order of elements (LIFO), while queues are perfect for processing elements in the order they arrive (FIFO). Understanding their operations and complexities helps in choosing the right data structure for the right problem. Both data structures offer constant-time operations for adding and removing elements, making them highly efficient for their intended use cases.
```
