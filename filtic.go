package gdsp

import (
	"runtime"

	"github.com/colinc86/parallel"
)

// Filtic creates the initial condition vector for the filter function. x (optional)
// and y contain the last input and output from the filter function. a and b are
// the filter coefficients.
func Filtic(b Vector, a Vector, y Vector, x Vector) Vector {
	na := len(a)
	nb := len(b)
	m := MaxI(na, nb) - 1

	if m < 1 {
		return MakeVector(0.0, 0)
	}

	xc := x.Copy()
	if len(xc) < nb-1 {
		xc = xc.PaddedTrailing(0.0, nb-1-len(x))
	}

	yc := y.Copy()
	if len(yc) < na-1 {
		yc = yc.PaddedTrailing(0.0, na-1-len(y))
	}

	vinit := MakeVector(0.0, m)
	vx := MakeVector(0.0, m)

	if na-1 > 0 {
		bf := VSDiv(VNeg(a.Reverse(na-1, 1)), a[0])
		af := MakeVector(1.0, 1).PaddedTrailing(0.0, len(bf)-1)
		yf := yc.SubVector(0, na-1)
		f, _ := Filter(bf, af, yf, nil)

		p := parallel.NewFixedProcess(runtime.NumCPU())
		p.Execute(len(f), func(i int) {
			vIndex := na - 2 - i
			vinit[vIndex] = f[i]
		})
	}

	if nb-1 > 0 {
		bf := VSDiv(b.Reverse(nb-1, 1), a[0])
		af := MakeVector(1.0, 1).PaddedTrailing(0.0, len(bf)-1)
		xf := xc.SubVector(0, nb-1)
		f, _ := Filter(bf, af, xf, nil)

		p := parallel.NewFixedProcess(runtime.NumCPU())
		p.Execute(len(f), func(i int) {
			vIndex := nb - 2 - i
			vx[vIndex] = f[i]
		})
	}

	return VAdd(vinit, vx)
}

// FilticC creates the initial condition vector for the filter function. x (optional)
// and y contain the last input and output from the filter function. a and b are
// the filter coefficients.
func FilticC(b VectorComplex, a VectorComplex, y VectorComplex, x VectorComplex) VectorComplex {
	na := len(a)
	nb := len(b)
	m := MaxI(na, nb) - 1

	if m < 1 {
		return MakeVectorComplex(0.0, 0)
	}

	xc := x.Copy()
	if len(xc) < nb-1 {
		xc = xc.PaddedTrailing(0.0, nb-1-len(x))
	}

	yc := y.Copy()
	if len(yc) < na-1 {
		yc = yc.PaddedTrailing(0.0, na-1-len(y))
	}

	vinit := MakeVectorComplex(0.0, m)
	vx := MakeVectorComplex(0.0, m)

	if na-1 > 0 {
		bf := VSDivC(VNegC(a.Reverse(na-1, 1)), a[0])
		af := MakeVectorComplex(1.0, 1).PaddedTrailing(0.0, len(bf)-1)
		yf := yc.SubVector(0, na-1)
		f, _ := FilterC(bf, af, yf, nil)

		p := parallel.NewFixedProcess(runtime.NumCPU())
		p.Execute(len(f), func(i int) {
			vIndex := na - 2 - i
			vinit[vIndex] = f[i]
		})
	}

	if nb-1 > 0 {
		bf := VSDivC(b.Reverse(nb-1, 1), a[0])
		af := MakeVectorComplex(1.0, 1).PaddedTrailing(0.0, len(bf)-1)
		xf := xc.SubVector(0, nb-1)
		f, _ := FilterC(bf, af, xf, nil)

		p := parallel.NewFixedProcess(runtime.NumCPU())
		p.Execute(len(f), func(i int) {
			vIndex := nb - 2 - i
			vx[vIndex] = f[i]
		})
	}

	return VAddC(vinit, vx)
}
