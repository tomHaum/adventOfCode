package part2

import (
	"fmt"
	"strconv"
	"strings"
)

func run(input string) int {
	cards := make(map[int]card)
	copies := make(map[int]int)

	for _, line := range strings.Split(input, "\n") {
		line = strings.TrimSpace(line)
		c := parseCard(line)
		cards[c.id] = c
		copies[c.id] = 1
	}
	fmt.Println()
	for id := 1; id <= len(cards); id++ {
		c := cards[id]
		for i := 0; i < c.matches; i++ {
			copies[id+1+i] += copies[id]
		}
	}

	sum := 0
	for id, _ := range cards {
		sum += copies[id]
	}
	return sum
}

type card struct {
	id int

	matches int
}

func parseCard(line string) card {
	c := card{}
	line = strings.TrimSpace(line)
	cardText, allNumbers, _ := strings.Cut(line, ":")
	idFields := strings.Fields(cardText)
	id, _ := strconv.Atoi(idFields[1])
	c.id = id
	winnersTxt, handTxt, _ := strings.Cut(allNumbers, "|")

	winners := make(map[int]interface{})
	for _, txt := range strings.Fields(winnersTxt) {
		x, _ := strconv.Atoi(txt)
		winners[x] = true
	}

	count := 0
	for _, txt := range strings.Fields(handTxt) {
		x, _ := strconv.Atoi(txt)
		_, ok := winners[x]
		if ok {
			count++
		}
	}

	return card{id: id, matches: count}
}
