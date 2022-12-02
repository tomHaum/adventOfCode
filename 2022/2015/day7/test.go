package main

import (
	"adventOfCode2015/day7/data"
	"fmt"
	"time"
)

func main() {
	fmt.Println("starting day 7")
	preParse := time.Now()
	var myData = data.GetData("data/input.txt")

	startTime := time.Now()

	var gateA = myData.Gates["a"]
	var signalA = gateA.Evaluate(&myData)
	//test(&myData)
	intermediateTime := time.Now()
	fmt.Printf("Part 1 | Signal a: %v\n", signalA)
	//clear the cache
	myData.GateCache = make(map[data.Wire]data.Signal)
	//enforce b to value of a
	myData.GateCache["b"] = signalA
	signalA = gateA.Evaluate(&myData)
	fmt.Printf("Part 2 | Nice String: %v\n", signalA)
	endTime := time.Now()

	fmt.Printf("time taken | Part 1: %v | Part 2: %v | Parsing", intermediateTime.Sub(startTime).Microseconds(), endTime.Sub(intermediateTime).Microseconds(), startTime.Sub(preParse).Microseconds())
}

func test(circuit *data.Circuit) {
	printWire(circuit, "d")
	printWire(circuit, "e")
	printWire(circuit, "f")
	printWire(circuit, "g")
	printWire(circuit, "h")
	printWire(circuit, "i")
	printWire(circuit, "x")
	printWire(circuit, "y")
}

func printWire(circuit *data.Circuit, wire data.Wire) {
	fmt.Printf("%v: %v\n", wire, (*circuit).Gates[wire].Evaluate(circuit))
}
