package main

import (
	"adventOfCode2022/day3/data"
	"fmt"
)

func main() {
	fmt.Println("starting day 2")
	var myData = data.GetData("data/input.txt")
	totalPriority := 0
	badgePriorities := 0
	for i, sack := range myData {
		for item, _ := range sack.Duplicates {
			p := data.GetPriority(item)
			totalPriority += p
		}

		if (i+1)%3 == 0 {
			p := data.GetPriority(myData[i].Badge)
			badgePriorities += p
		}
	}
	fmt.Printf("Part 1 | Signal a: %v\n", totalPriority)
	fmt.Printf("Part 2 | Signal b: %v\n", badgePriorities)
}
