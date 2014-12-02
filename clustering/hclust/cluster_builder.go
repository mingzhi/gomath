package hclust

type ClusterBuilder interface {
	Merge(i, j int, distance float64)
}
