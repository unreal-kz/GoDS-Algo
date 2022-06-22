package main

import "fmt"

//Graph structure
type Graph struct {
	vertices []*Vertex
}

//Vertex structure
type Vertex struct {
	key      int
	adjacent []*Vertex
}

// Add Vertex
func (g *Graph) AddVertex(k int) {
	if contain(g.vertices, k) {
		err := fmt.Errorf("Vertex %v has not been added since it already exists", k)
		fmt.Println(err.Error())
	} else {
		g.vertices = append(g.vertices, &Vertex{key: k})
	}
}

// Add Edge
func (g *Graph) AddEdge(from, to int) {
	//get vertex
	fromVertex := g.getVertex(from)
	toVertex := g.getVertex(to)
	//check error
	if fromVertex == nil || toVertex == nil {
		err := fmt.Errorf("Invalid edge (%v-->%v)", from, to)
		fmt.Println(err.Error())
	} else if contain(fromVertex.adjacent, to) {
		err := fmt.Errorf("Existing edge (%v-->%v)", from, to)
		fmt.Println(err.Error())
	} else {
		//add egde
		fromVertex.adjacent = append(fromVertex.adjacent, toVertex)
	}
}

//getVertex
func (g *Graph) getVertex(k int) *Vertex {
	for i, v := range g.vertices {
		if v.key == k {
			return g.vertices[i]
		}
	}
	return nil
}

//contains
func contain(s []*Vertex, k int) bool {
	for _, v := range s {
		if k == v.key {
			return true
		}
	}
	return false
}

//custom print
func (g *Graph) Print() {
	for _, v := range g.vertices {
		fmt.Printf("\nVertex %v: ", v.key)
		for _, k := range v.adjacent {
			fmt.Printf(" %v ", k.key)
		}
	}
	fmt.Println()
}
func main() {
	test := &Graph{}

	for i := 0; i < 5; i++ {
		test.AddVertex(i)
	}

	test.AddEdge(1, 2)
	test.AddEdge(1, 6)
	test.AddEdge(3, 2)

	test.Print()
}
