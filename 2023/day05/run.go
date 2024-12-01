package day05

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
	"sync"
)

const MaxUint = ^uint(0)
const MaxInt = int(MaxUint >> 1)

func Run(input string) (int, int) {
	input = strings.ReplaceAll(input, "\r\n", "\n")
	chunked := strings.Split(input, "\n\n")
	var seeds []int
	almanac := make(Almanac)
	for _, c := range chunked {
		split := strings.Split(c, "\n")
		if seeds == nil {
			seeds = getSeeds(split)
			continue
		}

		t := parseTransformer(split)
		almanac[t.Source] = t
	}

	fmt.Printf("%v\n", seeds)

	locations := make([]int, 0)
	for _, s := range seeds {
		locations = append(locations, almanac.Lookup("seed", "location", s))
	}
	fmt.Printf("%v\n", locations)
	slices.Sort(locations)
	seedPairs := [][]int{}
	for i := 0; i < len(seeds); i += 2 {
		seedPairs = append(seedPairs, seeds[i:i+2])
	}

	fmt.Println(seedPairs)

	length := 0
	for _, p := range seedPairs {
		length += p[1]
	}
	seeds = make([]int, length)
	j := 0
	for _, p := range seedPairs {
		for i := p[0]; i < p[0]+p[1]; i++ {
			seeds[j] = i
			j++
		}
	}

	fmt.Println("Done generating seeds:", len(seeds))
	part2mutext := sync.Mutex{}
	wg := sync.WaitGroup{}
	go func() {
		wg.Add(1)
	}()

	minPart2 := MaxInt
	doChunk := func(chunk []int) {
		wg.Add(1)
		min := MaxInt
		for _, s := range chunk {
			v := almanac.Lookup("seed", "location", s)
			if v < min {
				min = v
			}
		}
		part2mutext.Lock()
		defer part2mutext.Unlock()
		if min < minPart2 {
			minPart2 = min
		}
		wg.Done()
	}
	for _, c := range chunk(seeds, 1000000) {
		go doChunk(c)
	}

	go func() {
		wg.Done()
	}()
	wg.Wait()
	//
	//minPart2 := MaxInt
	//part2Mutex := sync.Mutex{}
	//wg := sync.WaitGroup{}
	//
	//seedTran := almanac["seed"]
	//soilTran := almanac["soil"]
	//fertTran := almanac["fertilizer"]
	//waterTran := almanac["water"]
	//lightTran := almanac["light"]
	//tempTran := almanac["temperature"]
	//humTran := almanac["humidity"]
	//
	//valFunc := func(input int) int {
	//	input = seedTran.Tranform(input)
	//	input = soilTran.Tranform(input)
	//	input = fertTran.Tranform(input)
	//	input = waterTran.Tranform(input)
	//	input = lightTran.Tranform(input)
	//	input = tempTran.Tranform(input)
	//	input = humTran.Tranform(input)
	//	return input
	//}
	//
	//calcRange := func(lower, length int) {
	//	wg.Add(1)
	//	min := MaxInt
	//	for i := lower; i < i+length; i++ {
	//		//l := almanac.Lookup("seed", "location", i)
	//		l := valFunc(i)
	//		if l < minPart2 {
	//			minPart2 = l
	//		}
	//	}
	//
	//	part2Mutex.Lock()
	//	if min < minPart2 {
	//		minPart2 = min
	//	}
	//	part2Mutex.Unlock()
	//	wg.Done()
	//}
	//
	//lower := -1
	//for _, x := range seeds {
	//	if lower == -1 {
	//		lower = x
	//		continue
	//	}
	//	go calcRange(lower, x)
	//	lower = -1
	//}
	//time.Sleep(100)
	//wg.Wait()
	return locations[0], minPart2
}

type Almanac map[string]Transformer

func (a Almanac) Lookup(source, dest string, value int) int {
	t, ok := a[source]
	if !ok {
		return -1
	}

	result := t.Tranform(value)
	if dest == t.Destination {
		return result
	}

	return a.Lookup(t.Destination, dest, result)
}

func chunk[T any](input []T, size int) [][]T {
	chunked := make([][]T, 0)

	for i := 0; i < len(input); i += size {
		end := i + size
		if end > len(input) {
			end = len(input)
		}

		chunked = append(chunked, input[i:end])
	}

	return chunked
}

func getSeeds(split []string) []int {
	rawSeeds := strings.TrimPrefix(split[0], "seeds: ")
	splitSeeds := strings.Fields(rawSeeds)
	seeds := make([]int, len(splitSeeds))
	for i, s := range splitSeeds {
		n, _ := strconv.Atoi(s)
		seeds[i] = n
	}

	return seeds
}

type Transformer struct {
	Source         string
	Destination    string
	Tranformations []Transformation
}

func (t Transformer) Tranform(input int) int {
	for _, tfm := range t.Tranformations {
		if tfm.Valid(input) {
			return tfm.Delta + input
		}
	}

	// default to same number
	return input
}

func parseTransformer(input []string) Transformer {
	mapping, _, _ := strings.Cut(input[0], " ")
	split := strings.Split(mapping, "-")
	source := split[0]
	dest := split[2]

	transformations := make([]Transformation, 0)
	for _, l := range input[1:] {
		transformations = append(transformations, parseTransformation(l))
	}

	return Transformer{
		Source:         source,
		Destination:    dest,
		Tranformations: transformations,
	}
}

type Transformation struct {
	SourceStart int
	SourceEnd   int
	Delta       int
}

func (t Transformation) Valid(input int) bool {
	return input >= t.SourceStart && input <= t.SourceEnd
}

func parseTransformation(line string) Transformation {
	split := strings.Split(line, " ")

	destinationStart := toInt(split[0])
	sourceStart := toInt(split[1])
	length := toInt(split[2])

	delta := destinationStart - sourceStart
	sourceEnd := sourceStart + length - 1
	return Transformation{
		SourceStart: sourceStart,
		SourceEnd:   sourceEnd,
		Delta:       delta,
	}
}

func toInt(s string) int {
	x, _ := strconv.Atoi(s)
	return x
}
