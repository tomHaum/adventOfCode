package data

import (
	"bufio"
	"os"
)

// GetData true for up false for down
func GetData(fileName string) []bool {
	file, err := os.Open(fileName)
	Check(err)
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanRunes)

	floors := make([]bool, 0)

	for scanner.Scan() {
		switch rune := scanner.Text(); rune {
		case "(":
			floors = append(floors, true)
		case ")":
			floors = append(floors, false)
		default:
			panic("unexpceted character")
		}
	}

	return floors
}

func Check(e error) {
	if e != nil {
		panic(e)
	}
}
