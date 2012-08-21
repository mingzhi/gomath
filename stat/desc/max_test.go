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
	"github.com/mingzhi/goutils/assert"
	"math"
	"testing"
)

func TestMaxSpecialValues(t *testing.T) {
	testArray := []float64{0.0, math.NaN(), math.Inf(-1), math.Inf(0)}
	max := NewMax()
	if !math.IsNaN(max.GetResult()) {
		t.Error("Max: it should be NaN")
	}

	max.Increment(testArray[0])
	if !assert.EqualFloat64(max.GetResult(), 0.0, 1e-10, 1) {
		t.Errorf("Max: result: %f, but expect: %f", max.GetResult(), 0.0)
	}

	max.Increment(testArray[1])
	if !assert.EqualFloat64(max.GetResult(), 0.0, 1e-10, 1) {
		t.Errorf("Max: result: %f, but expect: %f", max.GetResult(), 0.0)
	}

	max.Increment(testArray[2])
	if !assert.EqualFloat64(max.GetResult(), 0.0, 1e-10, 1) {
		t.Errorf("Max: result: %f, but expect: %f", max.GetResult(), 0.0)
	}

	max.Increment(testArray[3])
	if !assert.EqualFloat64(max.GetResult(), math.Inf(0), 1e-10, 1) {
		t.Errorf("Max: result: %f, but expect: %f", max.GetResult(), math.Inf(0))
	}

	maxV := EvaluateMax(testArray, 0, len(testArray))
	if !assert.EqualFloat64(maxV, math.Inf(0), 1e-10, 1) {
		t.Errorf("Max: result: %f, but expect: %f", maxV, math.Inf(0))
	}
}
