package dsp

import (
	"fmt"
	"math"
	"testing"
)

func TestFilticLowPass1(t *testing.T) {
	b := MakeVectorFromArray([]float64{0.25, 0.25})
	a := MakeVectorFromArray([]float64{1.0, -0.5})
	zfRef := 0.75
	zf := Filtic(b, a, []float64{1.0}, []float64{1.0})
	if zf[0] != zfRef {
		t.FailNow()
	}
}

func TestFilticLowPass2(t *testing.T) {
	b := MakeVectorFromArray([]float64{0.25, 0.25})
	a := MakeVectorFromArray([]float64{1.0, -0.5})
	zfRef := 0.5
	zf := Filtic(b, a, []float64{1.0}, nil)
	if zf[0] != zfRef {
		t.FailNow()
	}
}

func TestFilticHighPass1(t *testing.T) {
	b := MakeVectorFromArray([]float64{0.25, -0.25})
	a := MakeVectorFromArray([]float64{1.0, 0.5})
	zfRef := -0.25
	zf := Filtic(b, a, []float64{0.0}, []float64{1.0})
	if zf[0] != zfRef {
		t.FailNow()
	}
}

func TestFilticHighPass2(t *testing.T) {
	b := MakeVectorFromArray([]float64{0.25, -0.25})
	a := MakeVectorFromArray([]float64{1.0, 0.5})
	zfRef := 0.0
	zf := Filtic(b, a, []float64{0.0}, nil)
	if zf[0] != zfRef {
		t.FailNow()
	}
}

func TestFiltic(t *testing.T) {
	b := MakeVectorFromArray([]float64{0.25, -0.25, 0.15, -0.12, 0.2, 0.2})
	a := MakeVectorFromArray([]float64{1.0, 0.5, 0.3, 0.1, 0.2, 0.3})
	zf := Filtic(b, a, []float64{1.0, 0.2, 0.7}, []float64{1.0, 2.0, 3.0})
	fmt.Printf("%v\n", zf)
	if math.Abs(zf[0] - -0.94) > 0.000001 {
		t.Errorf("Error 0 %f", zf[0])
	}
	if math.Abs(zf[1]-0.05) > 0.000001 {
		t.Errorf("Error 1 %f", zf[1])
	}
	if math.Abs(zf[2]-0.53) > 0.000001 {
		t.Errorf("Error 2 %f", zf[2])
	}
	if math.Abs(zf[3]-0.34) > 0.000001 {
		t.Errorf("Error 3 %f", zf[3])
	}
	if math.Abs(zf[4] - -0.1) > 0.000001 {
		t.Errorf("Error 4 %f", zf[4])
	}
}
