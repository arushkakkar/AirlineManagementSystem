package graph

import (
	S "AirlineManagementSystem/stack"
	"fmt"
)

type Node struct {
	Value     string
	InDegree  int
	OutDegree int
}

type Edge struct {
	Weight    int
	StartNode string
	EndNode   string
}

type Graph struct {
	Nodes map[string]*Node
	Edges []Edge
}

func (g *Graph) AddNode(code string) {

	g.Nodes[code] = &Node{
		Value:     code,
		InDegree:  0,
		OutDegree: 0,
	}
}

func (g *Graph) FindNode(code string) *Node {
	return g.Nodes[code]
}

func (g *Graph) AddEdge(n1, n2 string, w int) {

	from := g.FindNode(n1)
	to := g.FindNode(n2)

	edge := Edge{
		Weight:    w,
		StartNode: n1,
		EndNode:   n2,
	}

	g.Edges = append(g.Edges, edge)

	from.OutDegree += 1
	to.InDegree += 1
}

func (g *Graph) RemoveNode(node string) {
	for i := 0; i < len(g.Edges); i++ {
		if g.Edges[i].StartNode == node || g.Edges[i].EndNode == node {
			g.Edges = append(g.Edges[:i], g.Edges[i+1:]...)
			i -= 1
		}
	}
	delete(g.Nodes, node)
}

func (g *Graph) RemoveEdge(start, end string, weight int) {
	for i := range g.Edges {
		if g.Edges[i].StartNode == start && g.Edges[i].EndNode == end && g.Edges[i].Weight == weight {
			g.Edges = append(g.Edges[:i], g.Edges[i+1:]...)
		}
	}
}

func (g *Graph) getOutNodes(n *Node) []*Node {
	var outNodes []*Node
	for i := range g.Edges {
		if g.Edges[i].StartNode == n.Value {
			outNodes = append(outNodes, g.FindNode(g.Edges[i].EndNode))
		}
	}
	return outNodes
}

func (g *Graph) getInNodes(n *Node) []*Node {
	var inNodes []*Node
	for i := range g.Edges {
		if g.Edges[i].EndNode == n.Value {
			inNodes = append(inNodes, g.FindNode(g.Edges[i].StartNode))
		}
	}
	return inNodes
}

func (g *Graph) GetOutEdges(s string) []Edge {
	var outEdges []Edge
	for i := range g.Edges {
		if g.Edges[i].StartNode == s {
			outEdges = append(outEdges, g.Edges[i])
		}
	}
	return outEdges
}

func (g *Graph) GetInEdges(s string) []Edge {
	var inEdges []Edge
	for i := range g.Edges {
		if g.Edges[i].EndNode == s {
			inEdges = append(inEdges, g.Edges[i])
		}
	}
	return inEdges
}

func (g *Graph) PrintNodes() {
	for k := range g.Nodes {
		fmt.Println(g.Nodes[k].Value)
	}
}

func (g *Graph) PrintEdges() {
	for k := range g.Edges {
		fmt.Println(g.Edges[k])
	}
}

func (g *Graph) DFS(s, e string) []string {
	start := g.FindNode(s)
	end := g.FindNode(e)

	visited := make([]string, 1)
	stackArray := make([]string, 1)

	visited[0] = start.Value
	stackArray[0] = start.Value

	stack := S.Stack{
		Values: stackArray,
		Top:    0,
	}

	return g.dfsHelper(end, visited, stack)
}

func (g *Graph) dfsHelper(end *Node, visited []string, stack S.Stack) []string {
	current := g.FindNode(stack.Values[stack.Top])
	if end.Value == current.Value {
		return visited
	}

	outNodes := g.getOutNodes(current)

	for adj := range outNodes {
		visitedB := false
		for i := range visited {
			if visited[i] == outNodes[adj].Value {
				visitedB = true
			}
		}

		if visitedB {
			continue
		}

		stack.Push(outNodes[adj].Value)
		visited = append(visited, outNodes[adj].Value)
		return g.dfsHelper(end, visited, stack)
	}

	return visited
}
