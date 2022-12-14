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
 * Check if the sparse matrix is empty
 */
func (m *SparseMatrix) IsEmpty() bool {
	return m.head == nil
}

/**
 * Get size of the sparse matrix (rows and columns)
 */
func (m *SparseMatrix) Size() (int, int) {
	return m.rows, m.cols
}

/**
 * Get the number of rows in the sparse matrix
 */
func (m *SparseMatrix) Rows() int {
	return m.rows
}

/**
 * Get the number of columns in the sparse matrix
 */
func (m *SparseMatrix) Cols() int {
	return m.cols
}

/**
 * Has a value at a specific row and column
 */
func (m *SparseMatrix) Has(row, col int) bool {
	current := m.head
	for current != nil {
		if current.row == row && current.col == col {
			return true
		}
		current = current.next
	}
	return false
}

/**
 * Has an index in the real matrix
 */
func (m *SparseMatrix) HasIndex(index int) bool {
	current := m.head
	for current != nil {
		if current.row*m.cols+current.col == index {
			return true
		}
		current = current.next
	}
	return false
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

/**
 * Get the value at a specific index, if not found return nil
 */
func (m *SparseMatrix) GetIndex(index int) *float64 {
	current := m.head
	for current != nil {
		if current.row*m.cols+current.col == index {
			return &current.value
		}
		current = current.next
	}
	return nil
}

/**
 * + Operator
 */
func (m *SparseMatrix) AddMatrix(m2 *SparseMatrix) *SparseMatrix {
	if m.rows != m2.rows || m.cols != m2.cols {
		return nil
	}
	result := NewSparseMatrix(m.rows, m.cols)
	current := m.head
	for current != nil {
		result.Add(current.row, current.col, current.value+m2.Get(current.row, current.col))
		current = current.next
	}
	return result
}

/**
 * - Operator
 */
func (m *SparseMatrix) SubMatrix(m2 *SparseMatrix) *SparseMatrix {
	if m.rows != m2.rows || m.cols != m2.cols {
		return nil
	}
	result := NewSparseMatrix(m.rows, m.cols)
	current := m.head
	for current != nil {
		result.Add(current.row, current.col, current.value-m2.Get(current.row, current.col))
		current = current.next
	}
	return result
}

/**
 * * Operator
 */
func (m *SparseMatrix) MulMatrix(m2 *SparseMatrix) *SparseMatrix {
	if m.cols != m2.rows {
		return nil
	}
	result := NewSparseMatrix(m.rows, m2.cols)
	current := m.head
	for current != nil {
		for i := 0; i < m2.cols; i++ {
			result.Add(current.row, i, result.Get(current.row, i)+current.value*m2.Get(current.col, i))
		}
		current = current.next
	}
	return result
}

/**
 * / Operator
 */
func (m *SparseMatrix) DivMatrix(m2 *SparseMatrix) *SparseMatrix {
	if m.cols != m2.rows {
		return nil
	}
	result := NewSparseMatrix(m.rows, m2.cols)
	current := m.head
	for current != nil {
		for i := 0; i < m2.cols; i++ {
			result.Add(current.row, i, result.Get(current.row, i)+current.value/m2.Get(current.col, i))
		}
		current = current.next
	}
	return result
}
