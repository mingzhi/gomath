//   Copyright (C) 2012 Mingzhi Lin
//
// Permission is hereby granted, free of charge, to any person obtaining 
// a copy of this software and associated documentation files (the "Software"), 
// to deal in the Software without restriction, including without limitation 
// the rights to use, copy, modify, merge, publish, distribute, sublicense, 
// and/or sell copies of the Software, and to permit persons to whom the 
// Software is furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included 
// in all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS 
// OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, 
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE 
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER 
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, 
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE 
// SOFTWARE.

package random

import (
	"fmt"
	"math"
	"math/rand"
	"testing"
)

func TestPoisson(t *testing.T) {
	src := rand.NewSource(1)

	// < SWITCH_MEAN
	mean := 9.9998
	poisson := NewPoisson(mean, src)
	err := testDiscreteDistribution(poisson)
	if err != nil {
		t.Error(err)
	}

	// >= SWITCH_MEAN
	mean = 99.747564
	poisson = NewPoisson(mean, src)
	err = testDiscreteDistribution(poisson)
	if err != nil {
		t.Error(err)
	}

}

func TestExponential(t *testing.T) {
	src := rand.NewSource(1)

	lambda := 2.3
	exp := NewExponential(lambda, src)
	err := testContinuousDistribution(exp)
	if err != nil {
		t.Error(err)
	}
}

func TestBinomial(t *testing.T) {
	src := rand.NewSource(1)

	n := 100000
	p := 0.004
	binomial := NewBinomial(n, p, src)
	err := testDiscreteDistribution(binomial)
	if err != nil {
		t.Error(err)
	}

	n = 100000
	p = 0.00004
	binomial = NewBinomial(n, p, src)
	err = testDiscreteDistribution(binomial)
	if err != nil {
		t.Error(err)
	}
}

func BenchmarkPoisson(b *testing.B) {
	b.StopTimer()
	src := rand.NewSource(1)
	rng := rand.New(src)
	mean := 99.0
	poisson := NewPoisson(mean, rng)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		poisson.Int()
	}
}

func testContinuousDistribution(dd ContinuousDistribution) (err error) {
	n := 1000000
	bins := 100
	count := make([]float64, bins)
	p := make([]float64, bins)

	var status_i bool

	for i := 0; i < n; i++ {
		r := dd.Float64()
		if int(r) >= 0 && int(r) < bins {
			count[int(r)]++
		}
	}

	for i := 0; i < bins; i++ {
		p[i] = dd.Cdf(float64(i+1)) - dd.Cdf(float64(i)) - dd.Pdf(float64(i))
	}

	for i := 0; i < bins; i++ {
		d := math.Abs(count[i] - float64(n)*p[i])
		if p[i] != 0 {
			s := d / math.Sqrt(float64(n)*p[i])
			status_i = (s > 5) && (d > 1)
		} else {
			status_i = count[i] != 0
		}

		if status_i {
			errmessage := fmt.Sprintf("i=%d (%g observed vs %g expected)", i, count[i]/float64(n), p[i])
			err = Error{Message: errmessage}
			return
		}
	}
	return
}

func testDiscreteDistribution(dd DiscreteDistricution) (err error) {
	n := 1000000
	bins := 100
	count := make([]float64, bins)
	p := make([]float64, bins)

	var status_i bool

	for i := 0; i < n; i++ {
		r := dd.Int()
		if r >= 0 && r < bins {
			count[r]++
		}
	}

	for i := 0; i < bins; i++ {
		p[i] = dd.Pdf(i)
	}

	for i := 0; i < bins; i++ {
		d := math.Abs(count[i] - float64(n)*p[i])
		if p[i] != 0 {
			s := d / math.Sqrt(float64(n)*p[i])
			status_i = (s > 5) && (d > 1)
		} else {
			status_i = count[i] != 0
		}

		if status_i {
			errmessage := fmt.Sprintf("i=%d (%g observed vs %g expected)", i, count[i]/float64(n), p[i])
			err = Error{Message: errmessage}
			return
		}
	}
	return
}
