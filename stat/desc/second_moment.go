package desc

import (
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
	if n < 1 {
		sm.moment.m1 = 0
		sm.m2 = 0
	}
	sm.moment.Increment(d)
	m2 += (float64(n) - 1.0) * sm.moment.dev * sm.moment.nDev
}

func (sm *SecondMoment) Clear() {
	sm.moment.Clear()
	sm.m2 = math.NaN()
}

func (sm *SecondMoment) GetResult() float64 {
	return sm.m2
}

func (sm *SecondMoment) GetN() float64 {
	return sm.moment.GetN()
}
