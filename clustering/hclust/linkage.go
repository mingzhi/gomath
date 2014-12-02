package hclust

import (
	"math"
)

// Linage interface defines the agglomeration method to be used.
type Linakge interface {
	ComputeDissimilarity(dik, djk, dij, ci, cj, ck float64) float64
}

// LinageMethod is a wrapper for agglomeration methods.
type LinakgeMethod struct {
	Method ComputeDissimilarity
}

func NewLinkageMethod(method ComputeDissimilarity) *LinakgeMethod {
	return &LinakgeMethod{Method: method}
}

func (l *LinakgeMethod) ComputeDissimilarity(dik, djk, dij, ci, cj, ck float64) float64 {
	return l.Method(dik, djk, dij, ci, cj, ck)
}

/**
 * An LinkageMethod represents the Lance-Williams dissimilarity update formula
 * used for hierarchical agglomerative clustering.
 *
 * The general form of the Lance-Williams matrix-update formula:
 * d[(i,j),k] = ai*d[i,k] + aj*d[j,k] + b*d[i,j] + g*|d[i,k]-d[j,k]|
 *
 * Parameters ai, aj, b, and g are defined differently for different methods:
 *
 * Method          ai                   aj                   b                          g
 * -------------   ------------------   ------------------   ------------------------   -----
 * Single          0.5                  0.5                  0                          -0.5
 * Complete        0.5                  0.5                  0                          0.5
 * Average         ci/(ci+cj)           cj/(ci+cj)           0                          0
 *
 * Centroid        ci/(ci+cj)           cj/(ci+cj)           -ci*cj/((ci+cj)*(ci+cj))   0
 * Median          0.5                  0.5                  -0.25                      0
 * Ward            (ci+ck)/(ci+cj+ck)   (cj+ck)/(ci+cj+ck)   -ck/(ci+cj+ck)             0
 *
 * WeightedAverage 0.5                  0.5                  0                          0
 *
 * (ci, cj, ck are cluster cardinalities)
 * see http://www.mathworks.com/help/toolbox/stats/linkage.html
 * see http://www.stanford.edu/~maureenh/quals/html/ml/node73.html
 * see [The data analysis handbook. By Ildiko E. Frank, Roberto Todeschini. Pages 152-155]
 *
 */
type ComputeDissimilarity func(dik, djk, dij, ci, cj, ck float64) float64

func AverageLinkage(dik, djk, dij, ci, cj, ck float64) float64 {
	return (ci*dik + cj*djk) / (ci + cj)
}

func CentroidLinkage(dik, djk, dij, ci, cj, ck float64) float64 {
	return (ci*dik + cj*djk - ci*cj*dij/(ci+cj)) / (ci + cj)
}

func CompleteLinkage(dik, djk, dij, ci, cj, ck float64) float64 {
	return math.Max(dik, djk)
}

func MedianLinkage(dik, djk, dij, ci, cj, ck float64) float64 {
	return 0.5*dik + 0.5*djk - 0.25*dij
}

func SingleLinkage(dik, djk, dij, ci, cj, ck float64) float64 {
	return math.Min(dik, djk)
}

func WardLinkage(dik, djk, dij, ci, cj, ck float64) float64 {
	return ((ci+ck)*dik + (cj+ck)*djk - ck*dij) / (ci + cj + ck)
}

func WeightedAverageLinkage(dik, djk, dij, ci, cj, ck float64) float64 {
	return 0.5*dik + 0.5*djk
}
