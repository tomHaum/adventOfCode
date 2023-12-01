package main

import (
	"adventOfCode2015/day3/data"
	"fmt"
)

func main() {
	fmt.Println("starting day 1")
	dimensions := data.GetData("data/input.txt")
	wrappingPaper := 0
	ribbon := 0
	for _, box := range dimensions {
		wrappingPaper += calculateWrappingPaper(box)
		ribbon += calculateRibbon(box)
	}
	fmt.Printf("Part 1 | Total Wrapping Paper: %v\n", wrappingPaper)
	fmt.Printf("Part 2 | Total Ribbon: %v\n", ribbon)
}

func calculateWrappingPaper(box data.Dimension) int {
	side1 := box.Width * box.Height
	side2 := box.Height * box.Length
	side3 := box.Width * box.Length

	minSide := min3(side1, side2, side3)

	return (side1 + side1) + (side2 + side2) + (side3 + side3) + minSide
}

func calculateRibbon(box data.Dimension) int {
	side1 := (box.Width * 2) + (box.Height * 2)
	side2 := (box.Height * 2) + (box.Length * 2)
	side3 := (box.Width * 2) + (box.Length * 2)

	return min3(side1, side2, side3) + (box.Height * box.Width * box.Length)
}

func min2(a int, b int) int {
	if a < b {
		return a
	}
	return b
}
func min3(a int, b int, c int) int {
	ab := min2(a, b)
	return min2(ab, c)
}
