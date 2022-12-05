package main

import (
	"adventOfCode2022/day4/data"
	"fmt"
)

func main() {
	fmt.Println("starting day 2")
	var myData = data.GetData("data/input.txt")

	containsCount := 0
	overlap := 0
	for _, p := range myData {
		if pairContains(p) {
			containsCount++
		}

		if anyOverlap(p) {
			overlap++
		}
	}

	fmt.Printf("Part 1 | Signal a: %v\n", containsCount)
	fmt.Printf("Part 2 | Signal b: %v\n", overlap)
}

func pairContains(pair data.Pair) bool {
	// first range contains second
	if pair.LowerRange.LowerBound <= pair.UpperRange.LowerBound && pair.LowerRange.UpperBound >= pair.UpperRange.UpperBound {
		return true
	}

	// second range contains first
	if pair.UpperRange.LowerBound <= pair.LowerRange.LowerBound && pair.UpperRange.UpperBound >= pair.LowerRange.UpperBound {
		return true
	}
	return false
}

func anyOverlap(pair data.Pair) bool {
	if pair.LowerRange.LowerBound <= pair.UpperRange.LowerBound && pair.LowerRange.UpperBound >= pair.UpperRange.LowerBound {
		return true
	}
	return false
}
