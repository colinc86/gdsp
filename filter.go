package dsp

// Filter performs a 1-dimensional digital filter.
func Filter(b Vector, a Vector, x Vector, z Vector) (Vector, Vector) {
	n := len(a)

	zOut := z.Copy()
	if len(zOut) < n {
		zOut = zOut.PaddedTrailing(0.0, n-len(zOut))
	}

	y := MakeVector(0.0, len(x))
	bn := VSDiv(b, a[0])
	an := VSDiv(a, a[0])

	for m := 0; m < len(y); m++ {
		y[m] = bn[0]*x[m] + zOut[0]

		for i := 1; i < n; i++ {
			zOut[i-1] = bn[i]*x[m] + zOut[i] - an[i]*y[m]
		}
	}

	return y, zOut[:len(zOut)-1]
}

// FilterC performs a 1-dimensional digital filter.
func FilterC(b VectorComplex, a VectorComplex, x VectorComplex, z VectorComplex) (VectorComplex, VectorComplex) {
	n := len(a)

	zOut := z.Copy()
	if len(zOut) < n {
		zOut = zOut.PaddedTrailing(0.0, n-len(zOut))
	}

	y := MakeVectorComplex(0.0, len(x))
	bn := VSDivC(b, a[0])
	an := VSDivC(a, a[0])

	for m := 0; m < len(y); m++ {
		y[m] = bn[0]*x[m] + zOut[0]

		for i := 1; i < n; i++ {
			zOut[i-1] = bn[i]*x[m] + zOut[i] - an[i]*y[m]
		}
	}

	return y, zOut[:len(zOut)-1]
}
