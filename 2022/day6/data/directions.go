package data

import (
	"bufio"
	"os"
)

// GetData true for up false for down
func GetData1(fileName string) int {
	file, err := os.Open(fileName)
	Check(err)
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanRunes)

	char1 := '-'
	char2 := '-'
	char3 := '-'
	i := 0
	for scanner.Scan() {
		line := scanner.Text()
		char4 := int32(line[0])

		if i > 3 && char1 != char2 && char2 != char3 && char3 != char4 && char1 != char3 && char1 != char4 && char2 != char4 {
			return i
		}
		i++
		char1 = char2
		char2 = char3
		char3 = char4
	}
	return 0
}

func GetData2(fileName string, headerLength int) int {
	file, err := os.Open(fileName)
	Check(err)
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	scanner.Scan()
	line := scanner.Text()
	lineLength := len(line)

	for i := 0; i < lineLength; i++ {
		if isUnique(line[i : i+headerLength]) {
			return i + headerLength
		}
	}
	return 0
}

func isUnique(str string) bool {
	visited := make([]bool, 26)

	for _, c := range str {
		index := c - 'a'
		if visited[index] {
			return false
		}
		visited[index] = true
	}

	return true
}
func Check(e error) {
	if e != nil {
		panic(e)
	}
}
