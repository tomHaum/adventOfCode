package day06

import (
	"strconv"
	"strings"
)

func Run(input string) (int, int) {
	input = strings.ReplaceAll(input, "\r\n", "\n")
	split := strings.Split(input, "\n")
	times := split[0]
	_, times, _ = strings.Cut(times, ":")
	times = strings.TrimSpace(times)
	distances := split[1]
	_, distances, _ = strings.Cut(distances, ":")
	distances = strings.TrimSpace(distances)

	timesSplit := strings.Fields(times)
	distancesSplit := strings.Fields(distances)

	races := make([]Race, 0)
	for i, _ := range timesSplit {
		t, _ := strconv.Atoi(timesSplit[i])
		d, _ := strconv.Atoi(distancesSplit[i])

		races = append(races, Race{Time: t, Distance: d})
	}

	part1 := 1
	for _, r := range races {
		v := r.Solutions()
		part1 *= v
	}

	part2 := 0
	return part1, part2
}

type Race struct {
	Time     int
	Distance int
}

func (r Race) Solutions() int {
	count := 0
	for i := 0; i < r.Time; i++ {
		speed := i
		time := r.Time - i
		distance := time * speed

		if distance > r.Distance {
			count++
		}
	}

	return count
}
