package regression

import (
	"math"
)

// Estimates an ordinary least squares regression model
// with one independent variable.
//
// y = intercept +slope * x
//
// Standard errors for interception and slope are
// available as well as ANOVA, r-square and Pearson's r statistics.
type Simple struct {
	sumX  float64 // sum of x values
	sumXX float64 // total variation in x (sum of squared deviations from xbar)
	sumY  float64 // sum of y values
	sumYY float64 // total variation in y (sum of squared deviations from ybar)
	sumXY float64 // sum of product
	n     int     // number of observations
	xbar  float64 // mean of accumulated x values, used in updating formulas
	ybar  float64 // mean of accumulated y values, used in updating formulas
}

func NewSimple() *Simple {
	return &Simple{}
}

// Adds the observation (x, y) to the regression data set.
// x is the independent variable value
// y is the dependent variable value
func (s *Simple) Add(x, y float64) {
	if s.n == 0 {
		s.xbar = x
		s.ybar = y
	} else {
		fact1 := 1.0 + float64(s.n)
		fact2 := float64(s.n) / (1.0 + float64(s.n))
		dx := x - s.xbar
		dy := y - s.ybar
		s.sumXX += dx * dx * fact2
		s.sumYY += dy * dy * fact2
		s.sumXY += dx * dy * fact2
		s.xbar += dx / fact1
		s.ybar += dy / fact1
	}
	s.sumX += x
	s.sumY += y
	s.n++
}

// Appends data from another regression calculation to this one.
func (s *Simple) Append(reg Simple) {
	if s.n == 0 {
		s.xbar = reg.xbar
		s.ybar = reg.ybar
		s.sumXX = reg.sumXX
		s.sumYY = reg.sumYY
		s.sumXY = reg.sumXY
	} else {
		fact1 := float64(reg.n) / float64(reg.n+s.n)
		fact2 := float64(s.n*reg.n) / float64(s.n+reg.n)
		dx := reg.xbar - s.xbar
		dy := reg.ybar - s.ybar
		s.sumXX += reg.sumXX + dx*dx*fact2
		s.sumYY += reg.sumYY + dy*dy*fact2
		s.sumXY += reg.sumXY + dx*dy*fact2
		s.xbar += dx * fact1
		s.ybar += dy * fact1
	}

	s.sumX += reg.sumX
	s.sumY += reg.sumY
	s.n += reg.n
}

// Removes the observation (x,y) from the regression data set.
// This is the reverse process of Update function.
// This method premits the use of Simple regression
// in streaming mode where the regression is applied
// to a sliding window of observations.
func (s *Simple) Remove(x, y float64) {
	if s.n > 0 {
		fact1 := float64(s.n) - 1.0
		fact2 := float64(s.n) / (float64(s.n) - 1.0)
		dx := x - s.xbar
		dy := y - s.ybar
		s.sumXX -= dx * dx * fact2
		s.sumYY -= dy * dy * fact2
		s.sumXY -= dx * dy * fact2
		s.xbar -= dx / fact1
		s.ybar -= dy / fact1
	}
	s.sumX -= x
	s.sumY -= y
	s.n--
}

// Clear all data from the model.
func (s *Simple) Clear() {
	s.sumX = 0
	s.sumXX = 0
	s.sumY = 0
	s.sumYY = 0
	s.sumXY = 0
	s.n = 0
}

func (s *Simple) N() int {
	return s.n
}

func (s *Simple) Intercept() float64 {
	return (s.sumY - s.Slope()*s.sumX) / float64(s.n)
}

func (s *Simple) Slope() float64 {
	if s.n < 2 {
		return math.NaN()
	}
	return s.sumXY / s.sumXX
}

// Returns the sum of squared errors (SSE) associated with the regression models.
// The sum is computed using the computational formula:
// SSE = SYY - (SXY * SXY / SXX)
// where SYY is the sum of the squared deviations of the y values
// about their mean, SXX is similarly defined
// and SXY is the sum of the products of x and y mean deviations.
// The return value is constrained to be non-negative.
// Preconditions:
// At least two observations (with at least two different x values)
// must have been added before invoking this method. If this method is
// invoked before a model can be estimated, NaN is returned.
func (s *Simple) SumSquaredErrors() float64 {
	return math.Max(0, s.sumYY-(s.sumXY*s.sumXY/s.sumXX))
}

// Returns the sum of squared deviations of the y values about their mean
func (s *Simple) TotalSumSquares() float64 {
	if s.N() < 2 {
		return math.NaN()
	}
	return s.sumYY
}

// Returns the sum of squared deviations of the x values about their mean
func (s *Simple) XSumSquares() float64 {
	if s.N() < 2 {
		return math.NaN()
	}
	return s.sumXX
}

// Returns the sum of crossproducts
func (s *Simple) SumOfCrossProducts() float64 {
	return s.sumXY
}

// Returns the sum of squared errors divided by the degrees of freedom,
// usually abbreviated MSE.
// If there are fewer than three data pairs in the model,
// or if there is no variation in x, this returns NaN
func (s *Simple) MeanSquareError() float64 {
	if s.N() < 3 {
		return math.NaN()
	}
	return s.SumSquaredErrors() / float64(s.N()-2)
}

// Returns the "predicted" y value associated with
// the supplied x value, based on the data that has been
// added to the model when this method is actived.
//
// predict(x) = intercept + slope * x
//
// Preconditions:
//
// At least two observations (with at least two different x values)
// must have been added before invoking this method. If this method is
// invoked before a model can be estimated, NaN, is returned.
func (s *Simple) Predict(x float64) {
	b1 := s.Slope()
	b0 := s.Intercept()
	return b0 + b1*x
}

// Returns <a href="http://mathworld.wolfram.com/CorrelationCoefficient.html">
// Pearson's product moment correlation coefficient</a>,
// usually denoted r.
// Preconditions
// At least two observations (with at least two different x values)
// must have been added before invoking this method. If this method is
// invoked before a model can be estimated, NaN is
// returned.
func (s *Simple) R() float64 {
	r := math.Sqrt(s.RSquare())
	if s.Slope() < 0 {
		r = -r
	}
	return r

}

// Returns the coefficient of determination
// usually denoted r-square.
// Preconditions:
// At least two observations (with at least two different x values)
// must have been added before invoking this method.
// If this method is called before a model can be estimated,
// NaN will be returned.
func (s *Simple) RSquare() float64 {
	ssto := s.TotalSumSquares()
	return (ssto - s.SumSquaredErrors()) / ssto
}

// Returns the standard error of the intercept estimate,
// usually denoted s(b0).
// If there are fewer that <strong>three</strong> observations in the
// model, or if there is no variation in x, this returns
// <code>Double.NaN</code>.</p> Additionally, a <code>Double.NaN</code> is
// returned when the intercept is constrained to be zero.
func (s *Simple) InterceptStdErr() float64 {
	return math.Sqrt(s.MeanSquareError() * ((1.0 / float64(s.N())) + (s.xbar*s.xbar)/s.sumXX))
}

// Returns the standard error of the slope estimate,
// usually denoted s(b1).
// If there are fewer that <strong>three</strong> observations in the
// model, or if there is no variation in x, this returns
// <code>Double.NaN</code>.</p> Additionally, a <code>Double.NaN</code> is
// returned when the intercept is constrained to be zero.
func (s *Simple) SlopeStdErr() float64 {
	return math.Sqrt(s.MeanSquareError() / s.sumXX)
}

// Computes SSR from the slope.
func (s *Simple) RegressionSumSquares() float64 {
	return s.Slope() * s.Slope() * s.sumXX
}
