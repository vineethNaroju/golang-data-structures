package main

import (
	"encoding/json"
	"fmt"
)

type Graph struct {
	VertexCount int
	Next        [][]int
}

func NewGraph(VertexCount int) *Graph {
	return &Graph{
		VertexCount: VertexCount,
		Next:        make([][]int, VertexCount),
	}
}

func (g *Graph) AddEdge(u, v int) {
	g.Next[u] = append(g.Next[u], v)
	g.Next[v] = append(g.Next[v], u)
}

func (g *Graph) Print() {
	fmt.Println("Graph with vertex count", g.VertexCount)

	for idx, arr := range g.Next {
		fmt.Printf("Next[%d]: ", idx)

		for _, val := range arr {
			fmt.Printf("%d ", val)
		}

		fmt.Println()
	}
}

func (g *Graph) ToString() string {
	str, err := json.Marshal(g)

	if err != nil {
		return err.Error()
	}

	return string(str)
}

func main() {
	graph := NewGraph(3)

	graph.AddEdge(0, 1)
	graph.AddEdge(0, 2)
	graph.AddEdge(2, 1)

	graph.Print()

	fmt.Println(graph.ToString())
}
