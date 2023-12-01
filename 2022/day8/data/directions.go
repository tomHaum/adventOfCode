package data

import (
	"bufio"
	"os"
)

type Tree int32
type Forest struct {
	Trees  [][]Tree
	Height int
	Width  int
}

// GetData1 GetData true for up false for down
func GetData1(fileName string) Forest {
	file, err := os.Open(fileName)
	Check(err)
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	trees := make([][]Tree, 0)

	for scanner.Scan() {
		line := scanner.Text()
		row := make([]Tree, 0)
		for _, c := range line {
			t := c - '0'
			row = append(row, Tree(t))
		}
		trees = append(trees, row)
	}

	return Forest{trees, len(trees), len(trees[0])}
}

func Check(e error) {
	if e != nil {
		panic(e)
	}
}
