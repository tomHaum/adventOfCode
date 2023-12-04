package main

import (
	_ "embed"
	"fmt"
	"github.com/pterm/pterm"
)

//go:embed part1
var input string

var area *pterm.AreaPrinter

func main() {
	pterm.Info.Println("Day 3")
	area, _ = pterm.DefaultArea.WithFullscreen(true).Start()
	defer area.Stop()
	g := GetGrid(input)
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

}
