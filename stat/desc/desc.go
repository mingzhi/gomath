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
	"fmt"
)

type StorelessUnivariateStatistic interface {
	Increment(float64)
	GetResult() float64
	GetN() int
	Clear()
}

type Error struct {
	Message string
	Status  int
}

const (
	_ = iota
	NotPositive
	NumberIsTooLarge
	DimentionMismatch
)

func (err Error) Error() string {
	return fmt.Sprintf("Message: %s, Status: %d", err.Message, err.Status)
}

func test(values []float64, begin, length int, allowEmpty bool) (ok bool, err error) {
	if begin < 0 {
		err = Error{Message: "Begin index is not positive", Status: NotPositive}
		return
	}

	if length < 0 {
		err = Error{Message: "Length is not positive", Status: NotPositive}
		return
	}

	if begin+length > len(values) {
		err = Error{Message: "Number is too large", Status: NumberIsTooLarge}
		return
	}

	if length == 0 && !allowEmpty {
		return
	}

	ok = true
	return
}
