package main

import (
	"adventOfCode2022/day8/data"
	"fmt"
)

const MaxUint = ^uint(0)
const MinUint = 0
const MaxInt = int(MaxUint >> 1)
const MinInt = -MaxInt - 1

func main() {
	fmt.Println("starting day 2")
	fileName := "data/input.txt"
	forest := data.GetData1(fileName)
	visited := make([][]bool, forest.Height)
	for i := 0; i < forest.Height; i++ {
		row := make([]bool, forest.Width)
		for j := 0; j < forest.Width; j++ {
			row[j] = false
		}
		visited[i] = row
	}

	// top down
	for x := 0; x < forest.Width; x++ {
		maxHeight := forest.Trees[0][x]
		for y := 0; y < forest.Height; y++ {
			if maxHeight == 9 {
				break
			}
			if y == 0 {
				visited[y][x] = true
				continue
			}
			currentTree := forest.Trees[y][x]
			if maxHeight < currentTree {
				visited[y][x] = true
			}
			maxHeight = maxTree(maxHeight, currentTree)
		}
	}
	// bottom up
	for x := 0; x < forest.Width; x++ {
		maxHeight := forest.Trees[forest.Height-1][x]
		for y := forest.Height - 1; y >= 0; y-- {
			if maxHeight == 9 {
				break
			}
			if y == forest.Height-1 {
				visited[y][x] = true
				continue
			}
			currentTree := forest.Trees[y][x]
			if maxHeight < currentTree {
				visited[y][x] = true
			}
			maxHeight = maxTree(maxHeight, currentTree)
		}
	}
	// left to right
	for y := 0; y < forest.Height; y++ {
		maxHeight := forest.Trees[y][0]
		for x := 0; x < forest.Width; x++ {
			if maxHeight == 9 {
				break
			}
			if x == 0 {
				visited[y][x] = true
				continue
			}
			currentTree := forest.Trees[y][x]
			if maxHeight < currentTree {
				visited[y][x] = true
			}
			maxHeight = maxTree(maxHeight, currentTree)
		}
	}
	// right to left
	for y := 0; y < forest.Height; y++ {
		maxHeight := forest.Trees[y][forest.Width-1]
		for x := forest.Width - 1; x >= 0; x-- {
			if maxHeight == 9 {
				break
			}
			if x == forest.Width-1 {
				visited[y][x] = true
				continue
			}
			currentTree := forest.Trees[y][x]
			if maxHeight < currentTree {
				visited[y][x] = true
			}
			maxHeight = maxTree(maxHeight, currentTree)
		}
	}

	// count
	countVisible := 0
	highestScore := 0
	for y := 0; y < forest.Height; y++ {
		for x := 0; x < forest.Width; x++ {
			if visited[y][x] {
				countVisible++
			}
			var sceneicScore = scenicScore(&forest, x, y)
			//fmt.Printf("%v", sceneicScore)
			highestScore = maxInt(highestScore, sceneicScore)
		}
		//fmt.Println()
	}
	scenicScore(&forest, 3, 2)
	fmt.Printf("part 1: %v\n", countVisible)
	fmt.Printf("part 2: %v\n", highestScore)
}
func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func maxTree(a, b data.Tree) data.Tree {
	if a > b {
		return a
	}
	return b
}

func scenicScore(forest *data.Forest, x, y int) int {
	// edges
	if x == 0 || y == 0 || x == forest.Width-1 || y == forest.Height-1 {
		return 0
	}
	originalTree := forest.Trees[y][x]
	score := 1

	// up
	tempScore := 0
	for i := y - 1; i >= 0; i-- {
		tempScore++
		currentTree := forest.Trees[i][x]
		if currentTree >= originalTree {
			break
		}
	}
	score *= tempScore

	// down
	tempScore = 0
	for i := y + 1; i < forest.Height; i++ {
		tempScore++
		currentTree := forest.Trees[i][x]
		if currentTree >= originalTree {
			break
		}
	}
	score *= tempScore

	// left
	tempScore = 0
	for i := x - 1; i >= 0; i-- {
		tempScore++
		currentTree := forest.Trees[y][i]
		if currentTree >= originalTree {
			break
		}
	}
	score *= tempScore

	// right
	tempScore = 0
	for i := x + 1; i < forest.Width; i++ {
		tempScore++
		currentTree := forest.Trees[y][i]
		if currentTree >= originalTree {
			break
		}
	}
	score *= tempScore

	return score
}
