package data

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type Dimension struct {
	Length int
	Width  int
	Height int
}

// GetData true for up false for down
func GetData(fileName string) []Dimension {
	file, err := os.Open(fileName)
	Check(err)
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	dimensions := make([]Dimension, 0)

	for scanner.Scan() {
		line := scanner.Text()

		x1 := strings.Index(line, "x")
		restOfLine := line[x1+1:]
		x2 := strings.Index(restOfLine, "x") + x1 + 1

		lengthStr := line[:x1]
		length, _ := strconv.Atoi(lengthStr)
		widthStr := line[x1+1 : x2]
		width, _ := strconv.Atoi(widthStr)
		heightStr := line[x2+1:]
		height, _ := strconv.Atoi(heightStr)

		dimensions = append(dimensions, Dimension{length, width, height})
	}

	return dimensions
}

func Check(e error) {
	if e != nil {
		panic(e)
	}
}
