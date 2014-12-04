package hclust

type ClusterBuilder interface {
	Merge(i, j int, distance float64)
	Clusters() [][]int
}

/* A ClusterMatrixBuilder builds a matrix in which
 * each row represents a step in the clustering
 * and each column represents an observation or cluster.
 * In the first step (row 0), each column represents an observation.
 * In the last step, each column refers to the same cluster.
 * Each step represents a copy of the step above,
 * with two clusters merged into one.
 */
type ClusterMatrixBuilder struct {
	clusters    [][]int
	currentStep int
	distances   []float64
}

func NewClusterMatrixBuilder(n int) *ClusterMatrixBuilder {
	clusters := make([][]int, n)
	for i := 0; i < n; i++ {
		clusters[i] = make([]int, n)
		// init original step (each observation is its own cluster)
		clusters[0][i] = i
	}
	c := &ClusterMatrixBuilder{}
	c.clusters = clusters
	c.currentStep = 0
	c.distances = append(c.distances, 0)
	return c
}

func (c *ClusterMatrixBuilder) Merge(i, j int, d float64) {
	previousStep := c.currentStep
	c.currentStep++
	for k := 0; k < len(c.clusters); k++ {
		cluster := c.clusters[previousStep][k]
		if cluster == j {
			c.clusters[c.currentStep][k] = i
		} else {
			c.clusters[c.currentStep][k] = cluster
		}
	}
	c.distances = append(c.distances, d)
}

func (c *ClusterMatrixBuilder) Clusters() [][]int {
	return c.clusters
}

func (c *ClusterMatrixBuilder) Distances() []float64 {
	return c.distances
}
