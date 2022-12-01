package main

import (
	"adventOfCode2015/day1/data"
	"fmt"
)

func main() {
	fmt.Println("starting day 1")
	floors := data.GetData("data/input.txt")
	floorNumber := 0
	basement := 0
	for i, floor := range floors {
		if floor {
			floorNumber++
		} else {
			floorNumber--
		}
		if floorNumber == -1 && basement == 0 {
			basement = i + 1
		}
	}
	fmt.Printf("Part 1 | Ending Floor: %v\n", floorNumber)
	fmt.Printf("Part 2 | Found Basement: %v\n", basement)
}
