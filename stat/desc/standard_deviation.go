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

type StandardDeviation struct {
	vr *Variance
}

func NewStandardDeviation() *StandardDeviation {
	vr := NewVariance()
	return &StandardDeviation{vr: vr}
}

func NewStandardDeviationWithBiasCorrection() *StandardDeviation {
	sd := NewStandardDeviation()
	sd.SetBiasCorrection(true)
	return sd
}

func (sd *StandardDeviation) Increment(d float64) {
	sd.vr.Increment(d)
}

func (sd *StandardDeviation) GetResult() float64 {
	return math.Sqrt(sd.vr.GetResult())
}

func (sd *StandardDeviation) GetN() int {
	return sd.vr.GetN()
}

func (sd *StandardDeviation) Clear() {
	sd.vr.Clear()
}

func (sd *StandardDeviation) SetBiasCorrection(b bool) {
	sd.vr.SetBiasCorrection(b)
}
