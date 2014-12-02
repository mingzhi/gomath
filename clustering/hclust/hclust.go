package hclust

import (
	"math"
)

type HClust struct {
	Linkage        Linakge
	Metric         Metric
	DistanceMatrix [][]float64
}

func (h *HClust) Cluster(builder ClusterBuilder) {
	// number of obversations
	num := len(h.DistanceMatrix)
	// create stoarges
	indexUsed := make([]bool, num)
	clusterCardinalities := make([]int, num)
	for i := 0; i < num; i++ {
		indexUsed[i] = true
		clusterCardinalities[i] = 1
	}

	// perform agglomerations
	for i := 0; i < num-1; i++ {
		pair := minDistance(h.DistanceMatrix, indexUsed)
		i, j := pair.A, pair.B
		d := h.DistanceMatrix[i][j]
	}
}

func minDistance(distanceMatrix [][]float64, indexUsed []bool) Pair {
	min := math.Inf(1)
	var a, b int
	for i := 0; i < len(distanceMatrix); i++ {
		if indexUsed[i] {
			for j := 0; j < len(distanceMatrix); j++ {
				if i != j && indexUsed[j] && distanceMatrix[i][j] < min {
					min = distanceMatrix[i][j]
					a, b = i, j
				}
			}
		}
	}
	return NewPair(a, b)
}

type Pair struct {
	A, B int
}

func NewPair(a, b int) Pair {
	if a > b {
		a, b = b, a
	}
	return Pair{A: a, B: b}
}
