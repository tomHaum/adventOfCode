package main

import (
	"adventOfCode2015/day9/data"
	"fmt"
)

func main() {
	fmt.Println("starting day 8")
	var myData = data.GetData("data/input.txt")
	var routes = getRoutes(myData.DefaultRoute)
	costs := make([]int, 0)
	minCost := int(^uint(0) >> 1)
	maxCost := 0
	for _, r := range routes {
		cost := routeCost(&myData, r)
		costs = append(costs, cost)

		if cost < minCost {
			minCost = cost
		}

		if cost > maxCost {
			maxCost = cost
		}
	}

	fmt.Printf("Part 1 | Shortest Route: %v\n", minCost)
	fmt.Printf("Part 2 | longest route: %v\n", maxCost)
}

func getRoutes(cities data.Route) []data.Route {
	var permRecursive func(data.Route, int)
	results := make([]data.Route, 0)

	permRecursive = func(route data.Route, n int) {
		if n == 1 {
			tmp := make(data.Route, len(route))
			copy(tmp, route)
			results = append(results, tmp)
			return
		}

		for i := 0; i < n; i++ {
			permRecursive(route, n-1)
			if n%2 == 1 {
				tmp := route[i]
				route[i] = route[n-1]
				route[n-1] = tmp
			} else {
				tmp := route[0]
				route[0] = route[n-1]
				route[n-1] = tmp
			}
		}
	}

	permRecursive(cities, len(cities))
	return results
}

func routeCost(myMap *data.MyMap, route data.Route) int {
	if len(route) < 2 {
		return 0
	}

	previousCity := route[0]
	cost := 0

	for i := 1; i < len(route); i++ {
		currentCity := route[i]
		cost += getEdgeCost(myMap, previousCity, currentCity)

		previousCity = currentCity
	}

	return cost
}

func getEdgeCost(myMap *data.MyMap, origin data.City, destination data.City) int {
	edges := (*myMap).Cities[origin]

	for i := 0; i < len(edges); i++ {
		if edges[i].Destination == destination {
			return edges[i].Length
		}
	}

	panic("could not find edge: " + destination)
}
