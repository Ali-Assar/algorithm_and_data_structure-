# Hash Table Implementation in Go and Python

This repository contains an implementation of the Hash Table data structure in both Go and Python. A Hash Table (also known as a Hash Map) is a data structure that implements an associative array abstract data type, a structure that can map keys to values. Hash Tables are highly efficient, providing fast access to data via unique keys.

## What is a Hash Table?

A Hash Table uses a **hash function** to compute an index (also known as a hash code) into an array of buckets or slots, from which the desired value can be found. 

- **Hash Function**: A function that takes a key as input and produces an index in the array where the value associated with the key is stored.
- **Collision**: When two different keys produce the same hash code, it's called a collision. In our implementation, we resolve this using **chaining**, where each bucket contains a list of key-value pairs.

### Why Use a Hash Table?

- **Fast Lookup**: Hash tables provide O(1) average time complexity for insert, delete, and lookup operations. In the worst-case scenario (with many collisions), the time complexity can degrade to O(n), but with a good hash function, this is rare.
- **Efficient Data Storage**: They are ideal for scenarios where fast access to key-value pairs is required, like databases, caching mechanisms, and more.

## Key Operations and Their Complexity

### 1. Insert (Set Item)

**Operation**: Add a key-value pair to the Hash Table.

**Explanation**: The key is passed through the hash function to compute an index. If there's no collision (i.e., no other key has hashed to this index), the key-value pair is stored directly in the array. In case of a collision, the new key-value pair is appended to the list in that bucket.

**Algorithm**:
- Compute the hash of the key using the hash function.
- Store the key-value pair at the computed index.
- If there's already an item at this index (collision), use chaining (a list) to store the key-value pair.

**Complexity**: 
- Best Case: O(1)
- Worst Case: O(n) (in case of many collisions)

### 2. Retrieve (Get Item)

**Operation**: Retrieve the value associated with a given key.

**Explanation**: The key is passed through the hash function to compute the index. If a list exists at the computed index (due to collisions), it searches through the list for the key. Once found, the corresponding value is returned.

**Algorithm**:
- Compute the hash of the key.
- Search the list at the computed index for the key.
- Return the associated value if the key is found.

**Complexity**: 
- Best Case: O(1)
- Worst Case: O(n) (if there's a long chain of items at one index)

### 3. Retrieve All Keys

**Operation**: Retrieve a list of all the keys stored in the Hash Table.

**Explanation**: This operation iterates over each bucket in the table and collects all the keys present in the non-empty buckets.

**Algorithm**:
- Traverse the hash table.
- Collect keys from each non-empty bucket.

**Complexity**: O(n), where `n` is the number of key-value pairs in the table.

## Code Structure

```
/go/
    hst.go   // Go implementation of Hash Table
/python/
    hst.py   // Python implementation of Hash Table
```

Both implementations provide the same functionality and follow similar logic, offering an opportunity to compare the two languages in how they handle data structures.

## Conclusion

Hash Tables are one of the most powerful and efficient data structures for mapping keys to values. By understanding and implementing Hash Tables, you gain insight into how many modern applications, like caching and data indexing, work behind the scenes. Understanding hash functions and how to handle collisions is essential for optimizing the performance of hash tables.
```