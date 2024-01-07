package main

import (
	"testing"
)

func TestInsert(t *testing.T) {
	tests := []struct {
		key   int
		value int
	}{
		{1, 10},
		{2, 20},
		{3, 30},
		{10, 100}, // testing a key that hashes to the same index as another key
		{20, 200}, // testing a key that hashes to the same index as another key
		{30, 300}, // testing a key that hashes to the same index as another key
	}

	hashTable := NewHashTable(10)

	for _, tt := range tests {
		hashTable.Insert(tt.key, tt.value)
		node := hashTable.Search(tt.key)
		if node == nil {
			t.Errorf("Inserted key %v not found", tt.key)
		} else if node.Value != tt.value {
			t.Errorf("Expected value %v, got %v", tt.value, node.Value)
		}
	}
}