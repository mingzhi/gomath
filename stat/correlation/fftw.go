package correlation

import (
	"github.com/mingzhi/fftw"
	"math/cmplx"
)

type FFTW struct {
	foward   fftw.DftR2C1DPlan
	backward fftw.DftC2R1DPlan
}

func NewFFTW(n int, circular bool) FFTW {
	// zero padding.
	ftlength := n
	if !circular {
		ftlength = n * 2
	}

	flag := fftw.Measure

	var f FFTW
	f.foward = fftw.NewDftR2C1DPlan(ftlength, flag)
	f.backward = fftw.NewDftC2R1DPlan(ftlength, flag)

	return f
}

func (f *FFTW) Close() {
	f.foward.Close()
	f.backward.Close()
}

func (f FFTW) XCorr(x1, x2 []float64) []float64 {
	var v1, v2, temp []complex128
	var rs []float64

	v1 = f.foward.ExecuteNewArray(x2)
	v2 = f.foward.ExecuteNewArray(x1)

	for i := 0; i < len(v1); i++ {
		v := v1[i] * cmplx.Conj(v2[i])
		temp = append(temp, v)
	}

	rs = f.backward.ExecuteNewArray(temp)
	totl := len(f.backward.Real)

	res := []float64{}
	for i := 0; i < len(x1); i++ {
		res = append(res, rs[i]/float64(totl))
	}
	return res
}

func (f FFTW) AutoCorr(x []float64) []float64 {
	return f.XCorr(x, x)
}
