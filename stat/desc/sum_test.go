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

func TestSumSpecialValues(t *testing.T) {
	sum := NewSum()
	if !assert.EqualFloat64(0.0, sum.GetResult(), 1e-10, 1) {
		t.Errorf("Sum: result: %f, but expect: %f\n", sum.GetResult(), 0.0)
	}

	sum.Increment(1.0)
	if !assert.EqualFloat64(1.0, sum.GetResult(), 1e-10, 1) {
		t.Errorf("Sum: result: %f, but expect: %f\n", sum.GetResult(), 1.0)
	}

	sum.Increment(math.Inf(0))
	if !assert.EqualFloat64(math.Inf(0), sum.GetResult(), 1e-10, 1) {
		t.Errorf("Sum: result: %f, but expect: %f\n", sum.GetResult(), math.Inf(0))
	}

	sum.Increment(math.Inf(-1))
	if !math.IsNaN(sum.GetResult()) {
		t.Error("Sum: result should be NaN")
	}

	sum.Increment(1.0)
	if !math.IsNaN(sum.GetResult()) {
		t.Error("Sum: result should be NaN")
	}
}
