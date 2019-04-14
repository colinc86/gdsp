package gdsp

import "testing"

func TestXCorr1(t *testing.T) {
	u := []float64{1.0, 2.0, 3.0, 2.0, 1.0}
	v := []float64{1.0, 2.0, 3.0, 2.0, 1.0}
	c := XCorr(u, v)

	if len(c) != 1 {
		t.Fail()
	}
}

func TestXCorr2(t *testing.T) {
	u := []float64{2.0 / 3.0, 1.0, 2.0 / 3.0}
	v := []float64{1.0 / 3.0, 2.0 / 3.0, 1.0, 2.0 / 3.0, 1.0 / 3.0}
	c := XCorr(u, v)

	if len(c) != 3 {
		t.Fail()
	}
}
