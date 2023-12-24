package main

import (
	"reflect"
	"testing"
)

func TestParse(t *testing.T) {
	tests := []struct {
		name        string
		expression  string
		wantNumbers []int
		wantOps     []Operation
		wantErr     bool
	}{
		{
			name:        "simple addition",
			expression:  "1 + 2 + 3",
			wantNumbers: []int{1, 2, 3},
			wantOps:     []Operation{Add, Add},
			wantErr:     false,
		},
		{
			name:        "invalid number",
			expression:  "1 + a + 3",
			wantNumbers: nil,
			wantOps:     nil,
			wantErr:     true,
		},
		{
			name:        "single number",
			expression:  "1",
			wantNumbers: []int{1},
			wantOps:     []Operation{},
			wantErr:     false,
		},
		{
			name:        "multiple additions",
			expression:  "1 + 2 + 3 + 4",
			wantNumbers: []int{1, 2, 3, 4},
			wantOps:     []Operation{Add, Add, Add},
			wantErr:     false,
		},
		{
			name:        "whitespace around numbers",
			expression:  " 1 + 2 + 3 ",
			wantNumbers: []int{1, 2, 3},
			wantOps:     []Operation{Add, Add},
			wantErr:     false,
		},
		{
			name:        "no whitespace",
			expression:  "1+2+3",
			wantNumbers: []int{1, 2, 3},
			wantOps:     []Operation{Add, Add},
			wantErr:     false,
		},
		{
			name:        "simple subtraction",
			expression:  "5 - 2",
			wantNumbers: []int{5, 2},
			wantOps:     []Operation{Subtract},
			wantErr:     false,
		},
		{
			name:        "simple multiplication",
			expression:  "2 * 3",
			wantNumbers: []int{2, 3},
			wantOps:     []Operation{Multiply},
			wantErr:     false,
		},
		{
			name:        "simple division",
			expression:  "6 / 2",
			wantNumbers: []int{6, 2},
			wantOps:     []Operation{Divide},
			wantErr:     false,
		},
		{
			name:        "division by zero",
			expression:  "1 / 0",
			wantNumbers: []int{1, 0},
			wantOps:     []Operation{Divide},
			wantErr:     false,
		},
		{
			name:        "simple exponentiation",
			expression:  "2 ^ 3",
			wantNumbers: []int{2, 3},
			wantOps:     []Operation{Exponent},
			wantErr:     false,
		},
		{
			name:        "invalid operation",
			expression:  "1 ? 2",
			wantNumbers: nil,
			wantOps:     nil,
			wantErr:     true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotNumbers, gotOps, err := Parse(tt.expression)
			if (err != nil) != tt.wantErr {
				t.Errorf("Parse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotNumbers, tt.wantNumbers) {
				t.Errorf("Parse() gotNumbers = %v, want %v", gotNumbers, tt.wantNumbers)
			}

			// Compare the results of the operations instead of the operations themselves
			if len(gotOps) != len(tt.wantOps) {
				t.Errorf("Parse() gotOps = %v, want %v", gotOps, tt.wantOps)
			} else {
				for i, gotOp := range gotOps {
					if gotOp.Name != tt.wantOps[i].Name {
						t.Errorf("Parse() gotOps = %v, want %v", gotOps, tt.wantOps)
						break
					}
				}
			}
		})
	}
}

func TestCalculate(t *testing.T) {
	tests := []struct {
		name       string
		numbers    []int
		operations []Operation
		want       int
	}{
		{
			name:       "simple addition",
			numbers:    []int{1, 2, 3},
			operations: []Operation{Add, Add},
			want:       6,
		},
		{
			name:       "single number",
			numbers:    []int{1},
			operations: []Operation{},
			want:       1,
		},
		{
			name:       "multiple additions",
			numbers:    []int{1, 2, 3, 4},
			operations: []Operation{Add, Add, Add},
			want:       10,
		},
		{
			name:       "simple subtraction",
			numbers:    []int{5, 2},
			operations: []Operation{Subtract},
			want:       3,
		},
		{
			name:       "simple multiplication",
			numbers:    []int{2, 3},
			operations: []Operation{Multiply},
			want:       6,
		},
		{
			name:       "simple division",
			numbers:    []int{6, 2},
			operations: []Operation{Divide},
			want:       3,
		},
		{
			name:       "division by zero",
			numbers:    []int{1, 0},
			operations: []Operation{Divide},
			want:       0, // or whatever your function returns when dividing by zero
		},
		{
			name:       "simple exponentiation",
			numbers:    []int{2, 3},
			operations: []Operation{Exponent},
			want:       8,
		},
		{
			name:       "mixed operations",
			numbers:    []int{2, 3, 2},
			operations: []Operation{Multiply, Add},
			want:       8,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, err := Calculate(tt.numbers, tt.operations); got != tt.want {
				if err != nil {
					t.Errorf("Calculate() error = %v", err)
				}
				t.Errorf("Calculate() = %v, want %v", got, tt.want)
			}
		})
	}
}
