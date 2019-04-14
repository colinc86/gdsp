package dsp

// Interpolate interpolates a signal using a DFT.
func Interpolate(input Vector, upsampleMultiple int) Vector {
	if upsampleMultiple < 2 || len(input)%2 != 0 {
		return input
	}

	fft := FFT(input.ToComplex())
	leftHalf := fft[:len(fft)/2]
	rightHalf := fft[len(fft)/2:]

	paddedFFT := MakeVectorComplex(0.0, len(fft)*(upsampleMultiple-1))
	paddedFFT = append(leftHalf, paddedFFT...)
	paddedFFT = append(paddedFFT, rightHalf...)

	return VSMulC(IFFT(paddedFFT), ComplexRI(upsampleMultiple)).Real()
}

// InterpolateC interpolates a signal using a DFT.
func InterpolateC(input VectorComplex, upsampleMultiple int) VectorComplex {
	if upsampleMultiple < 2 || len(input)%2 != 0 {
		return input
	}

	fft := FFT(input)
	leftHalf := fft[:len(fft)/2]
	rightHalf := fft[len(fft)/2:]

	paddedFFT := MakeVectorComplex(0.0, len(fft)*(upsampleMultiple-1))
	paddedFFT = append(leftHalf, paddedFFT...)
	paddedFFT = append(paddedFFT, rightHalf...)

	return VSMulC(IFFT(paddedFFT), ComplexRI(upsampleMultiple))
}
