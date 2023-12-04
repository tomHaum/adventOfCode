package day02

import (
	"regexp"
	"strconv"
	"strings"
)

const MaxUint = ^uint(0)
const MaxInt = int(MaxUint >> 1)

func run(input string) (int, int) {
	isPossible := func(g game) bool {
		for _, gs := range g.sets {
			if gs.red > 12 {
				return false
			}

			if gs.green > 13 {
				return false
			}

			if gs.blue > 14 {
				return false
			}
		}

		return true
	}

	powerSet := func(g game) int {
		blueMin := 0
		redMin := 0
		greenMin := 0

		for _, gs := range g.sets {
			if gs.red > redMin {
				redMin = gs.red
			}

			if gs.green > greenMin {
				greenMin = gs.green
			}

			if gs.blue > blueMin {
				blueMin = gs.blue
			}
		}

		return blueMin * redMin * greenMin
	}

	sum := 0
	sumPower := 0
	for _, line := range strings.Split(input, "\n") {
		line = strings.TrimSpace(line)
		g := getGame(line)
		if isPossible(g) {
			sum += g.id
		}

		sumPower += powerSet(g)
	}
	return sum, sumPower
}

type game struct {
	id int

	sets []gameSet
}

type gameSet struct {
	red   int
	blue  int
	green int
}

func getGame(line string) game {
	id := getGameId(line)
	sets := getGameSets(line)

	return game{
		id:   id,
		sets: sets,
	}
}

var idRegex = regexp.MustCompile(`Game (\d+):`)

func getGameId(line string) int {
	matches := idRegex.FindStringSubmatch(line)
	x, _ := strconv.Atoi(matches[1])
	return x
}

func getGameSets(line string) []gameSet {
	_, games, _ := strings.Cut(line, ":")
	games = strings.TrimSpace(games)
	split := strings.Split(games, ";")
	sets := make([]gameSet, len(split))
	for i, g := range split {
		sets[i] = getGameSet(g)
	}

	return sets
}

func getGameSet(txt string) gameSet {
	split := strings.Split(txt, ",")
	blue := 0
	red := 0
	green := 0

	for _, str := range split {
		n, c, _ := strings.Cut(strings.TrimSpace(str), " ")
		x, _ := strconv.Atoi(n)

		switch c {
		case "blue":
			blue = x
		case "red":
			red = x
		case "green":
			green = x
		default:
			panic("unknown color")
		}
	}

	return gameSet{
		red:   red,
		blue:  blue,
		green: green,
	}
}
