package main

import (
	"reflect"
	"testing"
)

func TestSearch(t *testing.T) {
	tree := createBTree([]int{1, 2, 3, 4, 5}, 3)

	tests := []struct {
		key      int
		expected bool
	}{
		{1, true},
		{2, true},
		{6, false},
		{0, false},
		{4, true},
	}

	for _, test := range tests {
		result, _ := tree.root.Search(test.key)
		if (result != nil) != test.expected {
			t.Errorf("Expected %v, got %v", test.expected, result != nil)
		}
	}
}

// Assuming a GetKeys function exists
func getKeys(node *BTreeNode, keys *[]int) {
	if node != nil {
		for _, child := range node.child {
			getKeys(child, keys)
		}

		*keys = append(*keys, node.keys...)
	}
}

func TestInsert(t *testing.T) {
	tests := []struct {
		keys     []int
		expected []int
	}{
		{[]int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}},
		{[]int{5, 4, 3, 2, 1}, []int{1, 2, 3, 4, 5}},
		{[]int{3, 1, 4, 5, 2}, []int{1, 2, 3, 4, 5}},
		{[]int{10, 20, 30, 40, 50}, []int{10, 20, 30, 40, 50}},
		{[]int{50, 40, 30, 20, 10}, []int{10, 20, 30, 40, 50}},
		{[]int{30, 10, 40, 50, 20}, []int{10, 20, 30, 40, 50}},
		{[]int{}, []int{}},
		{[]int{1}, []int{1}},
		{[]int{2, 1}, []int{1, 2}},
	}

	for _, test := range tests {
		tree := &BTree{
			root: &BTreeNode{
				isLeaf: true,
				keys:   make([]int, 0),
				child:  make([]*BTreeNode, 0),
			},
			t: 3,
		}

		for _, key := range test.keys {
			tree.Insert(key)
		}

		var keys []int
		getKeys(tree.root, &keys)

		if test.expected == nil || len(test.expected) == 0 {
			if len(keys) > 0 {
				t.Errorf("Expected %v, got %v", test.expected, keys)
			}
		} else if !reflect.DeepEqual(keys, test.expected) {
			t.Errorf("Expected %v, got %v", test.expected, keys)
		}
	}
}
