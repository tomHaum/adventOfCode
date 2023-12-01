package data

import (
	"bufio"
	"os"
	"regexp"
	"strconv"
)

type Edge struct {
	Origin      City
	Destination City
	Length      int
}

type Route []City
type City string
type Cities map[City][]Edge

type MyMap struct {
	Edges        []Edge
	Cities       Cities
	DefaultRoute Route
}

// GetData true for up false for down
func GetData(fileName string) MyMap {
	file, err := os.Open(fileName)
	Check(err)
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	regex := regexp.MustCompile(`^(\w+) to (\w+) = (\d+)$`)

	edges := make([]Edge, 0)
	cities := make(Cities)
	sliceOfCities := make(Route, 0)
	myMap := MyMap{edges, cities, sliceOfCities}

	for scanner.Scan() {
		line := scanner.Text()
		matches := regex.FindStringSubmatch(line)

		if len(matches) != 4 {
			continue
		}

		origin := matches[1]
		destination := matches[2]

		length, err := strconv.Atoi(matches[3])
		Check(err)

		edge := Edge{City(origin), City(destination), length}
		addEdgeToMap(&myMap, edge)
	}

	for k, _ := range myMap.Cities {
		myMap.DefaultRoute = append(myMap.DefaultRoute, k)
	}

	return myMap
}

func addEdgeToMap(myMap *MyMap, edge Edge) {
	reverseEdge := Edge{edge.Destination, edge.Origin, edge.Length}

	myMap.Edges = append(myMap.Edges, edge, reverseEdge)

	addEdgeToCities(myMap, edge)
	addEdgeToCities(myMap, reverseEdge)
}

func addEdgeToCities(myMap *MyMap, edge Edge) {
	localMap := *myMap

	var connectedCities []Edge
	var ok bool
	if connectedCities, ok = localMap.Cities[edge.Origin]; !ok {
		connectedCities = make([]Edge, 0)

	}
	connectedCities = append(connectedCities, edge)
	localMap.Cities[edge.Origin] = connectedCities
}

func Check(e error) {
	if e != nil {
		panic(e)
	}
}
