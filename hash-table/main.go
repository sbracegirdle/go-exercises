package main

import (
	"fmt"
)

type Node struct {
	Key   int
	Value int
	Next  *Node
}

type HashTable struct {
	Size    int
	Buckets []*Node
}

func NewHashTable(size int) *HashTable {
	return &HashTable{
		Size:    size,
		Buckets: make([]*Node, size),
	}
}

func (h *HashTable) Insert(key int, value int) {
	index := key % h.Size
	node := &Node{Key: key, Value: value}
	node.Next = h.Buckets[index]
	h.Buckets[index] = node
}

func (h *HashTable) Search(key int) *Node {
	index := key % h.Size
	node := h.Buckets[index]
	for node != nil {
		if node.Key == key {
			return node
		}
		node = node.Next
	}
	return nil
}

func main() {
	fmt.Println("Hello, world!")
}
