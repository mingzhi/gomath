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
	"math"
)

type SecondMoment struct {
	moment *FirstMoment
	m2     float64
}

func NewSecondMoment() *SecondMoment {
	moment := NewFirstMoment()
	return &SecondMoment{
		moment: moment,
		m2:     math.NaN(),
	}
}

func (sm *SecondMoment) Increment(d float64) {
	if sm.moment.n < 1 {
		sm.moment.m1 = 0
		sm.m2 = 0
	}
	sm.moment.Increment(d)
	sm.m2 += (float64(sm.moment.n) - 1.0) * sm.moment.dev * sm.moment.nDev
}

func (sm *SecondMoment) Append(s2 *SecondMoment) {
	delta := sm.moment.GetResult() - s2.moment.GetResult()
	nA := sm.GetN()
	nB := s2.GetN()
	if nA == 0 {
		sm.moment.Append(s2.moment)
		sm.m2 = s2.m2
	} else if nB != 0 {
		sm.moment.Append(s2.moment)
		sm.m2 += s2.m2 + (delta*delta)*float64(nA*nB)/float64(nA+nB)
	}
}

func (sm *SecondMoment) Clear() {
	sm.moment.Clear()
	sm.m2 = math.NaN()
}

func (sm *SecondMoment) GetResult() float64 {
	return sm.m2
}

func (sm *SecondMoment) GetN() int {
	return sm.moment.GetN()
}

type secondMoment struct {
	Moment *FirstMoment
	M2     float64
}

func (sm *SecondMoment) MarshalBinary() ([]byte, error) {
	s := secondMoment{Moment: sm.moment, M2: sm.m2}
	var b1 bytes.Buffer
	enc := gob.NewEncoder(&b1)
	err := enc.Encode(s)
	return b1.Bytes(), err
}

func (sm *SecondMoment) UnmarshalBinary(data []byte) error {
	var s secondMoment
	b := bytes.NewBuffer(data)
	dec := gob.NewDecoder(b)
	err := dec.Decode(&s)

	sm.m2 = s.M2
	sm.moment = s.Moment

	return err
}
