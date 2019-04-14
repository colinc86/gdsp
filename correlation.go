package gdsp

// XCorr performs cross-correlation on real-valued vectors u and v. The output
// vector has length len(v) - len(u) + 1.
func XCorr(u Vector, v Vector) Vector {
	xcorr := MakeVector(0.0, len(v)-len(u)+1)
	for i := 0; i < len(xcorr); i++ {
		sum := 0.0

		for j := 0; j < len(u); j++ {
			sum += u[j] * v[i+j]
		}

		xcorr[i] = sum
	}
	return xcorr
}
