package main

import (
	"adventOfCode2022/day5/data"
	"fmt"
)

func main() {
	fmt.Println("starting day 2")
	fileName := "data/input.txt"
	towers, instructions := data.GetData(fileName)

	printTowers(towers)
	for _, i := range instructions {
		followInstruction(i, &towers)
		//printTowers(towers)
	}

	fmt.Printf("Part 1 | Signal a: ")
	for _, t := range towers {
		fmt.Printf("%c", t.Peek())
	}
	fmt.Println()

	towers, _ = data.GetData(fileName)
	for _, i := range instructions {
		followInstruction2(i, &towers)
		//printTowers(towers)
	}
	fmt.Printf("Part 2 | Signal a: ")
	for _, t := range towers {
		fmt.Printf("%c", t.Peek())
	}
	fmt.Println()
}

func followInstruction(instruction data.Instruction, towers *[]data.Tower) {
	for i := 0; i < instruction.Count; i++ {
		temp, _ := (*towers)[instruction.SourceIndex-1].Pop()
		(*towers)[instruction.DestinationIndex-1].Push(temp)
	}
}

func followInstruction2(instruction data.Instruction, towers *[]data.Tower) {
	crates := make([]data.Crate, 0)
	for i := 0; i < instruction.Count; i++ {
		temp, _ := (*towers)[instruction.SourceIndex-1].Pop()
		crates = append(crates, temp)
	}
	for i := 0; i < instruction.Count; i++ {
		c := crates[instruction.Count-i-1]
		(*towers)[instruction.DestinationIndex-1].Push(c)
	}
}

func printTowers(towers []data.Tower) {
	for _, t := range towers {
		fmt.Printf("%v", t.Index)
	}
	fmt.Println()

	i := 0

	for {
		printed := 0
		for _, t := range towers {
			if len(t.Crates) > i {
				fmt.Printf("%c", rune(t.Crates[i]))
				printed++
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
		if printed == 0 {
			break
		}
		i++
	}
}
