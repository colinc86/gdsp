package gdsp

import (
	"fmt"
	"math/cmplx"
)

// Arburg finds the autoregressive parameters for a model with order p using Burg's
// method on the real-valued input vector, x, and returns the parameters along with
// the estimated variance.
func Arburg(x Vector, p int) (Vector, float64) {
	if len(x) < p+1 {
		fmt.Printf("Vector length %d should be greater than p, %d.", len(x), p)
	}

	N := len(x)
	Kout := MakeVector(0.0, p)

	efp := x.SubVector(1, len(x))
	ebp := x.SubVector(0, len(x)-1)

	E := VMulESum(x, x) / float64(N)

	a := MakeVector(0.0, p+1)
	a[0] = 1.0

	for m := 0; m < p; m++ {
		k := -2.0 * VMulESum(ebp, efp) / (VMulESum(efp, efp) + VMulESum(ebp, ebp))
		Kout[m] = k

		ef := VAdd(efp.SubVector(1, len(efp)), VSMul(ebp.SubVector(1, len(ebp)), k))
		ebp = VAdd(ebp.SubVector(0, len(efp)-1), VSMul(efp.SubVector(0, len(ebp)-1), k))
		efp = ef

		revA := a.Reverse(m, 0)
		for i := 1; i <= m+1; i++ {
			a[i] += k * revA[i-1]
		}

		E = (1.0 - k*k) * E
	}

	return a, E
}

// ArburgC finds the autoregressive parameters for a model with order p using Burg's
// method on the complex-valued input vector, x, and returns the parameters along
// with the estimated variance.
func ArburgC(x VectorComplex, p int) (VectorComplex, complex128) {
	if len(x) < p+1 {
		fmt.Printf("Vector length %d should be greater than p, %d.", len(x), p)
	}

	N := len(x)
	Kout := MakeVectorComplex(0.0, p)

	efp := x.SubVector(1, len(x))
	ebp := x.SubVector(0, len(x)-1)

	E := VMulESumC(x.Conj(), x) / ComplexRI(N)

	a := MakeVectorComplex(0.0, p+1)
	a[0] = 1.0

	for m := 0; m < p; m++ {
		d := real(VMulESumC(efp.Conj(), efp) + VMulESumC(ebp.Conj(), ebp))
		q := VMulESumC(ebp.Conj(), efp)
		qr := -2.0 * real(q) / d
		qi := -2.0 * imag(q) / d
		k := complex(qr, qi)
		Kout[m] = k

		ef := VAddC(efp.SubVector(1, len(efp)), VSMulC(ebp.SubVector(1, len(ebp)), k))
		ebp = VAddC(ebp.SubVector(0, len(efp)-1), VSMulC(efp.SubVector(0, len(ebp)-1), cmplx.Conj(k)))
		efp = ef

		revA := a.Reverse(m, 0).Conj()
		for i := 1; i <= m+1; i++ {
			a[i] += k * revA[i-1]
		}

		E = (1.0 - cmplx.Conj(k)*k) * E
	}

	return a, E
}
