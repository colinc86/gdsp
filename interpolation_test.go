package dsp

import (
	"fmt"
	"testing"
)

func TestInterpolation(t *testing.T) {
	v := []float64{0.0, 1.0, 0.0, -1.0, 0.0, 1.0}
	i := Interpolate(v, 2)

	fmt.Printf("%v", i)
}
