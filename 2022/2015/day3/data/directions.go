package data

import (
	"bufio"
	"os"
)

type Direction byte

const (
	North Direction = 0
	South           = 1
	East            = 2
	West            = 3
)

const (
	north string = "^"
	south        = "v"
	east         = ">"
	west         = "<"
)

// GetData true for up false for down
func GetData(fileName string) []Direction {
	file, err := os.Open(fileName)
	Check(err)
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanRunes)

	directions := make([]Direction, 0)

	for scanner.Scan() {
		d := North
		character := scanner.Text()

		switch character {
		case north:
			d = North
		case south:
			d = South
		case east:
			d = East
		case west:
			d = West
		default:
			panic("unknown input")
		}

		directions = append(directions, d)
	}

	return directions
}

func Check(e error) {
	if e != nil {
		panic(e)
	}
}
