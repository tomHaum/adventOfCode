package data

import (
	"bufio"
	"os"
	"regexp"
	"strconv"
)

type Crate byte

type Tower struct {
	Index  int
	Crates []Crate
}

type Instruction struct {
	Count            int
	SourceIndex      int
	DestinationIndex int
}

// GetData true for up false for down
func GetData(fileName string) (Towers []Tower, Instructions []Instruction) {
	file, err := os.Open(fileName)
	Check(err)
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	towers := make([]Tower, 0)

	// initial conditions
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 || line[1] == '1' {
			// skip to instructions
			break
		}

		if len(towers) == 0 {
			for i := 0; i < (len(line)+1)/4; i++ {
				newTower := Tower{i + 1, make([]Crate, 0)}
				towers = append(towers, newTower)
			}
		}

		for i := 0; i < len(towers); i++ {
			characterIndex := 1 + (4 * +i)
			c := line[characterIndex]
			if c == ' ' { // 1,4,7
				continue
			}

			towers[i].Crates = append(towers[i].Crates, Crate(c))
		}
	}

	// reverse towers
	for _, tower := range towers {
		ReverseCrates(tower.Crates)
	}
	instructions := make([]Instruction, 0)
	instructionRegex := regexp.MustCompile(`move (\d+) from (\d+) to (\d+)`)

	// instructions
	for scanner.Scan() {
		line := scanner.Text()
		matches := instructionRegex.FindStringSubmatch(line)

		if len(matches) != 4 {
			continue
		}

		count, err1 := strconv.Atoi(matches[1])
		Check(err1)
		origin, err2 := strconv.Atoi(matches[2])
		Check(err2)
		destination, err3 := strconv.Atoi(matches[3])
		Check(err3)

		instructions = append(instructions, Instruction{count, origin, destination})
	}

	return towers, instructions
}

func ReverseCrates(crates []Crate) {
	for i, j := 0, len(crates)-1; i < j; i, j = i+1, j-1 {
		crates[i], crates[j] = crates[j], crates[i]
	}
}

func (tower *Tower) Peek() Crate {
	return tower.Crates[len((*tower).Crates)-1]
}

func (tower *Tower) Push(crate Crate) {
	(*tower).Crates = append((*tower).Crates, crate)
}
func (tower *Tower) Pop() (crate Crate, success bool) {
	if len(tower.Crates) == 0 {
		return Crate(0), false
	}
	lastIndex := len((*tower).Crates) - 1
	ret := (*tower).Crates[lastIndex]
	(*tower).Crates = (*tower).Crates[:lastIndex]
	return ret, true
}

func Check(e error) {
	if e != nil {
		panic(e)
	}
}
