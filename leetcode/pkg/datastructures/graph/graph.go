package graph

// Node represents a node in a graph.
// For simplicity, we'll assume node IDs are integers.
// Adjacency list is a common way to represent graphs.

// Graph using an adjacency list representation.
// The map keys are node IDs, and values are slices of their neighbors.
type Graph struct {
	Adj map[int][]int
}

// NewGraph creates a new graph.
func NewGraph() *Graph {
	return &Graph{Adj: make(map[int][]int)}
}

// AddEdge adds an edge to the graph.
// For an undirected graph, add edges in both directions.
func (g *Graph) AddEdge(u, v int) {
	g.Adj[u] = append(g.Adj[u], v)
	// If it's an undirected graph, uncomment the line below
	// g.Adj[v] = append(g.Adj[v], u)
}

// AddNode ensures a node exists in the graph, even if it has no edges yet.
func (g *Graph) AddNode(u int) {
	if _, ok := g.Adj[u]; !ok {
		g.Adj[u] = []int{}
	}
}

// Neighbors returns the neighbors of a node.
func (g *Graph) Neighbors(u int) []int {
	return g.Adj[u]
}

// Nodes returns all nodes in the graph.
func (g *Graph) Nodes() []int {
	var nodes []int
	for node := range g.Adj {
		nodes = append(nodes, node)
	}
	return nodes
}