package gdsp

import (
	"math"
	"runtime"

	"github.com/colinc86/parallel"
)

// GaussianLowpass performs a gaussian lowpass filter on the input signal.
func GaussianLowpass(input Vector, cutoff float64) Vector {
	inputZ := input.ToComplex()
	Y := FFT(inputZ)
	cutoffN := int(float64(len(Y)/2) * cutoff)
	gauss := MakeVectorComplex(0.0, len(Y))
	sigma := float64(cutoffN)

	p := parallel.NewFixedProcess(runtime.NumCPU())
	p.Execute(cutoffN+1, func(i int) {
		gauss[i] = complex(math.Exp(-math.Pow(float64(i), 2.0)/(2.0*math.Pow(sigma, 2.0))), 0.0)
	})

	gaussRev := gauss.SubVector(1, cutoffN+1).Reversed()
	p.Execute(cutoffN, func(i int) {
		gauss[len(Y)-cutoffN+i] = gaussRev[i]
	})

	lp := IFFT(VMulEC(Y, gauss)).Real()
	return lp
}
