package data

import (
	"bufio"
	"os"
	"strconv"
)

func GetData(fileName string) [][]int {
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	elves := make([][]int, 0)
	elf := make([]int, 1)
	for scanner.Scan() {
		line := scanner.Text()

		if len(line) == 0 {
			elves = append(elves, elf)
			elf = make([]int, 0)
			continue
		}

		calorie, err := strconv.Atoi(line)
		Check(err)

		elf = append(elf, calorie)
	}

	return elves
}

func Check(e error) {
	if e != nil {
		panic(e)
	}
}
