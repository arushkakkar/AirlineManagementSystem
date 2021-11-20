package main

import (
	g "AirlineManagementSystem/graph"
	"fmt"
)

func asdfa() {
	roster := g.Graph{
		Nodes: make(map[string]*g.Node),
	}
	roster.AddNode("LAX")
	roster.AddNode("JFK")
	roster.AddNode("ONT")
	roster.AddNode("IND")

	roster.AddEdge("LAX", "JFK", 390)
	roster.AddEdge("JFK", "LAX", 500)
	roster.AddEdge("JFK", "ONT", 450)
	roster.AddEdge("LAX", "ONT", 100)
	roster.AddEdge("ONT", "IND", 100)

	fmt.Println(roster.DFS("LAX", "IND"))
}
