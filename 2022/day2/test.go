package main

import (
	"adventOfCode2022/day3/data"
	"fmt"
)

func main() {
	fmt.Println("starting day 2")
	var myData = data.GetData("data/input.txt")
	score1 := 0
	score2 := 0
	for _, i := range myData {
		tempScore1 := data.GetScore1(i)
		score1 += tempScore1

		tempScore2 := data.GetScore2(i)
		score2 += tempScore2
	}
	fmt.Printf("P1art 1 | Signal a: %v\n", score1)
	fmt.Printf("Part 2 | Signal b: %v\n", score2)
}
