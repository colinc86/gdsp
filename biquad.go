package dsp

import (
	"math"
)

// BiquadType defines a type of biquad.
type BiquadType int

// A group of biquad types.
const (
	BiquadTypeLowpass  BiquadType = 0
	BiquadTypeHighpass BiquadType = 1
	BiquadTypeBandpass BiquadType = 2
)

// BiquadSetup contains the parameters to perform a biquad filter on a signal.
type BiquadSetup struct {
	// The filter type.
	Type BiquadType

	// A and B coefficients.
	A0, A1, A2, B1, B2 float64

	// Filter Q and cutoff frequency.
	Q, Fc float64

	// Z coefficients.
	Z1, Z2 float64
}

// MARK: Constructors

// NewBiquadSetup creates and returns a new biquad setup.
func NewBiquadSetup(t BiquadType, q float64, fc float64) *BiquadSetup {
	setup := new(BiquadSetup)
	setup.Type = t
	setup.Q = q
	setup.Fc = fc
	setup.A0 = 1.0

	var norm float64
	var K = math.Tan(math.Pi * fc)

	switch t {
	case BiquadTypeLowpass:
		norm = 1.0 / (1.0 + K/q + K*K)
		setup.A0 = K * K * norm
		setup.A1 = 2.0 * setup.A0
		setup.A2 = setup.A0
		setup.B1 = 2.0 * (K*K - 1.0) * norm
		setup.B2 = (1.0 - K/q + K*K) * norm
	case BiquadTypeHighpass:
		norm = 1.0 / (1.0 + K/q + K*K)
		setup.A0 = 1.0 * norm
		setup.A1 = -2.0 * setup.A0
		setup.A2 = setup.A0
		setup.B1 = 2.0 * (K*K - 1.0) * norm
		setup.B2 = (1.0 - K/q + K*K) * norm
	case BiquadTypeBandpass:
		norm = 1.0 / (1.0 + K/q + K*K)
		setup.A0 = K / q * norm
		setup.A1 = 0
		setup.A2 = -setup.A0
		setup.B1 = 2.0 * (K*K - 1.0) * norm
		setup.B2 = (1.0 - K/q + K*K) * norm
	}

	return setup
}

// MARK: Public methods

// Perform performs a biquad filter on a singal and returns the result.
func (s *BiquadSetup) Perform(a []float64) []float64 {
	s.Z1 = 0.0
	s.Z2 = 0.0

	output := make([]float64, len(a), len(a))
	for n := 0; n < len(a); n++ {
		output[n] = a[n]*s.A0 + s.Z1
		s.Z1 = a[n]*s.A1 + s.Z2 - s.B1*output[n]
		s.Z2 = a[n]*s.A2 - s.B2*output[n]
	}

	return output
}
