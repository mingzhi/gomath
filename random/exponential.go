package random

import (
	"math"
	"math/rand"
)

type Exponential struct {
	Lambda          float64
	randomGenerator *rand.Rand
}

// NewExponential returns a Exponential distribution with the provided lambda.
func NewExponential(lambda float64, src rand.Source) *Exponential {
	exp := Exponential{Lambda: lambda}
	exp.randomGenerator = rand.New(src)
	return &exp
}

// Cdf returns the cumulative distribution function.
func (exp Exponential) Cdf(x float64) (p float64) {
	if x <= 0.0 {
		p = 0.0
	} else {
		p = 1.0 - math.Exp(-x*exp.Lambda)
	}
	return
}

// Float64 returns a random number from the distribution.
func (exp Exponential) Float64() float64 {
	return rand.ExpFloat64() / exp.Lambda
}

// Pdf returns the probability distribution function.
func (exp Exponential) Pdf(x float64) (p float64) {
	if x >= 0.0 {
		p = exp.Lambda * math.Exp(-x*exp.Lambda)
	} else {
		p = 0.0
	}
	return
}

// Seed uses the provided seed value to initialize the generator to a deterministic state.
func (exp *Exponential) Seed(seed int64) {
	exp.randomGenerator.Seed(seed)
}
