# Graph Data Structure

## Introduction

A graph is a collection of nodes (or vertices) and edges that connect pairs of nodes. Graphs can be represented in multiple ways, including adjacency lists and adjacency matrices. 

This document will cover:

1. **Adjacency List** (implemented in the provided Python code).
2. **Adjacency Matrix** (theoretical discussion).
3. **Time and space complexity** of both approaches.

### 1. Adjacency List

#### What is an Adjacency List?

In an adjacency list, each node (or vertex) in the graph stores a list of the other nodes it is connected to. This makes it a space-efficient representation, particularly when the graph is sparse (i.e., has fewer edges relative to the number of vertices).

#### Time Complexity

- **Adding a vertex**: O(1) - Constant time because we are simply adding an entry to a dictionary.
- **Adding an edge**: O(1) - Since we are appending to a list in the dictionary, it takes constant time.
- **Removing an edge**: O(n) - To remove an edge, we need to find and remove it from the lists of both vertices, which takes linear time in the size of the list.
- **Removing a vertex**: O(n) - We need to delete the vertex and also remove it from the adjacency lists of other vertices, which takes time proportional to the number of edges.

#### Space Complexity

- **Adjacency List**: O(V + E), where V is the number of vertices and E is the number of edges. This is more efficient for sparse graphs.

---

### 2. Adjacency Matrix

#### What is an Adjacency Matrix?

In an adjacency matrix, a graph is represented as a 2D matrix of size `V x V`, where `V` is the number of vertices. Each cell at position `[i][j]` represents whether there is an edge between vertex `i` and vertex `j`.

If an edge exists, the value is `1` (or the weight of the edge), otherwise, it is `0`.

#### Time Complexity

- **Adding an edge**: O(1) - Constant time because we are simply updating a matrix cell.
- **Removing an edge**: O(1) - Constant time to update the matrix.
- **Space Complexity**: O(V^2), where V is the number of vertices. This can be inefficient for large, sparse graphs, as we need space for every possible pair of vertices, even if many pairs are not connected.

#### Example of an Adjacency Matrix (for understanding):
For a graph with vertices `A, B, C`, if we have edges `(A, B)` and `(B, C)`, the matrix would look like:

|    | A | B | C |
|----|---|---|---|
| A  | 0 | 1 | 0 |
| B  | 1 | 0 | 1 |
| C  | 0 | 1 | 0 |

---

### 3. Comparison of Adjacency List vs Matrix

| Operation       | Adjacency List       | Adjacency Matrix  |
|-----------------|----------------------|-------------------|
| **Add Vertex**  | O(1)                 | O(V^2)            |
| **Add Edge**    | O(1)                 | O(1)              |
| **Remove Edge** | O(n)                 | O(1)              |
| **Space**       | O(V + E)             | O(V^2)            |
| **Best for**    | Sparse graphs        | Dense graphs      |

---

### Visualization

For an adjacency list, a visual representation of the graph:

```
A --> B, C, D
B --> A, D, E
C --> A, D, E
D --> A, B, C
E --> B, C
```

This shows that each vertex points to a list of its connected neighbors.

---

## Conclusion

Adjacency lists are more space-efficient for sparse graphs, which makes them a better choice in most cases. Adjacency matrices, on the other hand, can be useful for dense graphs where the number of edges is close to the square of the number of vertices.

---

### Conversion to Go

Here’s the Go implementation of the same adjacency list graph:
