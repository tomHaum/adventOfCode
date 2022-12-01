package data

import (
	"bufio"
	"os"
)

// GetData true for up false for down
func GetData(fileName string) []string {
	file, err := os.Open(fileName)
	Check(err)
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	directions := make([]string, 0)

	for scanner.Scan() {
		line := scanner.Text()

		directions = append(directions, line)
	}

	return directions
}

func Check(e error) {
	if e != nil {
		panic(e)
	}
}
