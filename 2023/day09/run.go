package day09

import (
	"slices"
	"strconv"
	"strings"
)

func run(input string) (int, int) {
	input = strings.ReplaceAll(input, "\r\n", "\n")
	split := strings.Split(input, "\n")

	split = split[0:]
	sumPart1 := 0
	sumPart2 := 0
	for _, l := range split {
		numsTxt := strings.Fields(l)
		nums := stringsToInts(numsTxt)
		part1, part2 := Extrapolate(nums)
		sumPart1 += part1
		sumPart2 += part2
	}

	return sumPart1, sumPart2
}

func Extrapolate(values []int) (int, int) {
	newValues := slices.Clone(values)
	lastNums := make([]int, 0)
	firstNums := make([]int, 0)
outer:
	for {
		lastNums = append(lastNums, newValues[len(newValues)-1])
		firstNums = append(firstNums, newValues[0])
		for i := 0; i < len(newValues)-1; i++ {
			newValues[i] = newValues[i+1] - newValues[i]
		}
		newValues = newValues[:len(newValues)-1]
		for _, v := range newValues {
			if v != 0 {
				continue outer
			}
		}
		break
	}

	sumPart1 := 0
	for _, n := range lastNums {
		sumPart1 += n
	}

	part2 := 0
	for i, _ := range firstNums {
		part2 = firstNums[len(firstNums)-i-1] - part2
	}
	return sumPart1, part2
}

func stringsToInts(strs []string) []int {
	ints := make([]int, len(strs))
	for i, s := range strs {
		ints[i], _ = strconv.Atoi(s)
	}

	return ints
}
