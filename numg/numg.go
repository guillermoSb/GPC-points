package numg

import (
	"errors"
)

// Matrix Definition
type M [][]float32



// Multiplies to matrices a and b
// Returns a new matrix c
func MultiplyMatrices(a,b M) (M, error)  {
	// Cannot add matrices that are empty
	if len(a) == 0 || len(b) == 0 {
		return nil, errors.New(InvalidMatrixSizeErrorString)
	}
	// Size of matrices
	m := len(a)	// rows for matrix a
	n := len(a[0])	// number of columns for matrix a
	r := len(b[0])	// number of columns for mtrix b
	c := M{}	// Result matrix
	// Validate that can make the operation
	if m != len(b) {
		return nil, errors.New(InvalidMatrixSizeErrorString)
	}
	// Start filling the result matrix
	for i := 0; i < m; i++ {
		row := []float32{}
		for j := 0; j < r; j++ {
			entry := float32(0)
			// cij = SUM(aik bkj)
			for k := 0; k < n; k++ {
				entry += a[i][k] * b[k][j]	// Calculate the dot product
			}
			row = append(row, entry)
		}
		c = append(c,row)	// Append the row to the result
	}
	return c, nil
}

// Multiplies a matrix by a scalar
// Returns the result of the multiplication
func MultiplyMatrixByScalar(a M, c float32) (M, error) {
	// Cannot multiply an empty matrix
	if len(a) == 0 {
		return nil, errors.New(InvalidMatrixSizeErrorString)
	}
	m := len(a)	// rows
	n := len(a[0])	// columns
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			a[i][j] *= c
		}
	}
	return a, nil
}

// TODO: Transpose matrix




// Error Strings
const InvalidMatrixSizeErrorString = "Invalid size for matrix multiplicaiton."