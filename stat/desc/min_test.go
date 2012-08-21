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

func TestMinSpecialValues(t *testing.T) {
	min := NewMin()
	if !math.IsNaN(min.GetResult()) {
		t.Error("Min: the result should be NaN")
	}
	if min.GetN() != 0 {
		t.Error("Min: the N of Min should be 0")
	}

	testArray := []float64{0.0, math.NaN(), math.Inf(0), math.Inf(-1)}
	min.Increment(testArray[0])
	if !assert.EqualFloat64(min.GetResult(), 0.0, 1e-10, 1) {
		t.Errorf("Min: result: %f, but expect: %f", min.GetResult(), 0.0)
	}

	min.Increment(testArray[1])
	if !assert.EqualFloat64(min.GetResult(), 0.0, 1e-10, 1) {
		t.Errorf("Min: result: %f, but expect: %f", min.GetResult(), 0.0)
	}

	min.Increment(testArray[2])
	if !assert.EqualFloat64(min.GetResult(), 0.0, 1e-10, 1) {
		t.Errorf("Min: result: %f, but expect: %f", min.GetResult(), 0.0)
	}

	min.Increment(testArray[3])
	if !assert.EqualFloat64(min.GetResult(), math.Inf(-1), 1e-10, 1) {
		t.Errorf("Min: result: %f, but expect: %f", min.GetResult(), math.Inf(-1))
	}

	if min.GetN() != len(testArray) {
		t.Errorf("Min: N: %d, but expect: %d", min.GetN(), len(testArray))
	}

	minV := EvaluateMin(testArray, 0, len(testArray))
	if !assert.EqualFloat64(minV, math.Inf(-1), 1e-10, 1) {
		t.Errorf("Min: result: %f, but expect: %f", minV, math.Inf(-1))
	}
}
