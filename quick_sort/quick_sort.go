package main

import "fmt"

func swap(arr []int, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}

func pivot(arr []int, pivotIndex, endIndex int) int {
	swapIndex := pivotIndex

	for i := pivotIndex + 1; i <= endIndex; i++ {
		if arr[i] < arr[pivotIndex] {
			swapIndex++
			swap(arr, swapIndex, i)
		}
	}
	swap(arr, pivotIndex, swapIndex)
	return swapIndex
}

func quickSortHelper(arr []int, left, right int) {
	if left < right {
		pivotIndex := pivot(arr, left, right)
		quickSortHelper(arr, left, pivotIndex-1)
		quickSortHelper(arr, pivotIndex+1, right)
	}
}

func quickSort(arr []int) {
	quickSortHelper(arr, 0, len(arr)-1)
}

func main() {
	arr := []int{1, 2, 7, 8, 3, 10, 11, 4, 9, 5, 6}
	quickSort(arr)
	fmt.Println(arr)
}
