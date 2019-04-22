package gdsp

import "testing"

func TestConv1(t *testing.T) {
	u := MakeVectorFromArray([]float64{1.0, 2.0, 3.0})
	v := MakeVectorFromArray([]float64{1.0})
	conv := Conv(u, v)

	if !conv.IsCloseToVector(u, 0.00001) {
		t.Fail()
	}
}

func TestConv2(t *testing.T) {
	u := MakeVectorFromArray([]float64{1.0, 2.0, 3.0})
	v := MakeVectorFromArray([]float64{1.0, 2.0})
	ex := MakeVectorFromArray([]float64{1.0, 4.0, 7.0, 6.0})
	conv := Conv(u, v)

	if !conv.IsCloseToVector(ex, 0.00001) {
		t.Fail()
	}
}

func TestConv3(t *testing.T) {
	u := MakeVectorFromArray([]float64{1.0, 2.0, 3.0})
	v := MakeVectorFromArray([]float64{1.0, 2.0, 3.0})
	ex := MakeVectorFromArray([]float64{1.0, 4.0, 10.0, 12.0, 9.0})
	conv := Conv(u, v)

	if !conv.IsCloseToVector(ex, 0.00001) {
		t.Fail()
	}
}

func TestConv4(t *testing.T) {
	u := MakeVectorFromArray([]float64{1.0, 2.0})
	v := MakeVectorFromArray([]float64{1.0, 2.0, 3.0})
	ex := MakeVectorFromArray([]float64{1.0, 4.0, 7.0, 6.0})
	conv := Conv(u, v)

	if !conv.IsCloseToVector(ex, 0.00001) {
		t.Fail()
	}
}

func TestConv5(t *testing.T) {
	u := MakeVectorFromArray([]float64{1.0})
	v := MakeVectorFromArray([]float64{1.0, 2.0, 3.0})
	conv := Conv(u, v)

	if !conv.IsCloseToVector(v, 0.00001) {
		t.Fail()
	}
}
