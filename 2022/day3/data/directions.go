package data

import (
	"bufio"
	"os"
)

type Sack struct {
	Pouch1        Pouch
	Pouch2        Pouch
	OriginalValue string
	Duplicates    map[Item]bool
	Badge         Item
}
type Item rune
type Pouch map[Item]int

// GetData true for up false for down
func GetData(fileName string) []Sack {
	file, err := os.Open(fileName)
	Check(err)
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	sacks := make([]Sack, 0)
	badgeCandidates := make(map[Item]bool)
	lineNumber := 1
	for scanner.Scan() {
		line := scanner.Text()

		pouch1 := make(Pouch)
		pouch2 := make(Pouch)
		duplicates := make(map[Item]bool)
		uniques := make(map[Item]bool)
		for i, c := range line {
			item := Item(c)
			uniques[item] = true
			if i < len(line)/2 {
				pouch1[item]++
			} else {
				pouch2[item]++

				if _, ok := pouch1[item]; ok {
					duplicates[item] = true
				}
			}

		}

		// first bag
		if (lineNumber)%3 == 1 { //1, 5, 8
			// initialize set
			badgeCandidates = uniques
		} else {
			// intersect set
			intersect := make(map[Item]bool)

			for item, _ := range badgeCandidates {
				if uniques[item] {
					intersect[item] = true
				}
			}

			badgeCandidates = intersect
		}

		sack := Sack{pouch1, pouch2, line, duplicates, Item('a')}
		sacks = append(sacks, sack)

		// last bag
		if (lineNumber)%3 == 0 { //3, 6, 9
			badge := getBadge(badgeCandidates)
			// correct the badge
			sacks[lineNumber-1].Badge = badge
			sacks[lineNumber-2].Badge = badge
			sacks[lineNumber-3].Badge = badge
			// reset badge candidates
			delete(badgeCandidates, badge)
		}

		lineNumber++
	}

	return sacks
}

func getBadge(badges map[Item]bool) Item {
	if len(badges) != 1 {
		panic("too many badges")
	}
	for k, _ := range badges {
		return k
	}
	panic("no badges")
}

func GetPriority(item Item) int {
	if item < 'a' {
		return int(item) - int('A') + 1 + 26
	} else {
		return int(item) - int('a') + 1
	}
}

func Check(e error) {
	if e != nil {
		panic(e)
	}
}
