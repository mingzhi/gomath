package hclust

type Metric interface {
	Op(x, y int) float64
}
