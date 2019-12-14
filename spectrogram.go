package gdsp

// Spectrogram calculates a spectrogram from the given input signal using the specified
// window length and type.
func Spectrogram(input Vector, windowLength int, windowType WindowType) MatrixComplex {
	var output MatrixComplex
	for i := 0; i < len(input)-windowLength; i++ {
		window := Window(windowType, input[i:i+windowLength].ToComplex())
		output = append(output, FFT(window))
	}
	return output
}

// InverseSpectrogram calculates a signal from a spectrogram that was generated with the given
// window type.
func InverseSpectrogram(spectrogram MatrixComplex, windowType WindowType) Vector {
	var output Vector
	for i := 0; i < len(spectrogram); i++ {
		idft := IFFT(spectrogram[i])
		window := InverseWindow(windowType, idft).Real()
		output = append(output, window[len(window)/2-1])
	}
	return output
}
