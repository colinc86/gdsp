package gdsp

import "testing"

func TestMakeVector_1(t *testing.T) {
	n := 5
	r := 1.0
	i := 1.0
	v := MakeVectorComplex(complex(r, i), n)

	if len(v) != 5 {
		t.Errorf("Vector length %d should be %d.", len(v), n)
	} else {
		for _, c := range v {
			if real(c) != r || imag(c) != i {
				t.Errorf("Vector element %f + %fi should be %f + %fi.", real(c), imag(c), r, i)
				break
			}
		}
	}
}

func TestVMulE(t *testing.T) {
	u := MakeVector(2.0, 3)
	v := MakeVector(3.0, 3)
	w := VMulE(u, v)

	for _, r := range w {
		if r != 6.0 {
			t.Error("Error!")
		}
	}
}

func TestVMulESum(t *testing.T) {
	u := MakeVector(2.0, 3)
	v := MakeVector(3.0, 3)
	s := VMulESum(u, v)

	if s != 18.0 {
		t.Error("Error!")
	}
}

func TestVAdd(t *testing.T) {
	u := MakeVector(2.0, 3)
	v := MakeVector(3.0, 3)
	w := VAdd(u, v)

	for _, r := range w {
		if r != 5.0 {
			t.Error("Error!")
		}
	}
}

func TestVSub(t *testing.T) {
	u := MakeVector(2.0, 3)
	v := MakeVector(3.0, 3)
	w := VSub(u, v)

	for _, r := range w {
		if r != -1.0 {
			t.Error("Error!")
		}
	}
}

func TestVSMul(t *testing.T) {
	u := MakeVector(2.0, 3)
	w := VSMul(u, 3.0)

	for _, r := range w {
		if r != 6.0 {
			t.Error("Error!")
		}
	}
}

func TestVSDiv(t *testing.T) {
	u := MakeVector(2.0, 3)
	w := VSDiv(u, 2.0)

	for _, r := range w {
		if r != 1.0 {
			t.Error("Error!")
		}
	}
}

func TestVESum(t *testing.T) {
	u := MakeVector(2.0, 3)
	w := VESum(u)

	if w != 6.0 {
		t.Error("Error!")
	}
}

func TestVSumSq(t *testing.T) {
	u := MakeVector(2.0, 3)
	w := VSumSq(u)

	if w != 12.0 {
		t.Error("Error!")
	}
}

func TestVNeg(t *testing.T) {
	u := MakeVector(2.0, 3)
	w := VNeg(u)

	for _, r := range w {
		if r != -2.0 {
			t.Error("Error!")
		}
	}
}

func TestVectorCopy(t *testing.T) {
	u := MakeVector(2.0, 3)
	v := u.Copy()

	for _, r := range v {
		if r != 2.0 {
			t.Error("Error!")
		}
	}

	for i := range v {
		v[i] = 3.0
	}

	for _, r := range u {
		if r != 2.0 {
			t.Error("Error!")
		}
	}
}

func TestVectorReverse(t *testing.T) {
	u := MakeVector(0.0, 10)
	for i := 0; i < 10; i++ {
		u[i] = float64(i)
	}

	v := u.Reverse(8, 2)
	if v[0] != 8.0 || v[len(v)-1] != 2.0 {
		t.Error("Error!")
	}
}

func TestIsZero(t *testing.T) {
	v := MakeVector(0.0, 5)
	if !v.IsZero() {
		t.Error("Vector elements should all be zero!")
	}
}

func TestIsNotZero(t *testing.T) {
	v := MakeVector(1.0, 5)
	if v.IsZero() {
		t.Error("Vector elements should not be zero!")
	}
}

func TestPad(t *testing.T) {
	v := MakeVectorComplex(complex(1.0, 1.0), 5)
	pv := v.Padded(0.0, 2, 0.0, 2)
	if len(pv) != 9 {
		t.Errorf("Vector should have 9 elements (%d).", len(pv))
	}
}
