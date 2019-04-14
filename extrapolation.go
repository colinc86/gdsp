package dsp

// Extrapolate extrapolates the given signal.
func Extrapolate(input Vector, count int) Vector {
	if !input.IsZero() {
		aR, _ := Arburg(input, len(input)-1)
		bR := MakeVector(1.0, 1).PaddedTrailing(0.0, len(aR)-1)
		yR := input.SubVector(1, len(input)).Reversed()
		zR := Filtic(bR, aR, yR, nil)
		yeR, _ := Filter(bR, aR, MakeVector(0.0, count), zR)
		return yeR
	}

	return MakeVector(0.0, count)
}

// ExtrapolateC extrapolates a given signal.
func ExtrapolateC(input VectorComplex, count int) VectorComplex {
	a, _ := ArburgC(input, len(input)-1)
	b := MakeVectorComplex(ComplexRI(1), 1).PaddedTrailing(0.0, len(a)-1)
	y := input.SubVector(1, len(input)).Reversed()
	z := FilticC(b, a, y, nil)
	ye, _ := FilterC(b, a, MakeVectorComplex(0.0, count+10), z)
	return ye
}
