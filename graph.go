package main

import "fmt"

type Edge struct {
	Src  string
	Dst  string
	Next *Edge
}

type EdgesList struct {
	Head *Edge
}

type Graph struct {
	VertexList map[string]*EdgesList
}

func (g *Graph) Init() {
	g.VertexList = map[string]*EdgesList{}
}

func (g *Graph) AddEdge(src, dst string) {
	edge := &Edge{Src: src, Dst: dst, Next: g.VertexList[src].Head}
	g.VertexList[src].Head = edge
}

func (g *Graph) Print() {
	for key, value := range g.VertexList {
		if value.Head != nil {
			fmt.Printf("Vertex %s connected to", key)
			for value.Head != nil {
				fmt.Println(value.Head.Next)
			}
			fmt.Println()
		}
	}
}
