package graph

import (
	S "4080Stuff/stack"
	"fmt"
)

type Node struct {
	value     string
	inDegree  int
	outDegree int
}

type Edge struct {
	weight    int
	startNode string
	endNode   string
}

type Graph struct {
	Nodes map[string]*Node
	edges []Edge
}

func (g *Graph) AddNode(code string) {

	g.Nodes[code] = &Node{
		value:     code,
		inDegree:  0,
		outDegree: 0,
	}
}

func (g *Graph) findNode(code string) *Node {
	return g.Nodes[code]
}

func (g *Graph) AddEdge(n1, n2 string, w int) {

	from := g.findNode(n1)
	to := g.findNode(n2)

	edge := Edge{
		weight:    w,
		startNode: n1,
		endNode:   n2,
	}

	g.edges = append(g.edges, edge)

	from.outDegree += 1
	to.inDegree += 1

}

func (g *Graph) getOutNodes(n *Node) []*Node {
	var outNodes []*Node
	for i := range g.edges {
		if g.edges[i].startNode == n.value {
			outNodes = append(outNodes, g.findNode(g.edges[i].endNode))
		}
	}
	return outNodes
}

func (g *Graph) PrintNodes() {
	for k := range g.Nodes {
		fmt.Println(g.Nodes[k].value)
	}
}

func (g *Graph) PrintEdges() {
	for k := range g.edges {
		fmt.Println(g.edges[k])
	}
}

func (g *Graph) DFS(s, e string) []string {
	start := g.findNode(s)
	end := g.findNode(e)

	visited := make([]string, 1)
	stackArray := make([]string, 1)

	visited[0] = start.value
	stackArray[0] = start.value

	stack := S.Stack{
		Values: stackArray,
		Top:    0,
	}

	return g.dfsHelper(end, visited, stack)
}

func (g *Graph) dfsHelper(end *Node, visited []string, stack S.Stack) []string {
	current := g.findNode(stack.Values[stack.Top])
	if end.value == current.value {
		return visited
	}

	outNodes := g.getOutNodes(current)

	for adj := range outNodes {
		visitedB := false
		for i := range visited {
			if visited[i] == outNodes[adj].value {
				visitedB = true
			}
		}

		if visitedB {
			continue
		}

		stack.Push(outNodes[adj].value)
		visited = append(visited, outNodes[adj].value)
		return g.dfsHelper(end, visited, stack)
	}

	return nil
}