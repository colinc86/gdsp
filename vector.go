package gdsp

import (
	"math/cmplx"
)

// Vector types represent a real-valued vector.
type Vector []float64

// VectorComplex types represent a complex-valued vector.
type VectorComplex []complex128

// Matrix types represent an array of vectors.
type Matrix []Vector

// MatrixComplex types represent an array of complex vectors.
type MatrixComplex []VectorComplex

// MARK: Constructors

// MakeMatrix creates and returns a new matrix.
func MakeMatrix(repeating float64, rows int, columns int) Matrix {
	m := make([]Vector, 0, 0)
	for i := 0; i < rows; i++ {
		m = append(m, MakeVector(repeating, columns))
	}
	return m
}

// MakeMatrixComplex creates and returns a new matrix.
func MakeMatrixComplex(repeating complex128, rows int, columns int) MatrixComplex {
	m := make([]VectorComplex, 0, 0)
	for i := 0; i < rows; i++ {
		m = append(m, MakeVectorComplex(repeating, columns))
	}
	return m
}

// MakeVector creates a new vector.
func MakeVector(repeating float64, count int) Vector {
	v := make(Vector, count, count)
	for i := range v {
		v[i] = repeating
	}
	return v
}

// MakeVectorFromArray creates and returns a new vector from an array.
func MakeVectorFromArray(input []float64) Vector {
	return append([]float64(nil), input...)
}

// MakeVectorComplex creates a new vector.
func MakeVectorComplex(repeating complex128, count int) VectorComplex {
	v := make(VectorComplex, count, count)
	for i := range v {
		v[i] = repeating
	}
	return v
}

// MakeVectorComplexFromArray creates and returns a new vector from an array.
func MakeVectorComplexFromArray(input []complex128) VectorComplex {
	return append([]complex128(nil), input...)
}

// MakeVectorComplexFromSplit creates a new vector from real and imaginary parts.
func MakeVectorComplexFromSplit(real Vector, imag Vector) VectorComplex {
	v := MakeVectorComplex(0.0, len(real))
	for i := 0; i < len(real); i++ {
		v[i] = complex(real[i], imag[i])
	}
	return v
}

// MARK: Conversion methods

// ToComplex converts a real-valued vector to a complex-valued vector.
func (v Vector) ToComplex() VectorComplex {
	vc := MakeVectorComplex(0.0, len(v))
	for i, r := range v {
		vc[i] = complex(r, 0.0)
	}
	return vc
}

// Real returns the real componenets of the vector.
func (v VectorComplex) Real() Vector {
	rp := MakeVector(0.0, len(v))
	for i, c := range v {
		rp[i] = real(c)
	}
	return rp
}

// Imag returns the imaginary components of the vector.
func (v VectorComplex) Imag() Vector {
	rp := MakeVector(0.0, len(v))
	for i, c := range v {
		rp[i] = imag(c)
	}
	return rp
}

// Conj returns conjugate vector.
func (v VectorComplex) Conj() VectorComplex {
	vc := v.Copy()
	for i, c := range vc {
		vc[i] = cmplx.Conj(c)
	}
	return vc
}

// MARK: Vector-Vector Functions

// VMulE performs vector element to vector element multiplication and returns
// the result.
//
// v[i] *= u[i]
func VMulE(u Vector, v Vector) Vector {
	vc := v.Copy()
	for i := 0; i < len(vc); i++ {
		vc[i] *= u[i]
	}
	return vc
}

// VMulEC performs vector element to vector element multiplication and returns
// the result.
//
// v[i] *= u[i]
func VMulEC(u VectorComplex, v VectorComplex) VectorComplex {
	vc := v.Copy()
	for i := 0; i < len(vc); i++ {
		vc[i] *= u[i]
	}
	return vc
}

// VMulESum performs vector element to vector element multiplication, adds the
// products and returns the resulting sum.
//
// s += u[i] * v[i]
func VMulESum(u Vector, v Vector) float64 {
	s := 0.0
	for i := 0; i < len(u); i++ {
		s += u[i] * v[i]
	}
	return s
}

// VMulESumC performs vector element to vector element multiplication, adds the
// products and returns the resulting sum.
//
// s += u[i] * v[i]
func VMulESumC(u VectorComplex, v VectorComplex) complex128 {
	s := complex(0.0, 0.0)
	for i := 0; i < len(u); i++ {
		s += u[i] * v[i]
	}
	return s
}

// VAdd performs vector addition and returns the result.
//
// v[i] += u[i]
func VAdd(u Vector, v Vector) Vector {
	vc := v.Copy()
	for i := 0; i < len(vc); i++ {
		vc[i] += u[i]
	}
	return vc
}

// VAddC performs vector addition and returns the result.
//
// v[i] += u[i]
func VAddC(u VectorComplex, v VectorComplex) VectorComplex {
	vc := v.Copy()
	for i := 0; i < len(vc); i++ {
		vc[i] += u[i]
	}
	return vc
}

// VSub performs vector subtraction and returns the result.
//
// v[i] -= u[i]
func VSub(u Vector, v Vector) Vector {
	vc := u.Copy()
	for i := 0; i < len(vc); i++ {
		vc[i] -= v[i]
	}
	return vc
}

// VSubC performs vector subtraction and returns the result.
//
// v[i] -= u[i]
func VSubC(u VectorComplex, v VectorComplex) VectorComplex {
	vc := u.Copy()
	for i := 0; i < len(vc); i++ {
		vc[i] -= v[i]
	}
	return vc
}

// MARK: Vector-Scaler Functions

// VSMul performs vector-scaler multiplication and returns the result.
//
// v[i] *= s
func VSMul(v Vector, s float64) Vector {
	vc := v.Copy()
	for i := 0; i < len(vc); i++ {
		vc[i] *= s
	}
	return vc
}

// VSMulC performs vector-scaler multiplication and returns the result.
//
// v[i] *= s
func VSMulC(v VectorComplex, s complex128) VectorComplex {
	vc := v.Copy()
	for i := 0; i < len(vc); i++ {
		vc[i] *= s
	}
	return vc
}

// VSDiv performs vector-scaler division and returns the result.
//
// v[i] /= s
func VSDiv(v Vector, s float64) Vector {
	vc := v.Copy()
	for i := 0; i < len(vc); i++ {
		vc[i] /= s
	}
	return vc
}

// VSDivC performs vector-scaler division and returns the result.
//
// v[i] /= s
func VSDivC(v VectorComplex, s complex128) VectorComplex {
	vc := v.Copy()
	for i := 0; i < len(vc); i++ {
		vc[i] /= s
	}
	return vc
}

// MARK: Single Vector Functions

// VESum adds together the elements of v and returns the result.
func VESum(v Vector) float64 {
	s := 0.0
	for _, r := range v {
		s += r
	}
	return s
}

// VESumC adds together the elements of v and returns the result.
func VESumC(v VectorComplex) complex128 {
	s := complex(0.0, 0.0)
	for _, c := range v {
		s += c
	}
	return s
}

// VSumSq adds together the squared elements of v and returns the result.
func VSumSq(v Vector) float64 {
	s := 0.0
	for _, r := range v {
		s += r * r
	}
	return s
}

// VSumSqC adds together the squared elements of u and returns the result.
func VSumSqC(v VectorComplex) complex128 {
	s := complex(0.0, 0.0)
	for _, c := range v {
		s += c * c
	}
	return s
}

// VNeg returns the additive inverse of v.
//
// VSMul(v, -1.0)
func VNeg(v Vector) Vector {
	return VSMul(v, -1.0)
}

// VNegC returns the additive inverse of v.
//
// VSMulC(v, -1.0)
func VNegC(v VectorComplex) VectorComplex {
	return VSMulC(v, -1.0)
}

// MARK: Vector methods

// Copy creates and returns a new vector initializes with the elements of v.
func (v Vector) Copy() Vector {
	return append([]float64(nil), v...)
}

// Copy creates and returns a new vector initializes with the elements of v.
func (v VectorComplex) Copy() VectorComplex {
	return append([]complex128(nil), v...)
}

// IsZero checks the elements of u and returns false iff any real or imaginary
// component is non-zero.
func (v Vector) IsZero() bool {
	for _, r := range v {
		if r != 0.0 {
			return false
		}
	}
	return true
}

// IsZero checks the elements of u and returns false iff any real or imaginary
// component is non-zero.
func (v VectorComplex) IsZero() bool {
	for _, c := range v {
		if real(c) != 0.0 || imag(c) != 0.0 {
			return false
		}
	}
	return true
}

// Reversed reverses the vector.
func (v Vector) Reversed() Vector {
	rv := MakeVector(0.0, len(v))
	j := 0
	for i := len(v) - 1; i >= 0; i-- {
		rv[j] = v[i]
		j++
	}
	return rv
}

// Reverse reverses the vector by copying and returning the elements from the
// provided start to stop indexes.
func (v Vector) Reverse(start int, stop int) Vector {
	rv := MakeVector(0.0, 0)
	for i := start; i >= stop; i-- {
		rv = append(rv, v[i])
	}
	return rv
}

// Reversed reverses the vector.
func (v VectorComplex) Reversed() VectorComplex {
	rv := MakeVectorComplex(0.0, len(v))
	j := 0
	for i := len(v) - 1; i >= 0; i-- {
		rv[j] = v[i]
		j++
	}
	return rv
}

// Reverse reverses the vector by copying and returning the elements from the
// provided start to stop indexes.
func (v VectorComplex) Reverse(start int, stop int) VectorComplex {
	rv := MakeVectorComplex(0.0, 0)
	for i := start; i >= stop; i-- {
		rv = append(rv, v[i])
	}
	return rv
}

// SubVector returns a subvector.
func (v Vector) SubVector(start int, stop int) Vector {
	rv := MakeVector(0.0, stop-start)
	j := 0
	for i := start; i < stop; i++ {
		rv[j] = v[i]
		j++
	}
	return rv
}

// SubVector returns a subvector.
func (v VectorComplex) SubVector(start int, stop int) VectorComplex {
	rv := MakeVectorComplex(0.0, stop-start)
	j := 0
	for i := start; i < stop; i++ {
		rv[j] = v[i]
		j++
	}
	return rv
}

// Padded pads a vector with leading and trailing values.
func (v Vector) Padded(leadingValue float64, leadingCount int, trailingValue float64, trailingCount int) Vector {
	l := MakeVector(leadingValue, leadingCount)
	t := MakeVector(trailingValue, trailingCount)
	u := append(l, v...)
	return append(u, t...)
}

// Padded pads a vector with leading and trailing values.
func (v VectorComplex) Padded(leadingValue complex128, leadingCount int, trailingValue complex128, trailingCount int) VectorComplex {
	l := MakeVectorComplex(leadingValue, leadingCount)
	t := MakeVectorComplex(trailingValue, trailingCount)
	u := append(l, v...)
	return append(u, t...)
}

// PaddedLeading pads a vector with leading values.
func (v Vector) PaddedLeading(leadingValue float64, leadingCount int) Vector {
	return v.Padded(leadingValue, leadingCount, 0.0, 0)
}

// PaddedLeading pads a vector with leading values.
func (v VectorComplex) PaddedLeading(leadingValue complex128, leadingCount int) VectorComplex {
	return v.Padded(leadingValue, leadingCount, 0.0, 0)
}

// PaddedTrailing pads a vector with trailing values.
func (v Vector) PaddedTrailing(trailingValue float64, trailingCount int) Vector {
	return v.Padded(0.0, 0, trailingValue, trailingCount)
}

// PaddedTrailing pads a vector with trailing values.
func (v VectorComplex) PaddedTrailing(trailingValue complex128, trailingCount int) VectorComplex {
	return v.Padded(0.0, 0, trailingValue, trailingCount)
}
