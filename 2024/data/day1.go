package data

import (
	_ "embed"
	"iter"
	"strconv"
	"strings"
)

type Tuple[One any, Two any] struct {
	One One
	Two Two
}

//go:embed day1.txt
var data01 string

func Day1() iter.Seq[Tuple[int, int]] {
	return func(yield func(Tuple[int, int]) bool) {
		for _, l := range strings.Split(data01, "\n") {
			fields := strings.Fields(l)
			left, _ := strconv.Atoi(fields[0])
			right, _ := strconv.Atoi(fields[1])
			d := Tuple[int, int]{left, right}

			if !yield(d) {
				return
			}
		}
	}
}
