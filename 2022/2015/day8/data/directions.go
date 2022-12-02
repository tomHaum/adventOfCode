package data

import (
	"bufio"
	"os"
)

type MyLine struct {
	OriginalString string
	MyString       MyString
}
type MyString []MyCharacter
type MyCharacter uint8

// GetData true for up false for down
func GetData(fileName string) []MyLine {
	file, err := os.Open(fileName)
	Check(err)
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	myLines := make([]MyLine, 0)

	for scanner.Scan() {
		line := scanner.Text()
		myString := make(MyString, 0)

		myLine := MyLine{line, myString}
		myLines = append(myLines, myLine)
	}

	return myLines
}

func Check(e error) {
	if e != nil {
		panic(e)
	}
}
