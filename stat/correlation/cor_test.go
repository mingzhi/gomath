package correlation

import (
	"math"
	"testing"
)

var (
	longleyDataSimple [][]float64
)

func init() {
	longleyDataSimple = [][]float64{
		[]float64{60323, 83.0},
		[]float64{61122, 88.5},
		[]float64{60171, 88.2},
		[]float64{61187, 89.5},
		[]float64{63221, 96.2},
		[]float64{63639, 98.1},
		[]float64{64989, 99.0},
		[]float64{63761, 100.0},
		[]float64{66019, 101.2},
		[]float64{67857, 104.6},
		[]float64{68169, 108.4},
		[]float64{66513, 110.8},
		[]float64{68655, 112.6},
		[]float64{69564, 114.2},
		[]float64{69331, 115.7},
		[]float64{70551, 116.9},
	}
}

func TestLonglySimpleVar(t *testing.T) {
	biasCorrected := true
	delta := 1e-7
	{
		cov := NewBivariateCovariance(biasCorrected)
		for i := 0; i < len(longleyDataSimple); i++ {
			cov.Increment(longleyDataSimple[i][0], longleyDataSimple[i][0])
		}
		expected := 12333921.73333333246
		actual := cov.GetResult()
		if !EqualFloat64(actual, expected, delta, 0) {
			t.Errorf("BivariateCovariance Increment, result: %.10f, expected: %.10f, number of data: %d\n", actual, expected, cov.GetN())
		}
	}

	{
		cov := NewBivariateCovariance(biasCorrected)
		for i := 0; i < len(longleyDataSimple); i++ {
			cov.Increment(longleyDataSimple[i][0], longleyDataSimple[i][1])
		}
		expected := 36796.660000
		actual := cov.GetResult()
		if !EqualFloat64(actual, expected, delta, 0) {
			t.Errorf("BivariateCovariance Increment, result: %.10f, expected: %.10f, number of data: %d\n", actual, expected, cov.GetN())
		}
	}

}

// Asserts that two floats are equal to within a positive delta.
// typ = 0: Absolute delta; typ = 1: Relative delta.
// NaNs or Infs are considerred equal.
func EqualFloat64(actual, expected, delta float64, typ int) (status bool) {
	switch {
	case math.IsNaN(actual) || math.IsNaN(expected):
		status = math.IsNaN(actual) == math.IsNaN(expected)
		break
	case math.IsInf(actual, 0) || math.IsInf(expected, 0):
		status = math.IsInf(actual, 0) == math.IsInf(expected, 0)
		break
	case expected == 0:
		status = math.Abs(actual-expected) < math.Abs(delta)
		break
	case expected != 0:
		if typ == 0 {
			status = math.Abs(actual-expected) < math.Abs(delta)
		} else {
			status = math.Abs(actual-expected)/math.Abs(expected) < math.Abs(delta)
		}
	}
	return
}
