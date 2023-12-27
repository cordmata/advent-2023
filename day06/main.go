package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(part1())
	fmt.Println(part2())
}

func part1() int {
	races := parseRaces(actual)
	winnerCounts := []int{}
	for _, r := range races {
		var numWinners int
		for holdDuration := 0; holdDuration < r.time; holdDuration++ {
			dist := holdDuration * (r.time - holdDuration)
			if dist > r.distance {
				numWinners++
			}
		}
		winnerCounts = append(winnerCounts, numWinners)
	}

	product := 1
	for _, c := range winnerCounts {
		if c > 0 {
			product *= c
		}
	}
	return product
}

func part2() int {
	return -1
}

type race struct {
	time     int
	distance int
}

func parseRaces(in string) []race {
	times := []int{}
	distances := []int{}
	races := []race{}
	for _, line := range strings.Split(in, "\n") {
		parts := strings.Fields(line)
		label := parts[0]
		vals := parts[1:]
		for _, v := range vals {
			v, _ := strconv.Atoi(v)
			if label == "Time:" {
				times = append(times, v)
			}
			if label == "Distance:" {
				distances = append(distances, v)
			}
		}
	}
	for i, time := range times {
		races = append(races, race{time: time, distance: distances[i]})
	}
	return races
}

const example = `Time:      7  15   30
Distance:  9  40  200`

const actual = `Time:        45     98     83     73
Distance:   295   1734   1278   1210`
