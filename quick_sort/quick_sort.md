# QuickSort Algorithm

QuickSort is a highly efficient sorting algorithm that follows the **divide and conquer** strategy. It picks an element as a **pivot** and partitions the given array around the pivot such that elements smaller than the pivot are on the left side, and elements greater than the pivot are on the right side. This process is repeated recursively for the left and right subarrays.

## Steps of the Algorithm

1. **Choose a Pivot**: Select an element from the array as the pivot. In our example, we pick the first element of the array, but other strategies such as picking the last element or the middle one can be used.
   
2. **Partitioning**: Rearrange the array so that elements less than the pivot are on the left, and those greater than the pivot are on the right. 

3. **Recursion**: Recursively apply the same process to the left and right subarrays.

4. **Base Case**: The recursion ends when the subarray has fewer than two elements.

## Flowchart of QuickSort

Below is a simplified flowchart that demonstrates how QuickSort works:

```
+------------------+
|    Start         |
+------------------+
       |
       v
+------------------+
| Choose Pivot     |    
| (First element)  |
+------------------+
       |
       v
+------------------------+
| Partition Array        |
| Elements < Pivot on    |
| left, > Pivot on right |
+------------------------+
       |
       v
+---------------------+
| Recursively Sort    |
| Left and Right      |
| Subarrays           |
+---------------------+
       |
       v
+------------------+
|   Base Case:     |
|   Single Element |
+------------------+
       |
       v
+------------------+
|    End           |
+------------------+
```

This process is repeated recursively for each subarray until the entire array is sorted.

## Time Complexity (Big O)

The efficiency of QuickSort depends on the choice of the pivot and the partitioning process.

### Best Case: O(n log n)
- This occurs when the pivot divides the array into two nearly equal halves at every step.
- The array is partitioned `log n` times, and each partitioning step involves comparing `n` elements.

### Average Case: O(n log n)
- On average, the pivot will divide the array into two parts, where one is slightly larger than the other. The recursive process will still involve `log n` steps, and partitioning each step requires `n` comparisons.

### Worst Case: O(n²)
- The worst case happens when the pivot is always the smallest or largest element in the array, causing the array to be divided into a subarray of size `n-1` and a subarray of size `1`. This leads to `n` partitioning steps, with each partitioning requiring `n` comparisons.
- This situation typically occurs when the array is already sorted or reverse-sorted, and the first or last element is chosen as the pivot.

### Space Complexity: O(log n)
- QuickSort is an in-place sorting algorithm, meaning it does not require additional memory proportional to the size of the input array.
- However, the recursive calls to sort subarrays require space in the call stack, which is proportional to the depth of the recursion tree (`log n` in the best/average cases, and `n` in the worst case).

### Summary of Complexity

| Case        | Time Complexity |
|-------------|-----------------|
| Best Case   | O(n log n)       |
| Average Case| O(n log n)       |
| Worst Case  | O(n²)            |
| Space Complexity | O(log n)    |

## Advantages of QuickSort
- **Fast on average**: On average, QuickSort is faster than other sorting algorithms like MergeSort due to its lower constant factors.
- **In-place sorting**: QuickSort requires only a small amount of additional memory (O(log n) for recursion).

## Disadvantages of QuickSort
- **Worst-case performance**: QuickSort has poor worst-case performance (O(n²)), but this can be mitigated by using better pivot selection strategies like choosing a random pivot or the median.

