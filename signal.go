package dsp

// DetectLag detects positive lag between signalA (the filter) and signalB.
func DetectLag(signalA Vector, signalB Vector) int {
	dcorr := DCorr(signalA[:len(signalA)/2], signalB)

	if len(dcorr) > 1 {
		for i := 1; i < len(dcorr)-1; i++ {
			if dcorr[i-1] < dcorr[i] && dcorr[i] > dcorr[i+1] {
				return i
			}
		}
	}

	return 0
}
