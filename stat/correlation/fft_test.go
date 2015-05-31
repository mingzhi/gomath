package correlation

import (
	"github.com/mingzhi/fftw"
	"math"
	"math/rand"
	"testing"
)

const tolerance float64 = 1e-5

func TestAutoCorr(t *testing.T) {
	data := []float64{
		0.1576, 0.9706, 0.9572, 0.4854, 0.8003, 0.1419, 0.4218, 0.9157, 0.7922, 0.9595,
	}

	var expected []float64
	for _, circular := range []bool{true, false} {
		if circular {
			expected = []float64{5.34401, 4.13152, 4.19332, 4.27546, 4.06152, 4.92138, 4.06152, 4.27546, 4.19332, 4.13152}
		} else {
			expected = []float64{5.34401, 3.98031, 3.13718, 2.4438, 1.88223, 2.46069, 2.17929, 1.83166, 1.05614, 0.151217}
		}
		res1 := AutoCorrBruteForce(data, circular)
		res2 := AutoCorrFFT(data, circular)
		if len(res1) != len(res2) {
			t.Errorf("Results 1 length of %d, results 2 length of %d, circular is %v\n", len(res1), len(res2), circular)
		}
		for i := 0; i < len(expected); i++ {
			if math.Abs(res1[i]-res2[i]) > tolerance {
				t.Errorf("Result 1 %f, result 2 %f, at %d, circular is %v\n", res2[i], res1[i], i, circular)
			}
			if math.Abs(res1[i]-expected[i]) > tolerance {
				t.Errorf("Expected %f, got %f, at %d, circular is %v\n", expected[i], res1[i], i, circular)
			}
		}
	}

}

func TestXCorr(t *testing.T) {

	data1 := []float64{
		0.6557,
		0.0357,
		0.8491,
		0.9340,
		0.6787,
		0.7577,
		0.7431,
		0.3922,
		0.6555,
		0.1712,
	}
	data2 := []float64{
		0.1576,
		0.9706,
		0.9572,
		0.4854,
		0.8003,
		0.1419,
		0.4218,
		0.9157,
		0.7922,
		0.9595,
	}

	var expected []float64
	for _, circular := range []bool{true, false} {
		if circular {
			expected = []float64{3.41092, 3.89322, 3.67161, 3.65795, 4.21624, 3.94802, 4.17104, 3.81444, 3.69511, 4.2955}
		} else {
			expected = []float64{3.41092, 3.86624, 3.40214, 2.79604, 3.00792, 2.27675, 1.87809, 1.44342, 0.5537, 0.629144}
		}

		dft := NewFFTW(len(data1), fftw.OutOfPlace, fftw.Measure, circular)
		defer dft.Close()

		res1 := XCorrBruteForce(data1, data2, circular)
		res2 := XCorrFFT(data1, data2, circular)
		res3 := dft.XCorr(data1, data2)
		if len(res1) != len(res2) || len(res1) != len(res3) {
			t.Errorf("Results 1 length of %d, results 2 length of %d, circular is %v\n", len(res1), len(res2), circular)
		}
		for i := 0; i < len(expected); i++ {
			if math.Abs(res1[i]-res2[i]) > tolerance {
				t.Errorf("Result 2 %f, result 1 %f, at %d, circular is %v\n", res2[i], res1[i], i, circular)
			}
			if math.Abs(res1[i]-expected[i]) > tolerance {
				t.Errorf("Expected %f, got %f, at %d, circular is %v\n", expected[i], res1[i], i, circular)
			}
			if math.Abs(res1[i]-res3[i]) > tolerance {
				t.Errorf("Result 3 %f, result 1 %f, at %d, circular is %v\n", res3[i], res1[i], i, circular)
			}
		}
	}

}

func BenchmarkFFTAuto(b *testing.B) {
	data := make([]float64, 510)
	for i := 0; i < len(data); i++ {
		data[i] = rand.Float64()
	}

	for i := 0; i < b.N; i++ {
		circular := true
		AutoCorrFFT(data, circular)
	}

}

func BenchmarkFFTBF(b *testing.B) {
	data := make([]float64, 510)
	for i := 0; i < len(data); i++ {
		data[i] = rand.Float64()
	}

	for i := 0; i < b.N; i++ {
		circular := true
		AutoCorrBruteForce(data, circular)
	}

}
