package main

import "fmt"

type Matrix struct {
	data       [][]int
	rows, cols int
}

// NewMatrix creates a new matrix
func NewMatrix(rows, cols int) *Matrix {
	data := make([][]int, rows)
	for i := range data {
		data[i] = make([]int, cols)
	}
	return &Matrix{data, rows, cols}
}

// Add adds two matrices
func (m *Matrix) Add(n *Matrix) (*Matrix, error) {
	if m.rows != n.rows || m.cols != n.cols {
		return nil, fmt.Errorf("matrices are not the same size")
	}
	result := NewMatrix(m.rows, m.cols)
	for i := 0; i < m.rows; i++ {
		for j := 0; j < m.cols; j++ {
			result.data[i][j] = m.data[i][j] + n.data[i][j]
		}
	}
	return result, nil
}

// Multiply multiplies two matrices
func (m *Matrix) Multiply(n *Matrix) (*Matrix, error) {
	if m.cols != n.rows {
		return nil, fmt.Errorf("matrices are not compatible for multiplication")
	}
	result := NewMatrix(m.rows, n.cols)
	for i := 0; i < m.rows; i++ {
		for j := 0; j < n.cols; j++ {
			sum := 0
			for k := 0; k < m.cols; k++ {
				sum += m.data[i][k] * n.data[k][j]
			}
			result.data[i][j] = sum
		}
	}
	return result, nil
}

// Transpose returns the transpose of the matrix
func (m *Matrix) Transpose() *Matrix {
	result := NewMatrix(m.cols, m.rows)
	for i := 0; i < m.rows; i++ {
		for j := 0; j < m.cols; j++ {
			result.data[j][i] = m.data[i][j]
		}
	}
	return result
}