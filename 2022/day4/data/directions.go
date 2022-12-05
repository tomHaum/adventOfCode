package data

import (
	"bufio"
	"os"
	"regexp"
	"strconv"
)

type Range struct {
	LowerBound int
	UpperBound int
}
type Pair struct {
	LowerRange Range
	UpperRange Range
}

// GetData true for up false for down
func GetData(fileName string) []Pair {
	file, err := os.Open(fileName)
	Check(err)
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	pairRegex := regexp.MustCompile(`(\d+)-(\d+),(\d+)-(\d+)`)
	pairs := make([]Pair, 0)

	for scanner.Scan() {
		line := scanner.Text()

		matches := pairRegex.FindStringSubmatch(line)

		if len(matches) != 5 {
			continue
		}

		lower1, _ := strconv.Atoi(matches[1])
		upper1, _ := strconv.Atoi(matches[2])
		lower2, _ := strconv.Atoi(matches[3])
		upper2, _ := strconv.Atoi(matches[4])

		range1 := Range{lower1, upper1}
		range2 := Range{lower2, upper2}

		if range1.LowerBound > range2.LowerBound {
			temp := range1
			range1 = range2
			range2 = temp
		}

		pair := Pair{range1, range2}
		pairs = append(pairs, pair)
	}

	return pairs
}

func Check(e error) {
	if e != nil {
		panic(e)
	}
}
