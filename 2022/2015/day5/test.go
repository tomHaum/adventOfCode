package main

import (
	"adventOfCode2015/day5/data"
	"fmt"
)

type position struct {
	x int
	y int
}

func main() {
	fmt.Println("starting day 5")
	lines := data.GetData("data/input.txt")
	niceStrings1 := 0
	niceStrings2 := 0
	for _, s := range lines {
		if part1(s) {
			niceStrings1++
		}
		if part2(s) {
			niceStrings2++
		}
	}
	fmt.Printf("Part 1 | Nice String: %v\n", niceStrings1)

	fmt.Printf("Part 2 | Nice String: %v\n", niceStrings2)
}
func part1(s string) bool {
	return vowelCheck(s) && doubleAndBannedCheck(s)
}
func vowelCheck(s string) bool {
	vowelCount := 0
	for _, c := range s {
		switch c {
		case 'a', 'e', 'i', 'o', 'u':
			vowelCount++
		}
	}
	return vowelCount > 2
}

func doubleAndBannedCheck(s string) bool {
	previousCharacter := '-'
	hasDouble := false

	for _, c := range s {
		if c == previousCharacter {
			hasDouble = true
		}
		if (previousCharacter == 'a' && c == 'b') || (previousCharacter == 'c' && c == 'd') || (previousCharacter == 'p' && c == 'q') || (previousCharacter == 'x' && c == 'y') {
			return false
		}
		previousCharacter = c
	}

	return hasDouble
}

func part2(s string) bool {
	return doubleDouble(s) && spacedPair(s)
}

func doubleDouble(s string) bool {
	if len(s) < 4 {
		return false
	}
	if doubleDoubleInner(s[2:], s[0:2]) {
		return true
	}
	return doubleDouble(s[1:])
}

func doubleDoubleInner(s string, pair string) bool {
	if len(s) < 2 {
		return false
	}
	if pair[0] == s[0] && pair[1] == s[1] {
		return true
	}
	return doubleDoubleInner(s[1:], pair)
}

func spacedPair(s string) bool {
	parent := '-'
	grandParent := '-'

	for _, c := range s {
		if grandParent == c {
			return true
		}
		grandParent = parent
		parent = c
	}

	return false
}
