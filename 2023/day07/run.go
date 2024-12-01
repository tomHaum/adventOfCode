package day07

import (
	"cmp"
	"fmt"
	"slices"
	"strconv"
	"strings"
	"unicode"
)

func Run(input string) (int, int) {
	input = strings.ReplaceAll(input, "\r\n", "\n")

	split := strings.Split(input, "\n")
	hands := make([]Hand, len(split))
	for i, l := range split {
		h := ParseHand(l)
		h.originalOrder = i
		hands[i] = h
	}

	slices.SortFunc(hands, CompareHand)
	part1 := 0
	for i, h := range hands {
		//h.Print()
		mult := (i + 1) * h.Bet
		part1 += mult
	}

	slices.SortFunc(hands, CompareHand2)
	part2 := 0
	for i, h := range hands {
		//h.Print()
		mult := (i + 1) * h.Bet
		part2 += mult
	}
	return part1, part2
}

type Type string

const (
	Type5     Type = "5"
	Type4     Type = "4"
	TypeFull  Type = "Full"
	Type3     Type = "3"
	Type2Pair Type = "2"
	Type1Pair Type = "1"
	TypeHigh  Type = "High"
)

const aBigger = 1
const bBigger = -1
const bothEqual = 0

func CompareHand(a, b Hand) int {
	typeCompare := CompareType(a.Type, b.Type)
	if typeCompare != 0 {
		return typeCompare
	}

	for i, _ := range a.Cards {
		cardCompare := CompareCard(a.Cards[i], b.Cards[i])
		if cardCompare != 0 {
			return cardCompare
		}
	}

	panic("unexpected")
}

func CompareHand2(a, b Hand) int {
	typeCompare := CompareType(a.Type2, b.Type2)
	if typeCompare != 0 {
		return typeCompare
	}

	for i, _ := range a.Cards {
		cardCompare := CompareCard2(a.Cards[i], b.Cards[i])
		if cardCompare != 0 {
			return cardCompare
		}
	}

	panic("unexpected")

}

func CompareCard2(a, b Card) int {
	aIsJack := a == 'J'
	bIsJack := b == 'J'
	if aIsJack && bIsJack {
		return bothEqual
	}
	if aIsJack && !bIsJack {
		return bBigger
	}
	if !aIsJack && bIsJack {
		return aBigger
	}

	return CompareCard(a, b)
}

func CompareCard(a, b Card) int {
	aIsSuit := unicode.IsLetter(rune(a))
	bIsSuit := unicode.IsLetter(rune(b))

	if !aIsSuit && !bIsSuit {
		return cmp.Compare(rune(a)-'0', rune(b)-'0')
	}
	if aIsSuit && !bIsSuit {
		return aBigger
	}
	if !aIsSuit && bIsSuit {
		return bBigger
	}
	if aIsSuit && bIsSuit {
		switch a {
		case 'A':
			switch b {
			case 'A':
				return bothEqual
			default:
				return aBigger
			}
		case 'K':
			switch b {
			case 'A':
				return bBigger
			case 'K':
				return bothEqual
			default:
				return aBigger
			}
		case 'Q':
			switch b {
			case 'A', 'K':
				return bBigger
			case 'Q':
				return bothEqual
			default:
				return aBigger
			}
		case 'J':
			switch b {
			case 'A', 'K', 'Q':
				return bBigger
			case 'J':
				return bothEqual
			default:
				return aBigger
			}
		case 'T':
			switch b {
			case 'A', 'K', 'Q', 'J':
				return bBigger
			case 'T':
				return bothEqual
			default:
				panic("unexpected")
			}
		}
	}
	panic("unexpected")
}

func CompareType(a, b Type) int {
	switch a {
	case Type5:
		switch b {
		case Type5:
			return bothEqual
		default:
			return aBigger
		}
	case Type4:
		switch b {
		case Type5:
			return bBigger
		case Type4:
			return bothEqual
		default:
			return aBigger
		}
	case TypeFull:
		switch b {
		case Type5, Type4:
			return bBigger
		case TypeFull:
			return bothEqual
		default:
			return aBigger
		}
	case Type3:
		switch b {
		case Type5, Type4, TypeFull:
			return bBigger
		case Type3:
			return bothEqual
		default:
			return aBigger
		}
	case Type2Pair:
		switch b {
		case Type5, Type4, TypeFull, Type3:
			return bBigger
		case Type2Pair:
			return bothEqual
		default:
			return aBigger
		}
	case Type1Pair:
		switch b {
		case Type5, Type4, TypeFull, Type3, Type2Pair:
			return bBigger
		case Type1Pair:
			return bothEqual
		default:
			return aBigger
		}
	case TypeHigh:
		switch b {
		case Type5, Type4, TypeFull, Type3, Type2Pair, Type1Pair:
			return bBigger
		case TypeHigh:
			return bothEqual
		default:
			panic("unexpected")
		}
	}
	panic("unexpected")
}

type Hand struct {
	originalOrder int
	Cards         []Card
	Bet           int

	Type  Type
	Type2 Type
}

func ParseHand(line string) Hand {
	cards, betTxt, _ := strings.Cut(line, " ")
	bet, _ := strconv.Atoi(betTxt)

	h := Hand{Cards: []Card(cards), Bet: bet}
	h.SetType()
	h.SetType2()

	return h
}

func (h *Hand) SetType() {
	cardFreq := make(map[Card]uint8)
	for _, c := range h.Cards {
		cardFreq[c] += 1
	}

	freq := make([]uint8, 0, len(cardFreq))
	for _, n := range cardFreq {
		freq = append(freq, n)
	}

	handType := freqToType(freq)
	h.Type = handType

	return
}

func (h *Hand) SetType2() {
	cardFreq := make(map[Card]uint8)
	for _, c := range h.Cards {
		cardFreq[c] += 1
	}

	freq := make([]uint8, 0, len(cardFreq))
	var jokers uint8 = 0
	for c, n := range cardFreq {
		if c == 'J' {
			jokers = n
			continue
		}
		freq = append(freq, n)
	}

	if jokers != 0 {
		if len(freq) == 0 {
			freq = append(freq, jokers)
		} else {
			slices.Sort(freq)
			freq[len(freq)-1] += jokers
		}
	}

	handType := freqToType(freq)
	h.Type2 = handType

	return
}

func freqToType(freq []uint8) Type {
	slices.Sort(freq)

	switch len(freq) {
	case 1:
		return Type5
	case 2:
		if freq[1] == 4 {
			return Type4
		} else {
			return TypeFull
		}
	case 3:
		if freq[2] == 3 {
			return Type3
		} else {
			return Type2Pair
		}
	case 4:
		return Type1Pair
	case 5:
		return TypeHigh
	}
	panic("unexpected")
}

type Card rune

func (c Card) String() string {
	return fmt.Sprintf("%s", string(c))
}

func (c Card) GoString() string {
	return c.String()
}

func (h *Hand) Print() {
	fmt.Printf("%v\t%v\t%v\t%v\n", h.originalOrder, string(h.Cards), h.Bet, h.Type)
}
