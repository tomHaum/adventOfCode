package adventofcode2023

import "strings"

func run(input []string) (int, int) {
	sum1 := 0
	sum2 := 0
	for _, line := range input {
		sum1 += calibrationValue(line)
		sum2 += calibrateValue2(line)
	}

	return sum1, sum2
}
func day1Part1(input []string) int {
	sum := 0
	for _, line := range input {
		sum += calibrationValue(line)
	}

	return sum
}

func day1part2(input []string) int {
	sum := 0
	for _, line := range input {
		v := calibrateValue2(line)
		sum += v
	}

	return sum
}

func calibrationValue(line string) int {
	firstFound := false
	firstDigit := '-'
	lastDigit := '-'
	for _, c := range line {
		if isDigit(c) {
			if !firstFound {
				firstDigit = c
				lastDigit = c
				firstFound = true
			} else {
				lastDigit = c
			}
		}
	}

	return (10 * runeToInt(firstDigit)) + runeToInt(lastDigit)
}

func calibrateValue2(line string) int {
	firstFound := false
	firstDigit := -1
	lastDigit := -1
	for i, _ := range line {
		if v := isDigitPart2(line[i:]); v != -1 {
			if !firstFound {
				firstDigit = v
				lastDigit = v
				firstFound = true
			} else {
				lastDigit = v
			}
		}
	}

	return (10 * firstDigit) + lastDigit
}

func isDigit(c rune) bool {
	return runeToInt(c) >= 0 && runeToInt(c) <= 9
}

var digits = []string{
	"one",
	"two",
	"three",
	"four",
	"five",
	"six",
	"seven",
	"eight",
	"nine",
}

func isDigitPart2(str string) int {
	if isDigit(rune(str[0])) {
		return runeToInt(rune(str[0]))
	}

	for i, d := range digits {
		if strings.HasPrefix(str, d) {
			return 1 + i
		}
	}

	return -1
}

func runeToInt(c rune) int {
	return int(c - '0')
}
