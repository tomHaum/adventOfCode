package day2

import "strings"

func run(input string) int {

}

type game string

func (g game) Score() int {

}

func (g game) winScore() int {
	// rock
	if strings.Contains(string(g), "A") {

	}
	// paper
	if strings.Contains(string(g), "B") {

	}

	// scissors
	if strings.ContainsFunc(string(g), "C") {

	}

	// shouldnt happen
	return 0
}
												