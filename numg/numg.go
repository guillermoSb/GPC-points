package numg

import (
	"errors"
	"math"
)

// Matrix Definition
type M [][]float32
type V2 [2]float32
type V3 [3]float32



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

// Create an identity matrix of n size
func Identity(n int) (M, error) {
	if n <= 0 {
		return nil, errors.New(InvalidMatrixSizeForCreation)
	}
	r := M{}
	for i := 0; i < n; i++ {
		row := []float32{}
		for j := 0; j < n; j++ {
			if i == j {
				row = append(row, 1)
			} else {
				row = append(row,0)
			}
		}
		r = append(r,row)
	}

	return r, nil
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


func Subtract(A,B V2) V2{
	newV2 := V2{0,0}
	newV2[0] = B[0] - A[0]
	newV2[1] = B[1] - A[1]
	return newV2
}

func Cross(A,B V2) V3 {
	return V3{0,0,(A[0]*B[1])-(A[1]*B[0]) }
}

func Norm(A V3) float64 {
	return math.Sqrt(math.Pow(float64(A[0]),2.0) + math.Pow(float64(A[1]),2.0) + math.Pow(float64(A[2]),2.0))
}

// TODO: Transpose matrix




// Error Strings
const InvalidMatrixSizeErrorString = "Invalid size for matrix multiplicaiton."
const InvalidMatrixSizeForCreation = "Invalid size for matrix creation."