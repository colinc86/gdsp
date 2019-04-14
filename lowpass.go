package dsp

import (
	"math"
)

// GaussianLowpass performs a gaussian lowpass filter on the input signal.
func GaussianLowpass(input Vector, cutoff float64) Vector {
	inputZ := input.ToComplex()
	Y := FFT(inputZ)
	cutoffN := int(float64(len(Y)/2) * cutoff)
	gauss := MakeVectorComplex(0.0, len(Y))
	sigma := float64(cutoffN)

	for i := 0; i < cutoffN+1; i++ {
		gauss[i] = complex(math.Exp(-math.Pow(float64(i), 2.0)/(2.0*math.Pow(sigma, 2.0))), 0.0)
	}

	gaussRev := gauss.SubVector(1, cutoffN+1).Reversed()
	for i := 0; i < cutoffN; i++ {
		gauss[len(Y)-cutoffN+i] = gaussRev[i]
	}

	lp := IFFT(VMulEC(Y, gauss)).Real()
	return lp
}

// IIR performs an IIR filter on input with the given response.
func IIR(input Vector, response float64) Vector {
	output := make([]float64, len(input), len(input))
	for i, v := range input {
		if i == 0 {
			output[i] = v
		} else {
			output[i] = output[i-1]*(1.0-response) + v*response
		}
	}
	return output
}
