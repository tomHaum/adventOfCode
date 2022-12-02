package data

import (
	"bufio"
	"os"
	"regexp"
	"strconv"
)

type Command byte

var (
	Toggle Command = 0
	On     Command = 1
	Off    Command = 2
)

type Point struct {
	X int
	Y int
}

type Instruction struct {
	Command Command
	Start   Point
	End     Point
}

// GetData true for up false for down
func GetData(fileName string) []Instruction {
	file, err := os.Open(fileName)
	Check(err)
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	instructions := make([]Instruction, 0)
	validInstruction := regexp.MustCompile(`^(turn on|toggle|turn off) (\d+),(\d+) through (\d+),(\d+)$`)

	for scanner.Scan() {
		line := scanner.Text()
		matches := validInstruction.FindStringSubmatch(line)

		if len(matches) != 6 {
			continue
		}

		command := Toggle
		switch matches[1] {
		case "turn on":
			command = On
		case "toggle":
			command = Toggle
		case "turn off":
			command = Off
		default:
			panic("unknown input")
		}

		x1, err := strconv.Atoi(matches[2])
		Check(err)
		y1, err := strconv.Atoi(matches[3])
		Check(err)
		x2, err := strconv.Atoi(matches[4])
		Check(err)
		y2, err := strconv.Atoi(matches[5])
		Check(err)

		instructions = append(instructions, Instruction{command, Point{x1, y1}, Point{x2, y2}})
	}

	return instructions
}

func Check(e error) {
	if e != nil {
		panic(e)
	}
}
