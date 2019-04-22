package gdsp

import (
	"math"
)

// MaxI returns the maximum integer in a set of integers.
func MaxI(n ...int) int {
	max := math.MinInt64
	for _, v := range n {
		if v > max {
			max = v
		}
	}
	return max
}

// ComplexRI creates a complex number from a real integer.
func ComplexRI(r int) complex128 {
	return complex(float64(r), 0.0)
}

// IsClose returns whether or not the distance between values a and b is less
// than or equal to the given tolerance.
func IsClose(a float64, b float64, tolerance float64) bool {
	return math.Abs(a-b) < tolerance
}

// IsCloseC returns whether or not the distance between values a and b is less
// than or equal to the given tolerance.
func IsCloseC(a complex128, b complex128, tolerance float64) bool {
	return math.Abs(real(a)-real(b)) < tolerance && math.Abs(imag(a)-imag(b)) < tolerance
}
