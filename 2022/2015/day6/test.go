package main

import (
	"adventOfCode2015/day6/data"
	"fmt"
)

const NumberOfLights int = 1000

type Lights [NumberOfLights][NumberOfLights]bool
type BrightLights [NumberOfLights][NumberOfLights]int

func main() {
	fmt.Println("starting day 6")
	instructions := data.GetData("data/input.txt")
	var lights Lights

	for _, i := range instructions {
		action := getLightActionByCommand(i.Command)

		for x := i.Start.X; x <= i.End.X; x++ {
			for y := i.Start.Y; y <= i.End.Y; y++ {
				action(&lights, data.Point{x, y})
			}
		}
	}

	lightsOn := 0
	for i := 0; i < NumberOfLights; i++ {
		for j := 0; j < NumberOfLights; j++ {
			if lights[i][j] {
				lightsOn++
			}
		}
	}
	fmt.Printf("Part 1 | Lights On: %v\n", lightsOn)

	var brightLights BrightLights
	for _, i := range instructions {
		action := getBrightLightActionByCommand(i.Command)

		for x := i.Start.X; x <= i.End.X; x++ {
			for y := i.Start.Y; y <= i.End.Y; y++ {
				action(&brightLights, data.Point{x, y})
			}
		}
	}

	totalBrightness := 0
	for i := 0; i < NumberOfLights; i++ {
		for j := 0; j < NumberOfLights; j++ {
			totalBrightness += brightLights[i][j]
		}
	}
	fmt.Printf("Part 2 | Nice String: %v\n", totalBrightness)
}
func getLightActionByCommand(command data.Command) func(*Lights, data.Point) {
	switch command {
	case data.Off:
		return turnOffLight
	case data.Toggle:
		return toggleLight
	case data.On:
		return turnOnLight
	}
	panic("unknown command")
}

func getBrightLightActionByCommand(command data.Command) func(*BrightLights, data.Point) {
	switch command {
	case data.Off:
		return turnOffBrightLight
	case data.Toggle:
		return toggleBrightLight
	case data.On:
		return turnOnBrightLight
	}
	panic("unknown command")
}

func turnOnLight(lights *Lights, point data.Point) {
	lights[point.X][point.Y] = true
}
func turnOffLight(lights *Lights, point data.Point) {
	lights[point.X][point.Y] = false
}
func toggleLight(lights *Lights, point data.Point) {
	lights[point.X][point.Y] = !lights[point.X][point.Y]
}
func turnOnBrightLight(lights *BrightLights, point data.Point) {
	lights[point.X][point.Y]++
}
func turnOffBrightLight(lights *BrightLights, point data.Point) {
	if lights[point.X][point.Y] > 0 {
		lights[point.X][point.Y]--
	}
}
func toggleBrightLight(lights *BrightLights, point data.Point) {
	lights[point.X][point.Y] += 2
}
