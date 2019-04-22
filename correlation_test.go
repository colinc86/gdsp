package gdsp

import (
	"testing"
)

func TestXCorr1(t *testing.T) {
	u := MakeVectorFromArray([]float64{1.0, 2.0, 3.0})
	v := MakeVectorFromArray([]float64{1.0})
	xCorr := XCorr(u, v)

	if !xCorr.IsCloseToVector(u, 0.00001) {
		t.Fail()
	}
}

func TestXCorr2(t *testing.T) {
	u := MakeVectorFromArray([]float64{1.0, 2.0, 3.0})
	v := MakeVectorFromArray([]float64{1.0, 2.0})
	ex := MakeVectorFromArray([]float64{5.0, 8.0, 3.0, 0.0})
	xCorr := XCorr(u, v)

	if !xCorr.IsCloseToVector(ex, 0.00001) {
		t.Fail()
	}
}

func TestXCorr3(t *testing.T) {
	u := MakeVectorFromArray([]float64{1.0, 2.0, 3.0})
	v := MakeVectorFromArray([]float64{1.0, 2.0, 3.0})
	ex := MakeVectorFromArray([]float64{14.0, 8.0, 3.0, 0.0, 3.0})
	conv := XCorr(u, v)

	if !conv.IsCloseToVector(ex, 0.00001) {
		t.Fail()
	}
}

func TestACorr1(t *testing.T) {
	v := MakeVectorFromArray([]float64{1.0})
	aCorr := ACorr(v)

	if !aCorr.IsCloseToVector(v, 0.00001) {
		t.Fail()
	}
}

func TestACorr2(t *testing.T) {
	v := MakeVectorFromArray([]float64{1.0, 2.0})
	ex := MakeVectorFromArray([]float64{5.0, 2.0, 0.0})
	aCorr := ACorr(v)

	if !aCorr.IsCloseToVector(ex, 0.00001) {
		t.Fail()
	}
}

func TestACorr3(t *testing.T) {
	v := MakeVectorFromArray([]float64{1.0, 2.0, 3.0})
	ex := MakeVectorFromArray([]float64{14.0, 8.0, 3.0, 0.0, 3.0})
	aCorr := ACorr(v)

	if !aCorr.IsCloseToVector(ex, 0.00001) {
		t.Fail()
	}
}
