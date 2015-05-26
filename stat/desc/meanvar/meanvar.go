package meanvar

import (
	"github.com/mingzhi/gomath/stat/desc"
)

type MeanVar struct {
	Mean *desc.Mean
	Var  *desc.Variance
}

func New() *MeanVar {
	var mv MeanVar
	mv.Mean = desc.NewMean()
	mv.Var = desc.NewVariance()
	return &mv
}

func (m *MeanVar) Increment(v float64) {
	m.Mean.Increment(v)
	m.Var.Increment(v)
}

func (m *MeanVar) Append(m2 *MeanVar) {
	m.Mean.Append(m2.Mean)
	m.Var.Append(m2.Var)
}
