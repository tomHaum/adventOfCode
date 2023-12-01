package main

import (
	"adventOfCode2022/day13/data"
	"encoding/json"
	"fmt"
)

const MaxUint = ^uint(0)
const MinUint = 0
const MaxInt = int(MaxUint >> 1)
const MinInt = -MaxInt - 1

func main() {
	json.Unmarshal()

	fmt.Println("starting day 2")
	fileName := "data/input.txt"
	cave, bounds := data.GetData1(fileName, false)
	sandDropped := 0
	for ; dropSand(&cave, &bounds, data.Point{X: 500}); sandDropped++ {
		//nop
	}
	//data.Print(cave, bounds)
	fmt.Printf("part 1: %v\n", sandDropped)

	cave, bounds = data.GetData1(fileName, true)
	sandDropped = 0
	for ; dropSand(&cave, &bounds, data.Point{X: 500}); sandDropped++ {
		//nop
	}
	//data.Print(cave, bounds)
	fmt.Printf("part 2: %v\n", sandDropped)
}

func dropSand(cavePtr *data.Cave, bounds *data.Bounds, grain data.Point) bool {
	cave := *cavePtr
	for {
		if cave[grain] != data.Empty {
			return false
		}
		if grain.Y > bounds.UpperY {
			return false
		}
		// drop down
		temp := cave[data.Point{grain.X, grain.Y + 1}]
		if temp == data.Empty {
			grain.Y++
			continue
		}
		//drop left
		temp = cave[data.Point{grain.X - 1, grain.Y + 1}]
		if temp == data.Empty {
			grain.X--
			grain.Y++
			continue
		}
		//drop right
		temp = cave[data.Point{grain.X + 1, grain.Y + 1}]
		if temp == data.Empty {
			grain.X++
			grain.Y++
			continue
		}

		cave[grain] = data.Sand
		return true
	}
}
