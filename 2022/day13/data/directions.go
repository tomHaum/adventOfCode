package data

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type State byte

var (
	Empty State = 0
	Rock  State = 1
	Sand  State = 2
)

type Point struct {
	X int
	Y int
}
type Bounds struct {
	LowerX int
	LowerY int
	UpperX int
	UpperY int
}

type Cave map[Point]State

// GetData1 GetData true for up false for down
func GetData1(fileName string, addLine bool) (Cave, Bounds) {
	file, err := os.Open(fileName)
	Check(err)
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	corrdinateRegex := regexp.MustCompile(`(\d+,\d+)`)
	cave := make(Cave)
	bounds := Bounds{-1, -1, -1, -1}
	for scanner.Scan() {
		line := scanner.Text()
		matches := corrdinateRegex.FindAllString(line, -1)

		previous := parseMatch(matches[0])
		if bounds.LowerX == -1 {
			bounds = Bounds{previous.X, previous.Y, previous.X, previous.Y}
		}
		updateBounds(&bounds, previous)
		for i := 1; i < len(matches); i++ {
			current := parseMatch(matches[i])
			DrawLine(&cave, previous, current)
			previous = current
			updateBounds(&bounds, previous)
			//Print(cave, bounds)
		}
	}
	if addLine {
		DrawLine(&cave, Point{-10000, bounds.UpperY + 2}, Point{10000, bounds.UpperY + 2})
		bounds.LowerX = -10000
		bounds.UpperY = 10000
		bounds.UpperY = bounds.UpperY + 2
	}

	return cave, bounds
}
func parseMatch(match string) Point {
	comma := strings.Index(match, ",")
	x, _ := strconv.Atoi(match[:comma])
	y, _ := strconv.Atoi(match[comma+1:])
	return Point{x, y}
}
func updateBounds(bounds *Bounds, point Point) {
	bounds.LowerX = minInt(bounds.LowerX, point.X)
	bounds.LowerY = minInt(bounds.LowerY, point.Y)
	bounds.UpperX = maxInt(bounds.UpperX, point.X)
	bounds.UpperY = maxInt(bounds.UpperY, point.Y)
}

func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func DrawLine(cavePtr *Cave, a, b Point) {
	cave := *cavePtr

	if a.X == b.X {
		//vertical
		min := minInt(a.Y, b.Y)
		max := maxInt(a.Y, b.Y)

		for ; min <= max; min++ {
			cave[Point{a.X, min}] = Rock
		}
	} else {
		//horizontal
		min := minInt(a.X, b.X)
		max := maxInt(a.X, b.X)

		for ; min <= max; min++ {
			cave[Point{min, a.Y}] = Rock
		}
	}
}
func Check(e error) {
	if e != nil {
		panic(e)
	}
}

func Print(cave Cave, bounds Bounds) {
	bounds.LowerX--
	bounds.LowerY--
	bounds.LowerY--
	bounds.LowerY--
	bounds.UpperX++
	bounds.UpperY++

	point := Point{bounds.LowerX, bounds.LowerY}
	for ; point.Y <= bounds.UpperY; point.Y++ {

		for point.X = bounds.LowerX; point.X <= bounds.UpperX; point.X++ {

			switch cave[point] {
			case Empty:
				fmt.Print(".")
				break
			case Rock:
				fmt.Print("#")
				break
			case Sand:
				fmt.Print("o")
				break
			}
		}
		fmt.Println()
	}
	fmt.Println()
	//
	//for key := range cave {
	//	fmt.Println(key)
	//}
	//fmt.Println()
}
