// Hierachical cluster analysis on a set of dissimilarities.
package hclust

import (
	"math"
)

type HClust struct {
	DistanceMatrix [][]float64
	Linkage        Linakge
}

func NewHClust(distanceMatrix [][]float64, linkage Linakge) *HClust {
	return &HClust{
		DistanceMatrix: distanceMatrix,
		Linkage:        linkage,
	}
}

func (h *HClust) Cluster(builder ClusterBuilder) {
	// number of obversations
	num := len(h.DistanceMatrix)
	// create stoarges
	indexUsing := make([]bool, num)
	clusterCardinalities := make([]int, num)
	for i := 0; i < num; i++ {
		indexUsing[i] = true
		clusterCardinalities[i] = 1
	}

	// perform agglomerations
	for temp := 0; temp < num-1; temp++ {
		pair := minDistance(h.DistanceMatrix, indexUsing)
		i, j := pair.A, pair.B
		d := h.DistanceMatrix[i][j]
		for k := 0; k < num; k++ {
			if k != i && k != j && indexUsing[k] {
				// Update distance between i (the smaller one) and k
				dik := h.DistanceMatrix[i][k]
				djk := h.DistanceMatrix[j][k]
				dij := h.DistanceMatrix[i][j]
				ci := float64(clusterCardinalities[i])
				cj := float64(clusterCardinalities[j])
				ck := float64(clusterCardinalities[k])
				distance := h.Linkage.ComputeDissimilarity(dik, djk, dij, ci, cj, ck)
				h.DistanceMatrix[i][k] = distance
				h.DistanceMatrix[k][i] = distance
			}
		}

		// Update cluster cardinality of i (the smaller one)
		clusterCardinalities[i] += clusterCardinalities[j]

		// Erase cluster j
		indexUsing[j] = false
		for k := 0; k < num; k++ {
			h.DistanceMatrix[j][k] = math.Inf(1)
			h.DistanceMatrix[k][j] = math.Inf(1)
		}

		// Update cluster builder
		builder.Merge(i, j, d)
	}
}

func minDistance(distanceMatrix [][]float64, indexUsing []bool) Pair {
	min := math.Inf(1)
	var a, b int
	for i := 0; i < len(distanceMatrix); i++ {
		if indexUsing[i] {
			for j := 0; j < len(distanceMatrix); j++ {
				if i != j && indexUsing[j] && distanceMatrix[i][j] < min {
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
