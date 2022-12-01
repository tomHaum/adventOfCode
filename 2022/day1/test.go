package main

import (
	"adventOfCode2022/day1/data"
	"adventOfCode2022/day1/heaps"
	"container/heap"
	"fmt"
)

func main() {
	fmt.Println("starting day 1")
	elves := data.GetData("data/elfs.txt")
	maxHeap := &Heaps.IntMaxHeap{}
	heap.Init(maxHeap)

	for _, elf := range elves {
		elfCalories := 0

		for _, calorie := range elf {
			elfCalories += calorie
		}
		heap.Push(maxHeap, elfCalories)
	}

	fmt.Printf("Part 1 | Max calories: %v\n", (*maxHeap)[0])
	topThree := (*maxHeap)[0:3]
	sumOfFirstThree := topThree[0] + topThree[1] + topThree[2]
	fmt.Printf("Part 2 | Top Three Total Caolories: %v\n", sumOfFirstThree)
}
