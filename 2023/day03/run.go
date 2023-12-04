package main

import (
	"fmt"
	"strings"
	"unicode"
)

func run(input string) (answer1 int, answer2 int) {
	g := GetGrid(input)

	printGrid(g)
	fmt.Sprintf("%v", g)
	numbers := g.FindAllNumbers()
	fmt.Sprintf("%v", numbers)

	partNumbers := make([]Number, 0)
	for _, n := range numbers {
		if n.IsPartNumber(g) {
			partNumbers = append(partNumbers, n)
		}
	}

	sum := 0
	for _, p := range partNumbers {
		sum += p.value
	}

	addedGears := map[cord]bool{}
	gearSum := 0
	for _, n := range partNumbers {
		if !n.gear {
			//gearSum += n.value
			continue
		}

		if addedGears[cord{n.gearX, n.gearY}] {
			continue
		}

		for _, gr := range partNumbers {
			if gr == n {
				continue
			}
			if gr.gearX == n.gearX && gr.gearY == n.gearY {
				gearSum += n.value * gr.value
				addedGears[cord{n.gearX, n.gearY}] = true
				break
			}
		}
	}
	return sum, gearSum
}

type cord struct {
	x int
	y int
}

func printGrid(g Grid) {
	for _, l := range g {
		for _, r := range l {
			fmt.Printf("%v", string(r))
		}
		fmt.Println()
	}
}

func GetGrid(input string) Grid {
	split := strings.Split(input, "\n")
	g := make(Grid, len(split))

	for y, l := range split {
		g[y] = []rune(strings.TrimSpace(l))
	}

	return g
}

type Schematic struct {
	grid Grid

	numbers []Number
}

type Grid [][]rune

func (g Grid) FindAllNumbers() []Number {
	numbers := make([]Number, 0)
	for y, l := range g {
		n := Number{}
		for x, r := range l {
			if unicode.IsNumber(r) {
				if n.value == 0 {
					n.x = x
					n.y = y
					n.value = int(r - '0')
					n.len = 1
				} else {
					n.value *= 10
					n.value += int(r - '0')
					n.len += 1
				}
			} else {
				if n.value != 0 {
					numbers = append(numbers, n)
					n = Number{}
				}
			}
		}

		if n.value != 0 {
			numbers = append(numbers, n)
		}
	}

	return numbers
}

type Number struct {
	x   int
	y   int
	len int

	value int
	gear  bool
	gearX int
	gearY int
}

var offsets = [][]int{
	{-1, -1}, {-1, 0}, {-1, 1},
	{0, -1}, {0, 0}, {0, 1},
	{1, -1}, {1, 0}, {1, 1},
}

func (n *Number) IsPartNumber(g Grid) bool {
	x := n.x
	i := 0
	for i < n.len {
		for _, o := range offsets {
			tempX := x
			tempY := n.y
			tempX += o[0]
			tempY += o[1]

			if tempX > 0 && tempX < len(g[n.y]) {
				if tempY > 0 && tempY < len(g) {
					r := g[tempY][tempX]
					if IsSymbol(r, tempX, tempY) {
						if r == '*' {
							n.gear = true
							n.gearX = tempX
							n.gearY = tempY
						}

						return true
					}
				}
			}
		}

		i++
		x++
	}

	return false
}

var symbols = map[rune]int{}

func IsSymbol(r rune, x, y int) bool {
	return !unicode.IsLetter(r) && !unicode.IsDigit(r) && r != '.'

	//symbols[r] = symbols[r] + 1
	//return true
}
