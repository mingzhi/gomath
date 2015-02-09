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
	"bytes"
	"fmt"
	"math"
)

type FirstMoment struct {
	n    int
	m1   float64
	dev  float64
	nDev float64
}

func NewFirstMoment() *FirstMoment {
	return &FirstMoment{
		n:    0,
		m1:   math.NaN(),
		dev:  math.NaN(),
		nDev: math.NaN(),
	}
}

func (fm *FirstMoment) Increment(d float64) {
	if fm.n == 0 {
		fm.m1 = 0
	}
	fm.n++
	n0 := fm.n
	fm.dev = d - fm.m1
	fm.nDev = fm.dev / float64(n0)
	fm.m1 += fm.nDev
}

func (fm *FirstMoment) Clear() {
	fm.m1 = math.NaN()
	fm.n = 0
	fm.dev = math.NaN()
	fm.nDev = math.NaN()
}

func (fm *FirstMoment) GetResult() float64 {
	return fm.m1
}

func (fm *FirstMoment) GetN() int {
	return fm.n
}

func (fm *FirstMoment) Append(fm2 *FirstMoment) {
	if fm.n+fm2.n > 0 {
		fm.m1 = (fm.m1*float64(fm.n) + fm2.m1*float64(fm2.n)) / float64(fm.n+fm2.n)
		fm.n = fm.n + fm2.n
	}
}

func (f *FirstMoment) MarshalBinary() ([]byte, error) {
	var b bytes.Buffer
	fmt.Fprintln(&b, f.m1, f.dev, f.nDev, f.n)
	return b.Bytes(), nil
}

func (f *FirstMoment) UnmarshalBinary(data []byte) error {
	b := bytes.NewBuffer(data)
	_, err := fmt.Fscanln(b, &f.m1, &f.dev, &f.nDev, &f.n)
	return err
}
