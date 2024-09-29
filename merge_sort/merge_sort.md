# Merge Sort Algorithm

Merge Sort is an efficient, stable, and comparison-based sorting algorithm. It works on the principle of **divide and conquer**, recursively breaking down a list into smaller sublists, sorting those sublists, and then merging them back together.

## Steps of the Algorithm

1. **Divide**: Split the list into two halves until each sublist has one element.
2. **Conquer**: Recursively sort each sublist.
3. **Merge**: Combine the two sorted sublists into a single sorted list.

## Flowchart of Merge Sort

Below is a simplified flowchart of the Merge Sort process:

```
+------------------+
|   Start          |
+------------------+
       |
       v
+------------------+
| Divide the array |
| into two halves  |
+------------------+
       |
       v
+------------------------+
| Recursively sort both  |
| the left and right     |
| halves                |
+------------------------+
       |
       v
+------------------------+
| Merge the sorted halves|
| into a single sorted   |
| list                  |
+------------------------+
       |
       v
+------------------+
|   End            |
+------------------+
```

## Time Complexity (Big O)

### Best Case, Average Case, and Worst Case: **O(n log n)**

Merge Sort always divides the list into two halves, performs recursive sorting on both halves, and merges them in linear time. As it always requires `log n` levels of recursion, the time complexity is `O(n log n)` in all cases.

### Space Complexity: **O(n)**

Merge Sort requires additional space proportional to the size of the array being sorted, making its space complexity `O(n)`.

## Advantages of Merge Sort

- **Stable**: Merge Sort maintains the relative order of records with equal keys.
- **Predictable performance**: Its `O(n log n)` time complexity applies to the best, average, and worst cases.
- **Suitable for linked lists**: Merge Sort works well with linked lists due to its sequential access pattern.

## Disadvantages of Merge Sort

- **Space complexity**: Merge Sort uses `O(n)` extra space, making it less memory-efficient compared to in-place algorithms like QuickSort.
