package day1

import (
	"adventOfCode2022/second-pass/util"
	"cmp"
	"slices"
	"strconv"
)

func run(input string) (int, int) {
	lines := util.SplitLines(input)
	elves := make([]elf, 0)
	e := make(elf, 0)
	for _, l := range lines {
		if l == "" {
			elves = append(elves, e)
			e = make(elf, 0)
			continue
		}
		e = append(e, calorie(l))
	}
	elves = append(elves, e)

	//maxElf := slices.MaxFunc(elves, func(a, b elf) int {
	//	return cmp.Compare(a.Total(), b.Total())
	//})

	slices.SortFunc(elves, func(a, b elf) int {
		return -1 * cmp.Compare(a.Total(), b.Total())
	})

	maxElf := elves[0].Total()
	top3 := maxElf + elves[1].Total() + elves[2].Total()
	return maxElf, top3
}

type elf []calorie

func (e elf) Total() int {
	sum := 0
	for _, c := range e {
		sum += c.Value()
	}

	return sum
}

type calorie string

func (c calorie) Value() int {
	v, _ := strconv.Atoi(string(c))
	return v
}
