package gdsp

// Matrix types represent an array of vectors.
type Matrix []Vector

// MatrixComplex types represent an array of complex vectors.
type MatrixComplex []VectorComplex

// MARK: Constructors

// MakeMatrix creates and returns a new matrix in row major order.
func MakeMatrix(repeating float64, rows int, columns int) Matrix {
	m := make([]Vector, 0, 0)
	for i := 0; i < rows; i++ {
		m = append(m, MakeVector(repeating, columns))
	}
	return m
}

// MakeMatrixComplex creates and returns a new matrix in row major order.
func MakeMatrixComplex(repeating complex128, rows int, columns int) MatrixComplex {
	m := make([]VectorComplex, 0, 0)
	for i := 0; i < rows; i++ {
		m = append(m, MakeVectorComplex(repeating, columns))
	}
	return m
}

// MARK: Conversion methods

// FlipOrder reverses the row and column order when indexing the matrix.
// The NewMatrix... constructors assume row major order and it is up to the
// programmer to keep track of what order their instance of Matrix has.
func (m Matrix) FlipOrder() Matrix {
	if len(m) == 0 || len(m[0]) == 0 {
		return nil
	}

	var flipped []Vector
	for i := 0; i < len(m[0]); i++ {
		var vec Vector
		for j := 0; j < len(m); j++ {
			vec = append(vec, m[j][i])
		}
		flipped = append(flipped, vec)
	}
	return flipped
}

// FlipOrderComplex reverses the row and column order when indexing the matrix.
// The NewMatrix... constructors assume row major order and it is up to the
// programmer to keep track of what order their instance of Matrix has.
func (m MatrixComplex) FlipOrderComplex() MatrixComplex {
	if len(m) == 0 || len(m[0]) == 0 {
		return nil
	}

	var flipped []VectorComplex
	for i := 0; i < len(m[0]); i++ {
		var vec VectorComplex
		for j := 0; j < len(m); j++ {
			vec = append(vec, m[j][i])
		}
		flipped = append(flipped, vec)
	}
	return flipped
}
