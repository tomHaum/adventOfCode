package day04

import (
	"fmt"
	"strconv"
	"strings"
)

func run(input string) (int, int) {
	cards := make([]card, 0)

	for _, l := range strings.Split(input, "\n") {
		l = strings.TrimSpace(l)

		c := card{raw: l}
		c.setNumbers()
		cards = append(cards, c)
	}

	sum := 0
	copies := map[int]int{}
	for _, c := range cards {
		sum += c.Value()
		copies[c.id] = 1
	}

	totalCards := 0
	for _, c := range cards {
		totalCards += copies[c.id]
		_, count := c.ValueAndCount()
		fmt.Printf("id: %v count: %v\n", c.id, count)
		for i := 0; i < count; i++ {
			copies[c.id+1+i] = copies[c.id+1+i] + copies[c.id]
		}
	}

	return sum, totalCards
}

type nothing struct{}
type card struct {
	raw string

	id      int
	winners map[int]nothing
	drawn   []int
}

func (c *card) setNumbers() {
	c.winners = make(map[int]nothing)
	c.drawn = make([]int, 0)
	cardText, allNumbers, _ := strings.Cut(c.raw, ":")
	_, cardId, _ := strings.Cut(cardText, " ")
	id, _ := strconv.Atoi(cardId)
	c.id = id

	winners, drawn, _ := strings.Cut(allNumbers, "|")

	for _, n := range strings.Split(strings.TrimSpace(winners), " ") {
		if n == "" || n == " " {
			continue
		}
		x, _ := strconv.Atoi(n)
		c.winners[x] = nothing{}
	}

	for _, n := range strings.Split(strings.TrimSpace(drawn), " ") {
		if n == "" || n == " " {
			continue
		}
		x, _ := strconv.Atoi(n)
		c.drawn = append(c.drawn, x)
	}
}

func (c *card) Value() int {
	value := 0
	for _, x := range c.drawn {
		_, ok := c.winners[x]
		if !ok {
			continue
		}

		if value == 0 {
			value = 1
		} else {
			value *= 2
		}
	}
	return value
}

func (c *card) ValueAndCount() (int, int) {
	value := 0
	count := 0
	for _, x := range c.drawn {
		_, ok := c.winners[x]
		if !ok {
			continue
		}

		if value == 0 {
			value = 1
			count = 1
		} else {
			value *= 2
			count++
		}
	}
	return value, count
}
