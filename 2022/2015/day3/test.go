package main

import (
	"adventOfCode2015/day3/data"
	"fmt"
)

type position struct {
	x int
	y int
}

func main() {
	fmt.Println("starting day 1")
	directions := data.GetData("data/input.txt")
	santaPosition := position{0, 0}

	visited := make(map[position]bool)
	visited[santaPosition] = true

	for _, direction := range directions {
		updatePosition(&santaPosition, direction)
		visited[santaPosition] = true
	}
	fmt.Printf("Part 1 | Visited Houses: %v\n", len(visited))

	santaPosition = position{0, 0}
	roboSanata := position{0, 0}
	visited = make(map[position]bool)
	for i, direction := range directions {
		var pos = &santaPosition
		if i%2 == 1 {
			pos = &roboSanata
		}
		updatePosition(pos, direction)
		visited[*pos] = true
	}
	fmt.Printf("Part 2 | Total Ribbon: %v\n", len(visited))
}

func updatePosition(p *position, direction data.Direction) {
	switch direction {
	case data.North:
		p.y++
	case data.South:
		p.y--
	case data.East:
		p.x++
	case data.West:
		p.x--
	}
}
