package main

import "fmt"

func bubbleSort(arr []int) []int {
	n := len(arr)
	for i := n - 1; i > 0; i-- {
		swapped := false
		for j := 0; j < i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				swapped = true
			}
		}
		if !swapped {
			break
		}
	}
	return arr
}

func selectionSort(arr []int) []int {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		minIndex := i
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}
		if i != minIndex {
			arr[i], arr[minIndex] = arr[minIndex], arr[i]
		}
	}
	return arr
}

func insertionSort(arr []int) []int {
	for i := 1; i < len(arr); i++ {
		currentValue := arr[i]
		j := i - 1
		for j >= 0 && arr[j] > currentValue {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = currentValue
	}
	return arr
}

func main() {
	arr := []int{5, 8, 1, 6, 7, 4, 3}
	
	fmt.Println(bubbleSort(arr))
	fmt.Println(selectionSort(arr))
	fmt.Println(insertionSort(arr))
}
