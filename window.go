package gdsp

import (
	"math"
	"math/cmplx"
	"runtime"

	"github.com/colinc86/parallel"
)

// Hann performs Hann windowing on the input vector.
func Hann(input VectorComplex) VectorComplex {
	vh := input.Copy()
	theta := 2.0 * math.Pi / float64(len(vh)-1)

	p := parallel.NewFixedProcess(runtime.NumCPU())
	p.Execute(len(vh), func(i int) {
		x := complex(theta*float64(i), 0)
		vh[i] *= (1.0 - cmplx.Cos(x)) / 2.0
	})
	return vh
}

// InverseHann performs inverse Hann windowing on the input vector.
func InverseHann(input VectorComplex) VectorComplex {
	vih := input.Copy()
	theta := 2.0 * math.Pi / float64(len(vih)-1)

	p := parallel.NewFixedProcess(runtime.NumCPU())
	p.Execute(len(vih), func(i int) {
		x := complex(theta*float64(i), 0)
		vih[i] /= (1.0 - cmplx.Cos(x)) / 2.0
	})
	return vih
}

// Hamming performs Hamming windowing on the input vector.
func Hamming(input VectorComplex) VectorComplex {
	vh := input.Copy()
	theta := 2.0 * math.Pi / float64(len(vh)-1)
	a := complex(25.0/46.0, 0.0)

	p := parallel.NewFixedProcess(runtime.NumCPU())
	p.Execute(len(vh), func(i int) {
		x := complex(theta*float64(i), 0)
		vh[i] *= a - (1.0-a)*cmplx.Cos(x)
	})
	return vh
}

// InverseHamming performs inverse Hamming windowing on the input vector.
func InverseHamming(input VectorComplex) VectorComplex {
	vih := input.Copy()
	theta := 2.0 * math.Pi / float64(len(vih)-1)
	a := complex(25.0/46.0, 0.0)

	p := parallel.NewFixedProcess(runtime.NumCPU())
	p.Execute(len(vih), func(i int) {
		x := complex(theta*float64(i), 0)
		vih[i] /= a - (1.0-a)*cmplx.Cos(x)
	})
	return vih
}

// Nuttal performs Nuttal windowing on the input vector.
func Nuttal(input VectorComplex) VectorComplex {
	vh := input.Copy()
	theta := 2.0 * math.Pi / float64(len(vh)-1)
	a0 := complex(0.355768, 0.0)
	a1 := complex(0.487396, 0.0)
	a2 := complex(0.144232, 0.0)
	a3 := complex(0.012604, 0.0)

	p := parallel.NewFixedProcess(runtime.NumCPU())
	p.Execute(len(vh), func(i int) {
		x := complex(theta*float64(i), 0)
		vh[i] *= a0 - a1*cmplx.Cos(x) + a2*cmplx.Cos(2.0*x) + a3*cmplx.Cos(3.0*x)
	})
	return vh
}

// InverseNuttal performs inverse Nuttal windowing on the input vector.
func InverseNuttal(input VectorComplex) VectorComplex {
	vih := input.Copy()
	theta := 2.0 * math.Pi / float64(len(vih)-1)
	a0 := complex(0.355768, 0.0)
	a1 := complex(0.487396, 0.0)
	a2 := complex(0.144232, 0.0)
	a3 := complex(0.012604, 0.0)

	p := parallel.NewFixedProcess(runtime.NumCPU())
	p.Execute(len(vih), func(i int) {
		x := complex(theta*float64(i), 0)
		vih[i] /= a0 - a1*cmplx.Cos(x) + a2*cmplx.Cos(2.0*x) + a3*cmplx.Cos(3.0*x)
	})
	return vih
}
