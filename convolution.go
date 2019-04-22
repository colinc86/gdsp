package gdsp

// Conv performs convolution on real-valued vectors u and v. The output
// vector has length len(v) - len(u) + 1.
func Conv(u Vector, v Vector) Vector {
	return ConvC(u.ToComplex(), v.ToComplex()).Real()
}

// ConvC performs convolution on real-valued vectors u and v. The output
// vector has length len(v) - len(u) + 1.
func ConvC(u VectorComplex, v VectorComplex) VectorComplex {
	mLen := 2.0 * MaxI(len(u), len(v))

	uDiff := mLen - len(u)
	zu := u.PaddedTrailing(0.0, uDiff)
	zufft := FFT(zu)

	vDiff := mLen - len(v)
	zv := v.PaddedTrailing(0.0, vDiff)
	zvfft := FFT(zv)

	uvfft := VMulEC(zufft, zvfft)
	return IFFT(uvfft).SubVector(0, len(u)+len(v)-1)
}
