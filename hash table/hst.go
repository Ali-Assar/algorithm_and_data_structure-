package main

import (
	"fmt"
)

type HashTable struct {
	dataMap [][][2]string
	size    int
}

// Hash function for the table
func hash(key string, size int) int {
	myHash := 0
	for _, letter := range key {
		myHash = (myHash + int(letter)*23) % size
	}
	return myHash
}

// Initialize the HashTable
func NewHashTable(size int) *HashTable {
	return &HashTable{dataMap: make([][][2]string, size), size: size}
}

// Set item in the table
func (ht *HashTable) SetItem(key string, value string) {
	index := hash(key, ht.size)
	if ht.dataMap[index] == nil {
		ht.dataMap[index] = make([][2]string, 0)
	}
	ht.dataMap[index] = append(ht.dataMap[index], [2]string{key, value})
}

// Get item from the table
func (ht *HashTable) GetItem(key string) string {
	index := hash(key, ht.size)
	if ht.dataMap[index] != nil {
		for _, pair := range ht.dataMap[index] {
			if pair[0] == key {
				return pair[1]
			}
		}
	}
	return ""
}

// Print the entire hash table
func (ht *HashTable) PrintTable() {
	for i, val := range ht.dataMap {
		fmt.Println(i, ":", val)
	}
}

// Get all the keys in the table
func (ht *HashTable) Keys() []string {
	allKeys := []string{}
	for _, bucket := range ht.dataMap {
		if bucket != nil {
			for _, pair := range bucket {
				allKeys = append(allKeys, pair[0])
			}
		}
	}
	return allKeys
}

// Example usage
func main() {
	myHashTable := NewHashTable(7)
	myHashTable.SetItem("bolts", "1400")
	myHashTable.SetItem("washers", "50")
	myHashTable.SetItem("lumber", "70")

	myHashTable.PrintTable()

	fmt.Println("Get 'bolts':", myHashTable.GetItem("bolts"))
	fmt.Println("Keys:", myHashTable.Keys())
}
