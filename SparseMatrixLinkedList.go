package main

import "fmt"

// Storing a sparse matrix
type Node struct {
	row, col int
	value    float64
	next     *Node
}

type SparseMatrix struct {
	head       *Node
	rows, cols int
}

/**
 * Create a new sparse matrix
 */
func NewSparseMatrix(rows, cols int) *SparseMatrix {
	return &SparseMatrix{rows: rows, cols: cols}
}

/**
 * Add a new value to the sparse matrix
 */
func (m *SparseMatrix) Add(row, col int, value float64) {
	node := &Node{row: row, col: col, value: value}
	if m.head == nil {
		m.head = node
	} else {
		current := m.head
		for current.next != nil {
			current = current.next
		}
		current.next = node
	}
}

/**
 * Get the value at a specific row and column
 */
func (m *SparseMatrix) Get(row, col int) float64 {
	current := m.head
	for current != nil {
		if current.row == row && current.col == col {
			return current.value
		}
		current = current.next
	}
	return 0
}

/**
 * Print the sparse matrix
 */
func (m *SparseMatrix) Print() {
	for i := 0; i < m.rows; i++ {
		for j := 0; j < m.cols; j++ {
			fmt.Printf("%f ", m.Get(i, j))
		}
		fmt.Println()
	}
}
