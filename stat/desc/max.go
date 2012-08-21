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

type Max struct {
	n int
	v float64
}

func NewMax() *Max {
	return &Max{n: 0, v: math.NaN()}
}

func (max *Max) Increment(x float64) {
	if x > max.v || math.IsNaN(max.v) {
		max.v = x
	}
	max.n++
}

func (max *Max) IncrementAll(values []float64, begin, length int) {
	allowEmpty := true
	ok, err := test(values, begin, length, allowEmpty)
	if ok && err == nil {
		for i := begin; i < begin+length; i++ {
			if !math.IsNaN(values[i]) {
				max.Increment(values[i])
			}
		}
	}
}

func (max *Max) Clean() {
	max.n = 0
	max.v = math.NaN()
}

func (max *Max) GetResult() float64 {
	return max.v
}

func (max *Max) GetN() int {
	return max.n
}

func EvaluateMax(values []float64, begin, length int) (max float64) {
	max = math.NaN()

	allowEmpty := true
	ok, err := test(values, begin, length, allowEmpty)
	if ok && err == nil {
		max = values[begin]
		for i := begin; i < begin+length; i++ {
			if !math.IsNaN(values[i]) {
				if max < values[i] || math.IsNaN(max) {
					max = values[i]
				}
			}
		}
	}

	return
}
