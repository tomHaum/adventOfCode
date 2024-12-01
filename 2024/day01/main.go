package main

import (
	"slices"
	"strconv"
	"strings"
)

func parse(data string) ([]int, []int) {
	left := []int{}
	right := []int{}

	for _, l := range strings.Split(data, "\n") {
		fields := strings.Fields(l)
		x, _ := strconv.Atoi(fields[0])
		y, _ := strconv.Atoi(fields[1])

		left = append(left, x)
		right = append(right, y)
	}

	return left, right
}

func Day1(left, right []int) (int, int) {
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

	totalFreq := 0
	for i := 0; i < len(left); i++ {
		c := 0
		for j := 0; j < len(right); j++ {
			if left[i] == right[j] {
				c++
			} else if left[i] < right[j] {
				break
			}
		}

		totalFreq += c * left[i]
	}

	return totalDiff, totalFreq
}
