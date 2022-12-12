# Sparse-Matrix Linked-List Go

Sparse-Matrix Linked-List Go is a sparse matrix implementation in Go using a linked list. This is a data structure that stores only non-zero values in a matrix. This is useful when you have a lot of zeros in your matrix. This data structure is also useful when you want to perform operations on sparse matrices.

Here, we are going to store the sparse matrix in a normal linked-list.

## Structure

**Types:**

```
// Storing a item in sparse matrix
type Node struct {
	row, col int
	value    float64
	next     *Node
}

// Storing a sparse matrix
type SparseMatrix struct {
	head       *Node
	rows, cols int
}
```

**Functions:**

- `func NewSparseMatrix(rows, cols int) *SparseMatrix`: Create a new sparse matrix

**Methods:**

- `func (m *SparseMatrix) Add(row, col int, value float64)`: Add a value to the sparse matrix
- `func (m *SparseMatrix) Get(row, col int) float64`: Get the value at a specific row and column
- `func (m *SparseMatrix) IsEmpty() bool`: Check if the sparse matrix is empty
- `func (m *SparseMatrix) Size() (int, int)`: Get size of the sparse matrix (rows and columns)
- `func (m *SparseMatrix) Rows() int`: Get the number of rows in the sparse matrix
- `func (m *SparseMatrix) Cols() int`: Get the number of columns in the sparse matrix
- `func (m *SparseMatrix) Has(row, col int) bool`: Has a value at a specific row and column
- `func (m *SparseMatrix) HasIndex(index int) bool`: Has an index in the real matrix
- `func (m *SparseMatrix) Print()`: Print the sparse matrix
- `func (m *SparseMatrix) GetIndex(index int) *float64`: Get the value at a specific index, if not found return nil
- `func (m *SparseMatrix) AddMatrix(m2 *SparseMatrix) *SparseMatrix`: Add two sparse matrices
- `func (m *SparseMatrix) SubMatrix(m2 *SparseMatrix) *SparseMatrix`: Subtract two sparse matrices
- `func (m *SparseMatrix) MulMatrix(m2 *SparseMatrix) *SparseMatrix`: Multiply two sparse matrices
- `func (m *SparseMatrix) DivMatrix(m2 *SparseMatrix) *SparseMatrix`: Divide two sparse matrices

## Example

```go
package main

import "fmt"

func main() {
	// Create a new sparse matrix
	matrix := NewSparseMatrix(3, 3)

	// Add some values
	matrix.Add(0, 0, 1)
	matrix.Add(0, 1, 2)
	matrix.Add(0, 2, 3)
	matrix.Add(1, 0, 4)
	matrix.Add(1, 1, 5)
	matrix.Add(1, 2, 6)
	matrix.Add(2, 0, 7)
	matrix.Add(2, 1, 8)
	matrix.Add(2, 2, 9)

	// Print the matrix
	matrix.Print()

	// Get the value at a specific row and column
	fmt.Println(matrix.Get(0, 0))

	// Check if the matrix is empty
	fmt.Println(matrix.IsEmpty())

	// Get the number of rows
	fmt.Println(matrix.Rows())

	// Get the number of columns
	fmt.Println(matrix.Cols())

	// Get the size of the matrix
	fmt.Println(matrix.Size())

	// Check if the matrix has a value at a specific row and column
	fmt.Println(matrix.Has(0, 0))

	// Check if the matrix has a value at a specific index
	fmt.Println(matrix.HasIndex(0))

	// Get the value at a specific index
	fmt.Println(*matrix.GetIndex(0))

	// Create a new sparse matrix
	matrix2 := NewSparseMatrix(3, 3)

	// Add some values
	matrix2.Add(0, 0, 1)
	matrix2.Add(0, 1, 2)
	matrix2.Add(0, 2, 3)
	matrix2.Add(1, 0, 4)
	matrix2.Add(1, 1, 5)
	matrix2.Add(1, 2, 6)
	matrix2.Add(2, 0, 7)
	matrix2.Add(2, 1, 8)
	matrix2.Add(2, 2, 9)

	// Print the matrix
	matrix2.Print()

	// Add two sparse matrices
	matrix3 := matrix.AddMatrix(matrix2)

	// Print the matrix
	matrix3.Print()

	// Subtract two sparse matrices
	matrix4 := matrix.SubMatrix(matrix2)

	// Print the matrix
	matrix4.Print()

	// Multiply two sparse matrices
	matrix5 := matrix.MulMatrix(matrix2)

	// Print the matrix
	matrix5.Print()

	// Divide two sparse matrices
	matrix6 := matrix.DivMatrix(matrix2)

	// Print the matrix
	matrix6.Print()
}
```

## License

This project is licensed under the GPL-3.0 License - see the [LICENSE](LICENSE) file for details.

© Copyright (c) 2022, Max Base
