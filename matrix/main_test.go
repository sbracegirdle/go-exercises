package main

import (
	"reflect"
	"testing"
)

func TestAdd(t *testing.T) {
	tests := []struct {
		name    string
		matrix1 *Matrix
		matrix2 *Matrix
		want    *Matrix
		wantErr bool
	}{
		{
			name:    "2x2 matrices",
			matrix1: &Matrix{data: [][]int{{1, 2}, {3, 4}}, rows: 2, cols: 2},
			matrix2: &Matrix{data: [][]int{{5, 6}, {7, 8}}, rows: 2, cols: 2},
			want:    &Matrix{data: [][]int{{6, 8}, {10, 12}}, rows: 2, cols: 2},
			wantErr: false,
		},
		{
			name:    "incompatible matrices",
			matrix1: &Matrix{data: [][]int{{1, 2}, {3, 4}}, rows: 2, cols: 2},
			matrix2: &Matrix{data: [][]int{{5, 6}}, rows: 1, cols: 2},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "2x3 and 3x2 matrices",
			matrix1: &Matrix{data: [][]int{{1, 2, 3}, {4, 5, 6}}, rows: 2, cols: 3},
			matrix2: &Matrix{data: [][]int{{7, 8}, {9, 10}, {11, 12}}, rows: 3, cols: 2},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "empty matrices",
			matrix1: &Matrix{data: [][]int{{}, {}}, rows: 2, cols: 0},
			matrix2: &Matrix{data: [][]int{{}, {}}, rows: 2, cols: 0},
			want:    &Matrix{data: [][]int{{}, {}}, rows: 2, cols: 0},
			wantErr: false,
		},
		{
			name:    "matrices with negative numbers",
			matrix1: &Matrix{data: [][]int{{-1, -2}, {-3, -4}}, rows: 2, cols: 2},
			matrix2: &Matrix{data: [][]int{{5, 6}, {7, 8}}, rows: 2, cols: 2},
			want:    &Matrix{data: [][]int{{4, 4}, {4, 4}}, rows: 2, cols: 2},
			wantErr: false,
		},
		{
			name:    "matrices with zero",
			matrix1: &Matrix{data: [][]int{{0, 0}, {0, 0}}, rows: 2, cols: 2},
			matrix2: &Matrix{data: [][]int{{1, 2}, {3, 4}}, rows: 2, cols: 2},
			want:    &Matrix{data: [][]int{{1, 2}, {3, 4}}, rows: 2, cols: 2},
			wantErr: false,
		},
		{
			name:    "matrices with one element",
			matrix1: &Matrix{data: [][]int{{1}}, rows: 1, cols: 1},
			matrix2: &Matrix{data: [][]int{{2}}, rows: 1, cols: 1},
			want:    &Matrix{data: [][]int{{3}}, rows: 1, cols: 1},
			wantErr: false,
		},
		{
			name:    "3x3 matrices",
			matrix1: &Matrix{data: [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}, rows: 3, cols: 3},
			matrix2: &Matrix{data: [][]int{{10, 11, 12}, {13, 14, 15}, {16, 17, 18}}, rows: 3, cols: 3},
			want:    &Matrix{data: [][]int{{11, 13, 15}, {17, 19, 21}, {23, 25, 27}}, rows: 3, cols: 3},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.matrix1.Add(tt.matrix2)
			if (err != nil) != tt.wantErr {
				t.Errorf("Matrix.Add() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Matrix.Add() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMultiply(t *testing.T) {
	tests := []struct {
		name    string
		matrix1 *Matrix
		matrix2 *Matrix
		want    *Matrix
		wantErr bool
	}{
		{
			name:    "2x2 matrices",
			matrix1: &Matrix{data: [][]int{{1, 2}, {3, 4}}, rows: 2, cols: 2},
			matrix2: &Matrix{data: [][]int{{5, 6}, {7, 8}}, rows: 2, cols: 2},
			want:    &Matrix{data: [][]int{{19, 22}, {43, 50}}, rows: 2, cols: 2},
			wantErr: false,
		},
		{
			name:    "incompatible matrices",
			matrix1: &Matrix{data: [][]int{{1, 2}, {3, 4}}, rows: 2, cols: 2},
			matrix2: &Matrix{data: [][]int{{5, 6}}, rows: 1, cols: 2},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "2x3 and 3x2 matrices",
			matrix1: &Matrix{data: [][]int{{1, 2, 3}, {4, 5, 6}}, rows: 2, cols: 3},
			matrix2: &Matrix{data: [][]int{{7, 8}, {9, 10}, {11, 12}}, rows: 3, cols: 2},
			want:    &Matrix{data: [][]int{{58, 64}, {139, 154}}, rows: 2, cols: 2},
			wantErr: false,
		},
		{
			name:    "empty matrices",
			matrix1: &Matrix{data: [][]int{{}, {}}, rows: 2, cols: 0},
			matrix2: &Matrix{data: [][]int{{}, {}}, rows: 2, cols: 0},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "matrices with negative numbers",
			matrix1: &Matrix{data: [][]int{{-1, -2}, {-3, -4}}, rows: 2, cols: 2},
			matrix2: &Matrix{data: [][]int{{5, 6}, {7, 8}}, rows: 2, cols: 2},
			want:    &Matrix{data: [][]int{{-19, -22}, {-43, -50}}, rows: 2, cols: 2},
			wantErr: false,
		},
		{
			name:    "matrices with zero",
			matrix1: &Matrix{data: [][]int{{0, 0}, {0, 0}}, rows: 2, cols: 2},
			matrix2: &Matrix{data: [][]int{{1, 2}, {3, 4}}, rows: 2, cols: 2},
			want:    &Matrix{data: [][]int{{0, 0}, {0, 0}}, rows: 2, cols: 2},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.matrix1.Multiply(tt.matrix2)
			if (err != nil) != tt.wantErr {
				t.Errorf("Matrix.Multiply() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Matrix.Multiply() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTranspose(t *testing.T) {
	tests := []struct {
		name string
		matrix *Matrix
		want *Matrix
	}{
		{
			name: "2x2 matrix",
			matrix: &Matrix{data: [][]int{{1, 2}, {3, 4}}, rows: 2, cols: 2},
			want: &Matrix{data: [][]int{{1, 3}, {2, 4}}, rows: 2, cols: 2},
		},
		{
			name: "2x3 matrix",
			matrix: &Matrix{data: [][]int{{1, 2, 3}, {4, 5, 6}}, rows: 2, cols: 3},
			want: &Matrix{data: [][]int{{1, 4}, {2, 5}, {3, 6}}, rows: 3, cols: 2},
		},
		{
			name: "3x2 matrix",
			matrix: &Matrix{data: [][]int{{1, 2}, {3, 4}, {5, 6}}, rows: 3, cols: 2},
			want: &Matrix{data: [][]int{{1, 3, 5}, {2, 4, 6}}, rows: 2, cols: 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.matrix.Transpose(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Matrix.Transpose() = %v, want %v", got, tt.want)
			}
		})
	}
}
