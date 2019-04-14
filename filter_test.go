package gdsp

import "testing"

func TestFilterLowPass1(t *testing.T) {
	b := MakeVectorFromArray([]float64{0.25, 0.25})
	a := MakeVectorFromArray([]float64{1.0, -0.5})
	zfRef := 1.25
	eRef := 0.875
	zf, e := Filter(b, a, []float64{1.0}, []float64{1.0})
	if zf[0] != zfRef {
		t.FailNow()
	}
	if e[0] != eRef {
		t.FailNow()
	}
}

func TestFilterLowPass2(t *testing.T) {
	b := MakeVectorFromArray([]float64{0.25, 0.25})
	a := MakeVectorFromArray([]float64{1.0, -0.5})
	zfRef := 0.25
	eRef := 0.375
	zf, e := Filter(b, a, []float64{1.0}, nil)
	if zf[0] != zfRef {
		t.FailNow()
	}
	if e[0] != eRef {
		t.FailNow()
	}
}

func TestFilterHighPass1(t *testing.T) {
	b := MakeVectorFromArray([]float64{0.25, -0.25})
	a := MakeVectorFromArray([]float64{1.0, 0.5})
	zfRef := 1.0
	eRef := -0.5
	zf, e := Filter(b, a, []float64{0.0}, []float64{1.0})
	if zf[0] != zfRef {
		t.FailNow()
	}
	if e[0] != eRef {
		t.FailNow()
	}
}

func TestFilterHighPass2(t *testing.T) {
	b := MakeVectorFromArray([]float64{0.25, -0.25})
	a := MakeVectorFromArray([]float64{1.0, 0.5})
	zfRef := 0.0
	eRef := 0.0
	zf, e := Filter(b, a, []float64{0.0}, nil)
	if zf[0] != zfRef {
		t.FailNow()
	}
	if e[0] != eRef {
		t.FailNow()
	}
}
