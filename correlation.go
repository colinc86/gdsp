package gdsp

// ACorr performs autocorrelation on real-valued vector u. The output vector has
// length 2 * len(u) - 1.
func ACorr(u Vector) Vector {
	return ACorrC(u.ToComplex()).Real()
}

// ACorrC performs autocorrelation on complex-valued vector u. The output vector
// has length 2 * len(u) - 1.
func ACorrC(u VectorComplex) VectorComplex {
	zu := u.PaddedTrailing(0.0, len(u))
	zufft := FFT(zu)
	zfft := VMulEC(zufft, zufft.Conj())
	return IFFT(zfft).SubVector(0, 2*len(u)-1)
}

// XCorr performs cross-correlation on real-valued vectors u and v. The output
// vector has length len(v) - len(u) + 1.
func XCorr(u Vector, v Vector) Vector {
	return XCorrC(u.ToComplex(), v.ToComplex()).Real()
}

// XCorrC performs cross-correlation on complex-valued vectors u and v. The output
// vector has length len(v) - len(u) + 1.
func XCorrC(u VectorComplex, v VectorComplex) VectorComplex {
	mLen := 2.0 * MaxI(len(u), len(v))

	uDiff := mLen - len(u)
	zu := u.PaddedTrailing(0.0, uDiff)
	zufft := FFT(zu)

	vDiff := mLen - len(v)
	zv := v.PaddedTrailing(0.0, vDiff)
	zvfft := FFT(zv)

	uvfft := VMulEC(zufft, zvfft.Conj())
	return IFFT(uvfft).SubVector(0, len(u)+len(v)-1)
}
