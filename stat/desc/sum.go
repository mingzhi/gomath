/*
 *   Copyright (C) 2012 Mingzhi Lin
 *
 * Permission is hereby granted, free of charge, to any person obtaining 
 * a copy of this software and associated documentation files (the "Software"), 
 * to deal in the Software without restriction, including without limitation 
 * the rights to use, copy, modify, merge, publish, distribute, sublicense, 
 * and/or sell copies of the Software, and to permit persons to whom the 
 * Software is furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included 
 * in all copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS 
 * OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, 
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE 
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER 
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, 
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE 
 * SOFTWARE.
 */
package desc

type Sum struct {
	n int
	v float64
}

func NewSum() (sum *Sum) {
	sum = &Sum{n: 0, v: 0}
	return
}

func (sum *Sum) Increment(x float64) {
	sum.v += x
	sum.n++
}

func (sum *Sum) GetResult() float64 {
	return sum.v
}

func (sum *Sum) GetN() int {
	return sum.n
}

func (sum *Sum) Clear() {
	sum.n = 0
	sum.v = 0
}

func (sum *Sum) IncrementAll(values []float64, begin, length int) {
	allowEmpty := true
	ok, err := test(values, begin, length, allowEmpty)
	if ok && err == nil {
		k := begin + length
		for i := begin; i < k; i++ {
			x := values[i]
			sum.Increment(x)
		}
	}
}

func (sum *Sum) IncrementAllWithWeigths(values, weights []float64, begin, length int) {
	allowEmpty := true
	ok, err := test(values, begin, length, allowEmpty)
	if ok && err == nil {
		k := begin + length
		for i := begin; i < k; i++ {
			x := values[i] * weights[i]
			sum.Increment(x)
		}
	}
}
