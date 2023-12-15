package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Printf("Day 6! \n")

	// dest - src - range
	// seed-to-soil map:
	// 50 98 2
	// 52 50 48
	data, _ := Read("d6.txt")
	times, distances := getRaces(data)

	var totalWays int64 = 0
	var ans int64 = 1

	for i, race := range times {
		fmt.Printf("Race: %v %v \n", race, distances[i])
		bestDist := distances[i]
		for j := 0; j < race; j++ {
			dist := getDistance(j, race)
			if dist > bestDist {
				totalWays++
			}
			// fmt.Printf("Distance: %d %v \n", j, dist)

		}
		// fmt.Printf("Total Ways: %d \n", totalWays)
		ans *= totalWays
		totalWays = 0
	}
	fmt.Printf("Answer: %d \n", ans)
	// fmt.Printf("Races: %v \n", times)
	// fmt.Printf("Distances: %v \n", distances)

}

func getRaces(raceStrings []string) ([]int, []int) {
	timeStrings := strings.Split(strings.TrimSpace(strings.TrimPrefix(raceStrings[0], "Time:")), " ")
	distanceStrings := strings.Split(strings.TrimSpace(strings.TrimPrefix(raceStrings[1], "Distance:")), " ")

	times := make([]int, 0)
	distances := make([]int, 0)

	for _, s := range timeStrings {

		if val, err := strconv.Atoi(s); err == nil {
			times = append(times, val)
		}

	}

	for _, s := range distanceStrings {
		if val, err := strconv.Atoi(s); err == nil {
			distances = append(distances, val)
		}
	}

	return times, distances
}

func getDistance(rate, limit int) int {
	if rate == 0 {
		return 0
	}
	if rate == limit {
		return 0
	}

	// Multiply the remaining time by the amount we held the button for.
	return (limit - rate) * rate
}
