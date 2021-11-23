package main

import (
	g "AirlineManagementSystem/graph"
	"fmt"
	"strconv"
)

//functions used for each option, can be modified according to the ops in graph
func airportInformation(flights g.Graph) {
	fmt.Println("\nPlease enter the airport name[3 letter code]:")
	var name string
	fmt.Scanln(&name)
	node := flights.FindNode(name)
	fmt.Println("Airport Info:")
	fmt.Println(node)
}

func addAirport(flights g.Graph) {
	fmt.Println("\nPlease enter the airport name:")
	var name string
	fmt.Scanln(&name)
	flights.AddNode(name)
	fmt.Println("done.")
	fmt.Println()
}

func removeAirport(flights g.Graph) {
	fmt.Println("\nPlease enter the airport name:")
	var name string
	fmt.Scanln(&name)
	flights.RemoveNode(name)
}

func addFlight(flights *g.Graph) {
	fmt.Println("\nPlease enter the name of the airport you are comming from :")
	var from string
	fmt.Scanln(&from)

	fmt.Println("\nPlease enter the name of the airport you are going to:")
	var to string
	fmt.Scanln(&to)

	fmt.Println("\nPlease enter the price of the flight")
	var price int
	fmt.Scanln(&price)

	flights.AddEdge(from, to, price)
	fmt.Println("done.")
	fmt.Println()
}

func removeFlight(flights *g.Graph) {
	fmt.Println("\nPlease enter the name of the airport you are comming from:")
	var from string
	fmt.Scanln(&from)
	fmt.Println("\nPlease enter the name of the airport you are going to:")
	var to string
	fmt.Scanln(&to)
	fmt.Println("\nPlease enter the price of the flight:")
	var price int
	fmt.Scanf("%d", &price)

	flights.RemoveEdge(from, to, price)
	fmt.Println("done.")
	fmt.Println()
}

func orderToVisit(flights *g.Graph) {
	fmt.Println("\nPlease enter the name of the airport you are comming from:")
	var from string
	fmt.Scanln(&from)
	fmt.Println("\nPlease enter the name of the airport you are going to:")
	var to string
	fmt.Scanln(&to)

	fmt.Println(flights.DFS(from, to))
	fmt.Println()

}

func getFlightsFrom(flights *g.Graph) {
	fmt.Println("\nPlease enter the airport name:")
	var name string
	fmt.Scanln(&name)
	fmt.Println()

	list := flights.GetOutEdges(name)
	for i := range list {
		fmt.Println("Flight from " + list[i].StartNode + "\nFlight to " + list[i].EndNode + "\nPrice: " + strconv.Itoa(list[i].Weight) + "\n\n")
	}
}

func getFlightsTo(flights *g.Graph) {
	fmt.Println("\nPlease enter the airport name:")
	var name string
	fmt.Scanln(&name)
	fmt.Println()

	list := flights.GetInEdges(name)
	for i := range list {
		fmt.Println("Flight from " + list[i].StartNode + "\nFlight to " + list[i].EndNode + "\nPrice: " + strconv.Itoa(list[i].Weight) + "\n\n")
	}
}

func menu() {
	flights := g.Graph{
		Nodes: make(map[string]*g.Node),
		Edges: make([]g.Edge, 0),
	}

	var input string
	// options below can be modified if not supported1
	fmt.Println("1. Display airport information")
	fmt.Println("2. Add a airport")
	fmt.Println("3. Remove a airport")
	fmt.Println("4. Add a flight from one airport to another airport")
	fmt.Println("5. Find an order to visit all airports starting from an airport")
	fmt.Println("6. Get all Flights from an Airport")
	fmt.Println("7. Get all Flights to an Airport")
	fmt.Println("8. Remove Flight")
	fmt.Println("0. Exit")
	fmt.Scanln(&input)

	for input != "0" {
		for input != "1" && input != "2" && input != "3" && input != "4" && input != "5" && input != "6" && input != "7" && input != "8" {
			fmt.Println("Invalid input, try again")
			fmt.Scanln(&input)
		}

		switch input {
		case "1":
			airportInformation(flights)
		case "2":
			addAirport(flights)
		case "3":
			removeAirport(flights)
		case "4":
			addFlight(&flights)
		case "5":
			orderToVisit(&flights)
		case "6":
			getFlightsFrom(&flights)
		case "7":
			getFlightsTo(&flights)
		case "8":
			removeFlight(&flights)
		}

		fmt.Println("1. Display airport information")
		fmt.Println("2. Add a airport")
		fmt.Println("3. Remove a airport")
		fmt.Println("4. Add a flight from one airport to another airport")
		fmt.Println("5. Find an order to visit all airports starting from an airport")
		fmt.Println("6. Get all Flights from an Airport")
		fmt.Println("7. Get all Flights to an Airport")
		fmt.Println("8. Remove Flight")
		fmt.Println("0. Exit")
		fmt.Scanln(&input)
	}
}

func main() {
	menu()
}
