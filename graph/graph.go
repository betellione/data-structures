package graph

type Graph struct {
	vertices map[string][]string
}

func NewGraph() *Graph {
	return &Graph{vertices: make(map[string][]string)}
}

func (g *Graph) AddVertex(vertex string) {
	if _, exists := g.vertices[vertex]; !exists {
		g.vertices[vertex] = []string{}
	}
}

func (g *Graph) AddEdge(vertex1, vertex2 string) {
	g.vertices[vertex1] = append(g.vertices[vertex1], vertex2)
	g.vertices[vertex2] = append(g.vertices[vertex2], vertex1)
}

func (g *Graph) GetVertices() []string {
	var vertices []string
	for vertex := range g.vertices {
		vertices = append(vertices, vertex)
	}
	return vertices
}

func (g *Graph) GetNeighbors(vertex string) []string {
	return g.vertices[vertex]
}
