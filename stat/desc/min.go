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

import (
	"math"
)

type Min struct {
	n int
	v float64
}

func NewMin() *Min {
	return &Min{n: 0, v: math.NaN()}
}

func (min Min) GetN() int {
	return min.n
}

func (min Min) GetResult() float64 {
	return min.v
}

func (min *Min) Increment(x float64) {
	if x < min.v || math.IsNaN(min.v) {
		min.v = x
	}
	min.n++
}

func (min *Min) IncrementAll(values []float64, begin, length int) {
	allowEmpty := true
	ok, err := test(values, begin, length, allowEmpty)
	if ok && err == nil {
		for i := 0; i < begin+length; i++ {
			if !math.IsNaN(values[i]) {
				min.Increment(values[i])
			}
			min.n++
		}
	}
}

func EvaluateMin(values []float64, begin, length int) float64 {
	allowEmpty := true
	min := math.NaN()
	ok, err := test(values, begin, length, allowEmpty)
	if ok && err == nil {
		for i := 0; i < begin+length; i++ {
			if !math.IsNaN(values[i]) {
				if values[i] < min || math.IsNaN(min) {
					min = values[i]
				}
			}
		}
	}
	return min
}
