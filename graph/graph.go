
package main

import "fmt"

type Graph struct {
	adjList map[string][]string
}

func NewGraph() *Graph {
	return &Graph{adjList: make(map[string][]string)}
}

func (g *Graph) AddVertex(vertex string) bool {
	if _, exists := g.adjList[vertex]; !exists {
		g.adjList[vertex] = []string{}
		return true
	}
	return false
}

func (g *Graph) AddEdge(v1, v2 string) bool {
	if _, ok1 := g.adjList[v1]; ok1 {
		if _, ok2 := g.adjList[v2]; ok2 {
			g.adjList[v1] = append(g.adjList[v1], v2)
			g.adjList[v2] = append(g.adjList[v2], v1)
			return true
		}
	}
	return false
}

func (g *Graph) RemoveEdge(v1, v2 string) bool {
	if _, ok1 := g.adjList[v1]; ok1 {
		if _, ok2 := g.adjList[v2]; ok2 {
			g.adjList[v1] = remove(g.adjList[v1], v2)
			g.adjList[v2] = remove(g.adjList[v2], v1)
			return true
		}
	}
	return false
}

func (g *Graph) RemoveVertex(vertex string) bool {
	if _, exists := g.adjList[vertex]; exists {
		for _, v := range g.adjList[vertex] {
			g.adjList[v] = remove(g.adjList[v], vertex)
		}
		delete(g.adjList, vertex)
		return true
	}
	return false
}

func (g *Graph) PrintGraph() {
	for vertex, edges := range g.adjList {
		fmt.Println(vertex, ":", edges)
	}
}

func remove(slice []string, s string) []string {
	for i, val := range slice {
		if val == s {
			return append(slice[:i], slice[i+1:]...)
		}
	}
	return slice
}

func main() {
	myGraph := NewGraph()
	myGraph.AddVertex("A")
	myGraph.AddVertex("B")
	myGraph.AddVertex("C")
	myGraph.AddVertex("D")
	myGraph.AddVertex("E")

	myGraph.AddEdge("A", "B")
	myGraph.AddEdge("A", "C")
	myGraph.AddEdge("A", "D")
	myGraph.AddEdge("B", "D")
	myGraph.AddEdge("C", "D")
	myGraph.AddEdge("C", "E")
	myGraph.AddEdge("B", "E")

	myGraph.PrintGraph()
}


