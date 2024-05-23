package gdsp

import (
	"math"
	"sort"
)

// Min returns the minimum value from a vector.
func Min(v Vector) (float64, int) {
	if len(v) == 0 {
		return 0.0, -1
	}

	minValue := math.MaxFloat64
	minIndex := -1
	for i, r := range v {
		if r < minValue {
			minValue = r
			minIndex = i
		}
	}
	return minValue, minIndex
}

func Mins(v Vector, length int) (Vector, Vector) {
	var mins, minIndices Vector
	for i := 0; i < len(v)-length; i++ {
		minValue, minIndex := Min(v[i : i+length])
		minIndex += i

		found := false
		for _, j := range minIndices {
			if float64(minIndex) == j {
				found = true
				break
			}
		}

		if !found {
			mins = append(mins, minValue)
			minIndices = append(minIndices, float64(minIndex))
		}
	}
	return mins, minIndices
}

// Max returns the maximum value from a vector.
func Max(v Vector) (float64, int) {
	if len(v) == 0 {
		return 0.0, -1
	}

	maxValue := -math.MaxFloat64
	maxIndex := -1
	for i, r := range v {
		if r > maxValue {
			maxValue = r
			maxIndex = i
		}
	}
	return maxValue, maxIndex
}

func Maxs(v Vector, length int) (Vector, Vector) {
	var maxs, maxIndices Vector
	for i := 0; i < len(v)-length; i++ {
		maxValue, maxIndex := Max(v[i : i+length])
		maxIndex += i

		found := false
		for _, j := range maxIndices {
			if float64(maxIndex) == j {
				found = true
				break
			}
		}

		if !found {
			maxs = append(maxs, maxValue)
			maxIndices = append(maxIndices, float64(maxIndex))
		}
	}
	return maxs, maxIndices
}

// MinReal returns the minimum value from a slice of floats.
func MinReal(v VectorComplex) float64 {
	if len(v) == 0 {
		return 0.0
	}

	minValue := math.MaxFloat64
	for _, c := range v {
		if real(c) < minValue {
			minValue = real(c)
		}
	}
	return minValue
}

// MaxReal returns the maximum value from a slice of floats.
func MaxReal(v VectorComplex) float64 {
	if len(v) == 0 {
		return 0.0
	}

	maxValue := -math.MaxFloat64
	for _, c := range v {
		if real(c) > maxValue {
			maxValue = real(c)
		}
	}
	return maxValue
}

// MinImag returns the minimum value from a slice of floats.
func MinImag(v VectorComplex) float64 {
	if len(v) == 0 {
		return 0.0
	}

	minValue := math.MaxFloat64
	for _, c := range v {
		if imag(c) < minValue {
			minValue = imag(c)
		}
	}
	return minValue
}

// MaxImag returns the maximum value from a slice of floats.
func MaxImag(v VectorComplex) float64 {
	if len(v) == 0 {
		return 0.0
	}

	maxValue := -math.MaxFloat64
	for _, c := range v {
		if imag(c) > maxValue {
			maxValue = imag(c)
		}
	}
	return maxValue
}

// Mean returns the mean of the elements of v.
func Mean(v Vector) float64 {
	return VESum(v) / float64(len(v))
}

// Median returns the median of the elements of v.
func Median(v Vector) float64 {
	vc := v.Copy()
	sort.Float64s(vc)
	if len(v)%2 == 0 {
		v1 := v[len(v)/2]
		v2 := v[len(v)/2+1]
		return (v1 + v2) / 2.0
	}

	return vc[(len(vc)-1)+1]
}

// StdDev returns the standard deviation of the elements of v.
func StdDev(v Vector) float64 {
	s := 0.0
	m := Mean(v)
	for _, r := range v {
		x := r - m
		s += x * x
	}
	return math.Sqrt(s / float64(len(v)-1))
}

// Normalize normalizes the given vector.
func Normalize(v Vector) Vector {
	nv := MakeVector(0.0, len(v))
	s := StdDev(v)
	m := Mean(v)
	for i, r := range v {
		nv[i] = (r - m) / s
	}
	return nv
}

// NormalizeStrict normalizes a vector using the max and min elements.
func NormalizeStrict(v Vector) (Vector, []float64) {
	maxValue, _ := Max(v)
	minValue, _ := Min(v)
	difference := maxValue - minValue

	nv := MakeVector(0.0, len(v))
	for i, r := range v {
		x := 0.0
		if difference != 0.0 {
			x = (r - minValue) / difference
		}
		nv[i] = x
	}
	return nv, []float64{minValue, maxValue}
}

// NormalizeStrictC normalizes a complex-valued vector.
func NormalizeStrictC(v VectorComplex) (VectorComplex, []float64) {
	real := v.Real()
	imag := v.Imag()

	nReal, rLimits := NormalizeStrict(real)
	nImag, iLimits := NormalizeStrict(imag)

	vc := MakeVectorComplexFromSplit(nReal, nImag)

	return vc, append(rLimits, iLimits...)
}
