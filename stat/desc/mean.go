package desc

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
