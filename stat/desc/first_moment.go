package desc

import (
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
	fm.nDev = fm.dev / n0
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
