# Basic Sorting Algorithms: Bubble Sort, Selection Sort, and Insertion Sort

In this document, we cover three basic sorting algorithms: **Bubble Sort**, **Selection Sort**, and **Insertion Sort**. Each of these algorithms operates on lists in place and has different advantages and time complexities.

## 1. Bubble Sort

Bubble Sort repeatedly steps through the list, compares adjacent elements, and swaps them if they are in the wrong order. This process is repeated until the list is sorted.

### Time Complexity

- **Best Case**: O(n) (if the list is already sorted)
- **Worst and Average Case**: O(n²)

### Flowchart of Bubble Sort

```
+---------------------+
| Start with unsorted  |
| list                |
+---------------------+
        |
        v
+---------------------+
| Compare adjacent    |
| elements, swap if   |
| needed              |
+---------------------+
        |
        v
+---------------------+
| Repeat for all      |
| elements            |
+---------------------+
        |
        v
+---------------------+
| Stop if no swaps    |
| were made in a pass |
+---------------------+
```

---

## 2. Selection Sort

Selection Sort divides the list into two parts: a sorted and an unsorted part. It repeatedly selects the smallest element from the unsorted portion and swaps it with the first unsorted element.

### Time Complexity

- **Best, Worst, and Average Case**: O(n²)

### Flowchart of Selection Sort

```
+---------------------+
| Find the smallest    |
| element in the array |
+---------------------+
        |
        v
+---------------------+
| Swap it with the     |
| first unsorted       |
| element              |
+---------------------+
        |
        v
+---------------------+
| Repeat for each      |
| unsorted element     |
+---------------------+
```

---

## 3. Insertion Sort

Insertion Sort builds the sorted list one item at a time. It takes each element and places it in its correct position relative to the already sorted portion of the list.

### Time Complexity

- **Best Case**: O(n) (already sorted list)
- **Worst and Average Case**: O(n²

)

### Flowchart of Insertion Sort

```
+---------------------+
| Start from the       |
| second element       |
+---------------------+
        |
        v
+---------------------+
| Compare with all     |
| previous elements    |
| and insert in the    |
| correct position     |
+---------------------+
        |
        v
+---------------------+
| Repeat for each      |
| element              |
+---------------------+
```

---

### Conclusion

Each of these algorithms has its own use case:
- **Bubble Sort**: Simple but inefficient.
- **Selection Sort**: More efficient but still O(n²).
- **Insertion Sort**: Efficient for nearly sorted data.

All of these sorts operate in O(n²) time in the worst case, but Insertion Sort performs better on smaller or nearly sorted lists.
