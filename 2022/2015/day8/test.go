package main

import (
	"adventOfCode2015/day9/data"
	"fmt"
)

func main() {
	fmt.Println("starting day 8")
	var myData = data.GetData("data/input.txt")
	totalOriginal := 0
	totalPart1 := 0
	totalPart2 := 0
	for _, line := range myData {
		lenOriginal := len(line.OriginalString)
		lenPart1 := getLenPart1(line.OriginalString)
		lenPart2 := getLenPart2(line.OriginalString)

		totalOriginal += lenOriginal
		totalPart1 += lenPart1
		totalPart2 += lenPart2
	}
	fmt.Printf("Part 1 | Signal a: %v\n", totalOriginal-totalPart1)
	fmt.Printf("Part 2 | Signal a: %v\n", totalPart2-totalOriginal)
}

func getLenPart1(line string) int {
	count := 0
	for i := 0; i < len(line); i++ {
		c := line[i]
		switch c {
		case '"':
			continue
		case '\\':
			i++ // move forward for the \
			c1 := line[i]
			switch c1 {
			case 'x':
				i += 2 // for the two hex characters
				fallthrough
			case '\\':
				fallthrough
			case '"':
				count++
			}
		default:
			count++
		}
	}

	return count
}

func getLenPart2(line string) int {
	count := 2
	for i := 0; i < len(line); i++ {
		c := line[i]
		switch c {
		case '\\':
			fallthrough
		case '"':
			count += 2
		default:
			count++
		}
	}

	return count
}
