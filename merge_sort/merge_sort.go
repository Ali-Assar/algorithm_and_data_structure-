package main

import "fmt"

func merge(list1, list2 []int) []int {
	combined := make([]int, 0, len(list1)+len(list2))
	i, j := 0, 0

	for i < len(list1) && j < len(list2) {
		if list1[i] < list2[j] {
			combined = append(combined, list1[i])
			i++
		} else {
			combined = append(combined, list2[j])
			j++
		}
	}

	combined = append(combined, list1[i:]...)
	combined = append(combined, list2[j:]...)
	return combined
}

func mergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	midIndex := len(arr) / 2
	left := mergeSort(arr[:midIndex])
	right := mergeSort(arr[midIndex:])

	return merge(left, right)
}

func main() {
	arr := []int{1, 2, 7, 8, 3, 10, 11, 4, 9, 5, 6}
	sortedArr := mergeSort(arr)
	fmt.Println(sortedArr)
}
