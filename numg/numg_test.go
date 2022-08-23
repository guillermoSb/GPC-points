package numg

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Matrix multiplication tests
func Test_multiply_two_2x2_matrices(t *testing.T) {
	
	// Arrange
	a := M{{1,2},
		   {2,2}}
	b := M{{3,3},
		   {3,3}}
	result := M{{9,9}, {12,12}}
	// Act
	c, _ := MultiplyMatrices(a,b)
	// Assert
	assert.Equal(t, result, c)
}


func Test_multiply_matrices_with_invalid_sices_returns_error(t *testing.T) {
	// Arrange
	a := M{{1,2}}	// 1 x 2
	b := M {{3,3}, {2,2}, {3,3}}	// 3 x 2
	// Act
	_, err := MultiplyMatrices(a,b)
	// Assert
	assert.EqualError(t, err, InvalidMatrixSizeErrorString)
}

func Test_cant_multiply_empty_matrices(t *testing.T) {
	// Arrange
	a := M{}	// 0x0
	b := M{{4,4}}	// 1x2
	// Act
	_, err := MultiplyMatrices(a,b)
	// Assert
	assert.EqualError(t, err, InvalidMatrixSizeErrorString)
}


// Matrix multiplication by a scalar

func Test_multiply_matrix_by_scalar(t *testing.T) {
	// Arrange
	a := M{{1,1}}
	c := float32(2)
	// Act
	r, _ := MultiplyMatrixByScalar(a,c)
	// Assert
	assert.Equal(t, M{{2,2}},r)
}

func Test_cant_multily_empty_matrix_by_scalar(t *testing.T) {
	// Arrange
	a := M{}
	// Act
	_, err := MultiplyMatrixByScalar(a, 2)
	// Assert
	assert.Error(t, err, InvalidMatrixSizeErrorString)
}


// Identity Matrix
func Test_generates_identity_matrix(t *testing.T) {
	// Act
	m, _ := Identity(4)	
	// Assert
	for i := 0; i < len(m); i++ {
		for j := 0; j < len(m); j++ {
			if i == j {
				assert.Equal(t, float32(1), m[i][j])
			}
		}
	}
}

func Test_calculates_cross_product(t *testing.T) {
	// Arrange
	A := V3{13,8,0}
	B := V3{-1,2,0}
	// Act
	C := Cross(A,B)
	// Assert
	assert.Equal(t, float32(34.0), C[2])
}

func Test_calcualtes_vector_norm(t *testing.T) {
	// Arrange
	A := V3{2,2,1}
	// Act
	r := Norm(A)
	// Assert
	assert.Equal(t, 3.0, r)
}

func Test_calculateInverseMatrix(t *testing.T) {
	// Arrange
	m := M{{1,1,1}, {2,1,1}, {1, 1, 2}}
	// Act
	r, _ := InverseOfMatrix(m)
	// Assert
	
	expected := M{
		{-1,1,0},
		{3,-1,-1},
		{-1,0,1},
	}

	assert.Equal(t, expected, r)
}

func Test_shuffle(t *testing.T){
	
	fmt.Println(MakePermutation())
}
