package main

import (
	g "AirlineManagementSystem/graph"
	"fmt"
)

//functions used for each option, can be modified according to the ops in graph
func airPortInformation(flights g.Graph) {
	//Ask for airport name (three letter code like "LAX" and use g.FindNode(name) to find airport and display its info)

}

func addAirPort(flights g.Graph) {
	//Ask for airport name and use flights.AddNode(name) to add airport

}

func removeAirPort(flights g.Graph) {
	//Ask for airport name and use flights.RemoveNode(name) to remove airport

}

func addFlight(flights g.Graph) {
	//Ask for from and to airports and price of flight (integer) and use flights.AddEdge(from, to, price) to add flight

}

func removeFlight(flights g.Graph) {
	//Ask for from and to airports and price of flight (integer) and use flights.RemoveEdge(from, to, price) to add flight

}

func orderToVisit(flights g.Graph) {
	//TODO

}

func getFlightsFrom(flights g.Graph) {
	//Pls just implement asking for an airport name, i will do the rest later

}

func getFlightsTo(flights g.Graph) {
	//Pls just implement asking for an airport name, i will do the rest later
}

func menu() {
	flights := g.Graph{
		Nodes: make(map[string]*g.Node),
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
	fmt.Println("0. Exit")
	fmt.Scanln(&input)

	for input != "0" {
		for input != "1" && input != "2" && input != "3" && input != "4" && input != "5" {
			fmt.Println("Invalid input, try again")
			fmt.Scanln(&input)
		}

		switch input {
		case "1":
			airPortInformation(flights)
		case "2":
			addAirPort(flights)
		case "3":
			removeAirPort(flights)
		case "4":
			addFlight(flights)
		case "5":
			orderToVisit(flights)
		case "6":
			getFlightsFrom(flights)
		case "7":
			getFlightsTo(flights)
		}

		fmt.Println()
		fmt.Println("1. Display airport information")
		fmt.Println("2. Add a airport")
		fmt.Println("3. Remove a airport")
		fmt.Println("4. Add a flight from one airport to another airport")
		fmt.Println("5. Find an order to visit all airports starting from an airport")
		fmt.Println("6. Get all Flights from an Airport")
		fmt.Println("7. Get all Flights to an Airport")
		fmt.Println("0. Exit")
		fmt.Scanln(&input)
	}
}

func main() {
	menu()
}
