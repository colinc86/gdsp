package gdsp

import (
	"fmt"
	"math"
	"testing"
)

func BenchmarkArburgC(t *testing.B) {
	v := []complex128{1.0, 2.0, 3.0, 4.0, 4.0, 3.0, 2.0, 1.0, 1.0, 2.0, 3.0, 4.0, 4.0, 3.0, 2.0, 1.0, 1.0, 2.0, 3.0, 4.0, 4.0, 3.0, 2.0, 1.0, 1.0, 2.0, 3.0, 4.0, 4.0, 3.0, 2.0, 1.0}

	for i := 0; i < t.N; i++ {
		ArburgC(v, len(v)-1)
	}
}

func TestArburg(t *testing.T) {
	vIn := []float64{0.7796877567724873, 0.7915302788549146, 0.8010646129962299, 0.8057222742623854, 0.8032915729866256, 0.7928232582592226, 0.775207484743498, 0.7525704516423012, 0.7272908506274501, 0.7012812282511067, 0.6764652490733608, 0.6546240491690392, 0.6365409527326914, 0.6219087759651647, 0.609896655287969, 0.5991957420539259}
	vOut := []float64{1.0000, -4.7023, 9.2221, -9.5862, 4.6628, 2.1789, -7.4553, 9.8842, -10.2159, 8.3933, -3.5362, -3.2161, 8.4095, -8.5107, 4.4159, -0.9439}
	eOut := 2.6695e-09
	v, e := Arburg(vIn, 15)
	acceptableDifference := 0.0001

	for i := range v {
		if math.Abs(v[i]-vOut[i]) > acceptableDifference {
			t.Errorf("%f at %d should be %f.", v[i], i, vOut[i])
		}
	}

	if math.Abs(e-eOut) > acceptableDifference {
		t.Errorf("%f should be %f.", e, eOut)
	}
}

func TestArburgC(t *testing.T) {
	v, e := ArburgC([]complex128{complex(2.0, 3.0), complex(1.0, -1.0), complex(4.0, 2.0), complex(3.0, -2.0)}, 3)
	fmt.Printf("%v, %f", v, e)
}
