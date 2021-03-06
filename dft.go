package gdsp

import (
	"math"
	"math/cmplx"
)

// DFT performs a discrete Fourier transform on the complex-valued input vector
// and returns the result. Pass true for the forward parameter to perform a forward
// transform and false for an inverse transform.
func DFT(input VectorComplex, forward bool) VectorComplex {
	N := len(input)
	theta := 2.0 * math.Pi / float64(N)
	coeff := 1.0

	if forward {
		coeff = -1.0
	}

	output := MakeVectorComplex(0.0, N)
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			x := float64(i * j)
			realP := math.Cos(theta * x)
			imagP := coeff * math.Sin(theta*x)

			output[i] += input[j] * complex(realP, imagP)
		}

		if !forward {
			output[i] /= complex(float64(N), 0.0)
		}
	}

	return output
}

// FFT performs a discrete Fourier transform on the complex-valued input vector
// using the Cooley-Turkey FFT algorithm. For an inverse FFT, see the IFFT function.
func FFT(input VectorComplex) VectorComplex {
	if len(input) == 1 {
		return input
	}

	if len(input)%2 != 0 {
		return DFT(input, true)
	}

	var evenInput VectorComplex
	var oddInput VectorComplex
	for i := 0; i < len(input); i += 2 {
		evenInput = append(evenInput, input[i])
		oddInput = append(oddInput, input[i+1])
	}

	evenDFT := FFT(evenInput)
	oddDFT := FFT(oddInput)

	for k := 0; k < len(input)/2; k++ {
		x := float64(k) / float64(len(input))
		t := evenDFT[k]
		ec := cmplx.Exp(complex(0.0, -2.0*math.Pi*x))
		evenDFT[k] = t + ec*oddDFT[k]
		oddDFT[k] = t - ec*oddDFT[k]
	}

	return append(evenDFT, oddDFT...)
}

// IFFT performs an inverse discrete Fourier transform on the complex-valued input
// vector using the Cooley-Turkey FFT algorithm. For a forward FFT, see the FFT
// function.
func IFFT(input VectorComplex) VectorComplex {
	inputConjugate := input.Conj()
	fft := FFT(inputConjugate)
	return VSDivC(fft.Conj(), complex(float64(len(input)), 0.0))
}
