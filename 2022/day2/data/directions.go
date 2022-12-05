package data

import (
	"bufio"
	"os"
)

type Hand int
type Instruction struct {
	Opponent Hand
	Me       Hand
}

var (
	Rock    Hand = 1
	Paper   Hand = 2
	Scissor Hand = 3
)

type Outcome int

var (
	Lose Outcome = 0
	Tie  Outcome = 3
	Win  Outcome = 6
)

// GetData true for up false for down
func GetData(fileName string) []Instruction {
	file, err := os.Open(fileName)
	Check(err)
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	instructions := make([]Instruction, 0)
	for scanner.Scan() {
		line := scanner.Text()
		me := Rock
		opponent := Rock

		switch line[0:1] {
		case "A":
			opponent = Rock
		case "B":
			opponent = Paper
		case "C":
			opponent = Scissor
		}

		switch line[2:] {
		case "X":
			me = Rock
		case "Y":
			me = Paper
		case "Z":
			me = Scissor
		}
		instructions = append(instructions, Instruction{opponent, me})
	}

	return instructions
}

func GetScore1(i Instruction) int {
	if i.Me == i.Opponent {
		return int(i.Me) + int(Tie)
	}
	if i.Me == Rock && i.Opponent == Scissor {
		return int(i.Me) + int(Win)
	}
	if i.Me == Rock && i.Opponent == Paper {
		return int(i.Me) + int(Lose)
	}

	if i.Me == Paper && i.Opponent == Rock {
		return int(i.Me) + int(Win)
	}
	if i.Me == Paper && i.Opponent == Scissor {
		return int(i.Me) + int(Lose)
	}

	if i.Me == Scissor && i.Opponent == Paper {
		return int(i.Me) + int(Win)
	}
	if i.Me == Scissor && i.Opponent == Rock {
		return int(i.Me) + int(Lose)
	}

	panic("invalid")
}

func GetScore2(i Instruction) int {
	outcome := translate(i.Me)
	if outcome == Tie {
		return int(i.Opponent) + int(Tie)
	}

	if i.Opponent == Rock && outcome == Win {
		return int(Win) + int(Paper)
	}
	if i.Opponent == Rock && outcome == Lose {
		return int(Lose) + int(Scissor)
	}

	if i.Opponent == Paper && outcome == Win {
		return int(Win) + int(Scissor)
	}
	if i.Opponent == Paper && outcome == Lose {
		return int(Lose) + int(Rock)
	}

	if i.Opponent == Scissor && outcome == Win {
		return int(Win) + int(Rock)
	}
	if i.Opponent == Scissor && outcome == Lose {
		return int(Lose) + int(Paper)
	}
	panic("invalid")
}

func translate(hand Hand) Outcome {
	switch hand {
	case Rock:
		return Lose
	case Paper:
		return Tie
	case Scissor:
		return Win
	}
	panic("invalid")
}

func Check(e error) {
	if e != nil {
		panic(e)
	}
}
