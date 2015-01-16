package stat

type Calculator interface {
	GetN() int
	GetResult() float64
}

type Appendable interface {
	Append(Calculator)
}
