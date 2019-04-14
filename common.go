package dsp

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
