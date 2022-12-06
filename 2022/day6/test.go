package main

import (
	"adventOfCode2022/day6/data"
	"fmt"
)

func main() {
	fmt.Println("starting day 2")
	fileName := "data/input.txt"
	result1 := data.GetData1(fileName)

	fmt.Printf("part 1: %v \n", result1+1)

	result2 := data.GetData2(fileName, 14)
	fmt.Printf("part 2: %v \n", result2)
	result1Again := data.GetData2(fileName, 4)
	fmt.Printf("part 1 again: %v \n", result1Again)
}
