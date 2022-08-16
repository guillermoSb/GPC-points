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

func V3DotProduct(A V3, B V3) float32 {
	return A[0] * B[0] + A[1] * B[1] + A[2] * B[2]
}

func V3MultiplyScalar(A V3, c float32) V3 {
	return V3{A[0]*c, A[1]*c, A[2]*c}
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


func InverseOfMatrix(a M) (M, error) {
	// Using Gauss-Jordan Elimination

	// A
	augmentedMatrix := M{}
	n := len(a)
	for i := 0; i < n; i++ {
		row := []float32{}
		for j := 0; j < 2*n; j++ {
			if j >= n {
				if j == ((n) + i) {
					row = append(row, 1)	
				} else {
					row = append(row,0)
				}
			} else {
				row = append(row, a[i][j])
			}
		}
		augmentedMatrix = append(augmentedMatrix, row)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i != j {
				ratio := augmentedMatrix[j][i] / augmentedMatrix[i][i]
				for k := 0; k < 2*n; k++ {
					augmentedMatrix[j][k] = augmentedMatrix[j][k] - ratio * augmentedMatrix[i][k]
				}
			}
		}
	}
	for i := 0; i < n; i++ {
		divisor := augmentedMatrix[i][i]
		for j := 0; j < n * 2; j++ {
			augmentedMatrix[i][j] = augmentedMatrix[i][j] / divisor
		}
	}	
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			augmentedMatrix[i][j] = float32(math.Abs(float64(augmentedMatrix[i][j])))
		}
	}
	result := M{}
	for i := 0; i < n; i++ {
		row := []float32{}
		for j := n; j < 2*n; j++ {
			row = append(row, float32(augmentedMatrix[i][j]))
		}
		result = append(result, row)
	}

	
	return result,nil
}

func Subtract(A,B V3) V3{
	newV := V3{0,0,0}
	newV[0] = A[0] - B[0]
	newV[1] = A[1] - B[1]
	newV[2] = A[2] - B[2]
	return newV
}

func NormalizeV3(A V3) V3 {
	m := float32(math.Sqrt(math.Pow(float64(A[0]), 2) + math.Pow(float64(A[1]), 2) + math.Pow(float64(A[2]), 2)))
	newA := V3{A[0]/m,A[1]/m,A[2]/m}
	return newA
}

func Cross(A,B V3) V3 {
	det1 := (A[1] * B[2]) - (B[1]*A[2])
	det2 := (A[0] * B[2]) - (B[0]*A[2])
	det3 := (A[0] * B[1]) - (B[0]*A[1])
	return V3{det1, -det2, det3}
}

func Norm(A V3) float64 {
	return math.Sqrt(math.Pow(float64(A[0]),2.0) + math.Pow(float64(A[1]),2.0) + math.Pow(float64(A[2]),2.0))
}

// TODO: Transpose matrix




// Error Strings
const InvalidMatrixSizeErrorString = "Invalid size for matrix multiplicaiton."
const InvalidMatrixSizeForCreation = "Invalid size for matrix creation."