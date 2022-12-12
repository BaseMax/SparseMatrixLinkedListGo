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
