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
	"encoding/gob"
)

type Mean struct {
	moment *FirstMoment
}

func NewMean() *Mean {
	moment := NewFirstMoment()
	return &Mean{
		moment: moment,
	}
}

func (m *Mean) Increment(d float64) {
	m.moment.Increment(d)
}

func (m *Mean) Clear() {
	m.moment.Clear()
}

func (m *Mean) GetResult() float64 {
	return m.moment.GetResult()
}

func (m *Mean) GetN() int {
	return m.moment.GetN()
}

func (m *Mean) Append(m2 *Mean) {
	m.moment.Append(m2.moment)
}

func (f *Mean) MarshalBinary() ([]byte, error) {
	var b bytes.Buffer
	enc := gob.NewEncoder(&b)
	if err := enc.Encode(f.moment); err != nil {
		panic(err)
	}
	return b.Bytes(), nil
}

func (f *Mean) UnmarshalBinary(data []byte) error {
	b := bytes.NewBuffer(data)
	dec := gob.NewDecoder(b)
	err := dec.Decode(&f.moment)
	return err
}
