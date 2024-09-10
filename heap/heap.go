package main

import "fmt"

// MaxHeap structure
type MaxHeap struct {
	heap []int
}

// Returns the index of the left child
func (h *MaxHeap) leftChild(index int) int {
	return 2*index + 1
}

// Returns the index of the right child
func (h *MaxHeap) rightChild(index int) int {
	return 2*index + 2
}

// Returns the index of the parent
func (h *MaxHeap) parent(index int) int {
	return (index - 1) / 2
}

// Swaps two elements in the heap
func (h *MaxHeap) swap(index1, index2 int) {
	h.heap[index1], h.heap[index2] = h.heap[index2], h.heap[index1]
}

// Sink down operation to maintain the heap property
func (h *MaxHeap) sinkDown(index int) {
	maxIndex := index
	for {
		leftIndex := h.leftChild(index)
		rightIndex := h.rightChild(index)

		if leftIndex < len(h.heap) && h.heap[leftIndex] > h.heap[maxIndex] {
			maxIndex = leftIndex
		}
		if rightIndex < len(h.heap) && h.heap[rightIndex] > h.heap[maxIndex] {
			maxIndex = rightIndex
		}

		if maxIndex != index {
			h.swap(index, maxIndex)
			index = maxIndex
		} else {
			return
		}
	}
}

// Insert a new value into the heap
func (h *MaxHeap) Insert(value int) {
	h.heap = append(h.heap, value)
	current := len(h.heap) - 1

	for current > 0 && h.heap[current] > h.heap[h.parent(current)] {
		h.swap(current, h.parent(current))
		current = h.parent(current)
	}
}

// Remove and return the max value (root) from the heap
func (h *MaxHeap) Remove() int {
	if len(h.heap) == 0 {
		return -1 // Return -1 if the heap is empty
	}
	if len(h.heap) == 1 {
		val := h.heap[0]
		h.heap = h.heap[:0]
		return val
	}

	maxValue := h.heap[0]
	h.heap[0] = h.heap[len(h.heap)-1]
	h.heap = h.heap[:len(h.heap)-1]
	h.sinkDown(0)

	return maxValue
}

func main() {
	myHeap := MaxHeap{}

	myHeap.Insert(99)
	myHeap.Insert(1)
	myHeap.Insert(22)
	myHeap.Insert(23)
	myHeap.Insert(50)

	fmt.Println(myHeap.heap)

	myHeap.Insert(100)
	fmt.Println(myHeap.heap)

	myHeap.Insert(125)
	fmt.Println(myHeap.heap)

	myHeap.Remove()
	fmt.Println(myHeap.heap)
}
