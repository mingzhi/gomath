package hclust

import (
	"testing"
)

func TestCompleteLinkage(t *testing.T) {
	distanceMatrix := [][]float64{
		{0, 8.408169, 7.185479, 6.622472, 9.008004},
		{8.408169, 0, 14.097252, 11.136869, 12.764842},
		{7.185479, 14.097252, 0, 8.643471, 11.994746},
		{6.622472, 11.136869, 8.643471, 0, 10.388818},
		{9.008004, 12.764842, 11.994746, 10.388818, 0},
	}
	linkage := NewLinkageMethod(CompleteLinkage)
	h := NewHClust(distanceMatrix, linkage)
	builder := NewClusterMatrixBuilder(len(distanceMatrix))
	h.Cluster(builder)
	partitions := [][]int{{0, 1, 2, 3, 4}, {0, 1, 2, 0, 4}, {0, 1, 0, 0, 4}, {0, 1, 0, 0, 0}, {0, 0, 0, 0, 0}}
	for k := 0; k < len(partitions); k++ {
		for i := 0; i < len(partitions); i++ {
			if builder.Clusters()[k][i] != partitions[k][i] {
				t.Errorf("%v\n", builder.Clusters())
			}
		}

	}
}
