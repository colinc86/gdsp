package gdsp

// Extrapolate extrapolates the given real-valued signal by n samples using an autoregressive
// model.
func Extrapolate(input Vector, n int) Vector {
	if !input.IsZero() {
		aR, _ := Arburg(input, len(input)-1)
		bR := MakeVector(1.0, 1).PaddedTrailing(0.0, len(aR)-1)
		yR := input.SubVector(1, len(input)).Reversed()
		zR := Filtic(bR, aR, yR, nil)
		yeR, _ := Filter(bR, aR, MakeVector(0.0, n), zR)
		return yeR
	}

	return MakeVector(0.0, n)
}

// ExtrapolateC extrapolates the given complex-valued signal by n samples using
// an autoregressive model.
func ExtrapolateC(input VectorComplex, count int) VectorComplex {
	a, _ := ArburgC(input, len(input)-1)
	b := MakeVectorComplex(ComplexRI(1), 1).PaddedTrailing(0.0, len(a)-1)
	y := input.SubVector(1, len(input)).Reversed()
	z := FilticC(b, a, y, nil)
	ye, _ := FilterC(b, a, MakeVectorComplex(0.0, count+10), z)
	return ye
}

// FrequencyExtrapolate extrapolates a signal by extrapolating its frequency component signals.
func FrequencyExtrapolate(input Vector, count int, windowLength int, windowType WindowType) (Vector, Vector) {
	s := Spectrogram(input, windowLength, windowType)
	sc := s.FlipOrderComplex()
	var esc MatrixComplex
	for _, s := range sc {
		esc = append(esc, append(s, ExtrapolateC(s, count+windowLength/2)...))
	}
	es := esc.FlipOrderComplex()
	extrapolatedInput := InverseSpectrogram(es, windowType)
	return extrapolatedInput[len(extrapolatedInput)-count-windowLength/2 : len(extrapolatedInput)-count], extrapolatedInput[len(extrapolatedInput)-count:]
}
