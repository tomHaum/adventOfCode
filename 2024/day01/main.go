package main

import (
	"fmt"
	"github.com/tbhaum/aoc/2024/data"
	"slices"
)

func main() {
	left := []int{}
	right := []int{}
	for x := range data.Day1() {
		left = append(left, x.One)
		right = append(right, x.Two)
	}

	slices.Sort(left)
	slices.Sort(right)

	totalDiff := 0
	for i := 0; i < len(left); i++ {
		diff := left[i] - right[i]
		if diff < 0 {
			diff *= -1
		}

		totalDiff += diff
	}

	fmt.Println(totalDiff)

	totalFreq := 0
	for i := 0; i < len(left); i++ {
		c := 0
		for j := 0; j < len(right); j++ {
			//fmt.Printf("%v %v\n", left[i], right[j])
			if left[i] == right[j] {
				c++
				fmt.Println("same")
			} else if left[i] < right[j] {
				break
			}
		}

		totalFreq += c * left[i]
	}

	fmt.Println(totalFreq)

}
