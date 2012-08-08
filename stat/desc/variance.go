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

import "math"

type Variance struct {
	isBiasCorrected bool
	moment          *SecondMoment
}

func NewVariance() *Variance {
	moment := NewSecondMoment()
	return &Variance{
		moment: moment,
	}
}

func NewVarianceWithBiasCorrection() *Variance {
	v := NewVariance()
	v.SetBiasCorrection(true)
	return v
}

func (v *Variance) Increment(d float64) {
	v.moment.Increment(d)
}

func (v *Variance) GetResult() (r float64) {
	if v.moment.GetN() == 0 {
		r = math.NaN()
	} else if v.moment.GetN() == 1 {
		r = 0.0
	} else {
		if v.isBiasCorrected {
			r = v.moment.GetResult() / (float64(v.moment.GetN() - 1.0))
		} else {
			r = v.moment.GetResult() / float64(v.moment.GetN())
		}
	}
	return
}

func (v *Variance) GetN() int {
	return v.moment.GetN()
}

func (v *Variance) Clear() {
	v.moment.Clear()
}

func (v *Variance) SetBiasCorrection(b bool) {
	v.isBiasCorrected = b
}
