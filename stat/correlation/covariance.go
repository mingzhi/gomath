package correlation

import (
	"bytes"
	"fmt"
	"math"
)

type BivariateCovariance struct {
	meanX, meanY  float64 // the mean of variable x and y
	n             int64   // number of observations
	estimator     float64 // the running covariance estimate
	biasCorrected bool    // flag for bias correction
}

func NewBivariateCovariance(biasCorrected bool) *BivariateCovariance {
	cov := BivariateCovariance{}
	cov.biasCorrected = biasCorrected
	return &cov
}

func (cov *BivariateCovariance) Increment(x, y float64) {
	cov.n++
	deltaX := x - cov.meanX
	deltaY := y - cov.meanY
	cov.meanX += deltaX / float64(cov.n)
	cov.meanY += deltaY / float64(cov.n)
	cov.estimator += ((float64(cov.n) - 1.0) / float64(cov.n)) * deltaX * deltaY
}

func (cov *BivariateCovariance) Append(cov2 *BivariateCovariance) {
	oldN := cov.n
	cov.n += cov2.n
	if cov.n > 0 {
		deltaX := cov2.meanX - cov.meanX
		deltaY := cov2.meanY - cov.meanY
		cov.meanX += deltaX * float64(cov2.n) / float64(cov.n)
		cov.meanY += deltaY * float64(cov2.n) / float64(cov.n)
		cov.estimator += cov2.estimator + float64(oldN)*float64(cov2.n)/float64(cov.n)*deltaX*deltaY
	}
}

func (cov *BivariateCovariance) GetN() int {
	return int(cov.n)
}

func (cov *BivariateCovariance) GetResult() float64 {
	n := cov.n
	if cov.biasCorrected {
		n = n - 1
	}
	if n <= 0 {
		return math.NaN()
	} else if n == 1 {
		return 0
	} else {
		return cov.estimator / float64(n)
	}
}

func (b *BivariateCovariance) MeanX() float64 {
	return b.meanX
}

func (b *BivariateCovariance) MeanY() float64 {
	return b.meanY
}

func (cov *BivariateCovariance) SetBiasCorrelation(bias bool) {
	cov.biasCorrected = bias
}

func (cov BivariateCovariance) MarshalBinary() ([]byte, error) {
	// A simple encoding: plain text.
	var b bytes.Buffer
	fmt.Fprintln(&b, cov.meanX, cov.meanY, cov.n, cov.estimator, cov.biasCorrected)
	return b.Bytes(), nil
}

func (cov *BivariateCovariance) UnmarshalBinary(data []byte) error {
	// A simple encoding: plain text.
	b := bytes.NewBuffer(data)
	_, err := fmt.Fscanln(b, &cov.meanX, &cov.meanY, &cov.n, &cov.estimator, &cov.biasCorrected)
	return err
}
