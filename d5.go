package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"unicode"
	// "regexp"
)

func d5() {
	fmt.Printf("Day 5! \n")

	// dest - src - range
	// seed-to-soil map:
	// 50 98 2
	// 52 50 48
	allMaps, _ := Read("d5.txt")

	maps, seeds := splitUpFarmList(allMaps)
	fmt.Printf("Split maps\n")
	// fmt.Printf("Maps len=(%d) %v   \n", len(maps), maps)
	var loc int64
	var mloc int64
	loc = 0
	mloc = math.MaxInt64
	fmt.Printf("Lowest Location %d \n", mloc)
	pt2Seeds := getSeedRanges(seeds)
	fmt.Printf("pt2Seeds %v \n", pt2Seeds)
	// for _, seed := range seeds { // part 1
	for _, seedPair := range pt2Seeds {
		fmt.Printf("seedPair %v \n", seedPair)

		startSeed := seedPair[0]
		endSeed := seedPair[1] + startSeed

		for seed := startSeed; seed <= endSeed; seed++ {
			loc = getLoc(seed, maps)
			// fmt.Printf("Location for seed %d  = %d \n", seed, loc)
			if loc < mloc {
				mloc = loc
				fmt.Printf("Lowest Location %d \n", mloc)
			}

		}

	}
	fmt.Printf("Lowest Location %d \n", mloc)

}

// [961540761 489996751]
// Lowest Location 171065507

func getSeedRanges(seeds []int64) [][]int64 {

	var pairs [][]int64
	var pair []int64
	for i, n := range seeds {

		pair = append(pair, n)
		if i%2 == 1 {
			pairs = append(pairs, pair)
			pair = nil
		}
	}
	return pairs
}

func getLoc(seed int64, maps [][][]int64) int64 {
	// last := seed
	current := seed
	// fmt.Printf("-> %d ", current)
	for _, group := range maps {
		found := false
		for _, mapping := range group {
			if !found && current >= mapping[0] && current < mapping[0]+mapping[2] {
				delta := current - mapping[0]
				// fmt.Printf(" (%d+%d) ", delta, mapping[1])
				current = mapping[1] + delta // dest start plus delta from current and src
				found = true
			}
		}
		// fmt.Printf("-> %d ", current)
	}

	return current
}

func splitUpFarmList(allMaps []string) ([][][]int64, []int64) {
	seedListStr := allMaps[0]
	seeds := getSeeds(seedListStr)

	maps := make([][][]int64, 7)

	m := make([]int64, 3)
	var group [][]int64
	// header := ""
	mc := 0
	for _, line := range allMaps[3:] {
		// fmt.Printf("Line(%d): %s \n", i, line)
		if line == "" {
			// fmt.Printf("Line [newline] (%d): %s \n", i, line)
			maps[mc] = group
			// maps[mc] = m
			mc++
			m = make([]int64, 3)
			// Todo reint group slice to an empty slice
			group = nil

		} else if unicode.IsDigit(rune(line[0])) {
			m = addToGroup(line)

			group = append(group, m)
		}
		//  else if strings.Contains(line, "-") {
		// 	header = line
		// 	fmt.Printf("Header: %s \n", header)
		// }
	}
	maps[mc] = group
	return maps, seeds
}

func addToGroup(srcDestRangeStr string) []int64 {
	srcDestRange := strsToNums(strings.Split(srcDestRangeStr, " "))
	d := srcDestRange[0] // Dest start
	s := srcDestRange[1] // Src Start
	r := srcDestRange[2] // Range
	return []int64{s, d, r}
}

func getSeeds(seedStr string) []int64 {
	seedsStrArr := strings.Split(seedStr[len("seeds: "):], " ")
	seeds := strsToNums(seedsStrArr)
	fmt.Printf("Considering seeds: %v \n", seeds)
	return seeds
}

func strsToNums(numStr []string) []int64 {
	nums := make([]int64, len(numStr))
	for i, n := range numStr {
		nums[i], _ = strconv.ParseInt(n, 10, 64)
	}
	return nums
}
