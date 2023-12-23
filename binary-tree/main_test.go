package main

import (
	"reflect"
	"testing"
)

func TestInsertAndFind(t *testing.T) {
	tests := []struct {
		name     string
		inserts  []int
		finds    []int
		expected []bool
	}{
		{
			name:     "Test 1",
			inserts:  []int{5, 3, 2, 4, 7, 6, 8},
			finds:    []int{5, 3, 2, 4, 7, 6, 8, 9},
			expected: []bool{true, true, true, true, true, true, true, false},
		},
		{
			name:     "Test 2",
			inserts:  []int{1, 2, 3, 4, 5},
			finds:    []int{1, 2, 3, 4, 5, 6},
			expected: []bool{true, true, true, true, true, false},
		},
		{
			name:     "Test 3",
			inserts:  []int{5, 4, 3, 2, 1},
			finds:    []int{5, 4, 3, 2, 1, 0},
			expected: []bool{true, true, true, true, true, false},
		},
		{
			name:     "Test 4 - Empty Tree",
			inserts:  []int{},
			finds:    []int{1},
			expected: []bool{false},
		},
		{
			name:     "Test 5 - Single Node Tree",
			inserts:  []int{1},
			finds:    []int{1, 2},
			expected: []bool{true, false},
		},
		{
			name:     "Test 6 - Non-Existent Value",
			inserts:  []int{1, 2, 3, 4, 5},
			finds:    []int{6},
			expected: []bool{false},
		},
		{
			name:     "Test 7 - Duplicate Values",
			inserts:  []int{1, 2, 2, 3, 3, 3},
			finds:    []int{2, 3},
			expected: []bool{true, true},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := &Tree{}
			for _, v := range tt.inserts {
				tr.Insert(v)
			}
			for i, v := range tt.finds {
				node := tr.Find(v)
				found := node != nil
				if found != tt.expected[i] {
					t.Errorf("got %v, want %v", found, tt.expected[i])
				}
			}
		})
	}
}

func TestNext(t *testing.T) {
	tests := []struct {
		name     string
		inserts  []int
		value    int
		expected *int
	}{
		{
			name:     "Test 1",
			inserts:  []int{5, 3, 2, 4, 7, 6, 8},
			value:    5,
			expected: intPtr(6),
		},
		{
			name:     "Test 2",
			inserts:  []int{1, 2, 3, 4, 5},
			value:    3,
			expected: intPtr(4),
		},
		{
			name:     "Test 3",
			inserts:  []int{5, 4, 3, 2, 1},
			value:    1,
			expected: intPtr(2),
		},
		{
			name:     "Test 4 - Non-Existent Value",
			inserts:  []int{1, 2, 3, 4, 5},
			value:    6,
			expected: nil,
		},
		{
			name:     "Test 5 - Maximum Value",
			inserts:  []int{1, 2, 3, 4, 5},
			value:    5,
			expected: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := &Tree{}
			for _, v := range tt.inserts {
				tr.Insert(v)
			}
			node := tr.Next(tt.value)
			var got *int
			if node != nil {
				got = &node.Value
			}
			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("got %v, want %v", got, tt.expected)
			}
		})
	}
}

// intPtr is a helper function to get a pointer to an int
func intPtr(i int) *int {
	return &i
}
