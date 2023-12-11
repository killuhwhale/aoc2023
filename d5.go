package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
	// "regexp"
)

func main() {
	fmt.Printf("Day 5! \n")

	// dest - src - range
	// seed-to-soil map:
	// 50 98 2
	// 52 50 48
	allMaps, _ := Read("d5.txt")

	maps, seeds := splitUpFarmList(allMaps)
	fmt.Printf("Split maps\n")
	loc := 0
	mloc := 2 ^ 63
	for _, seed := range seeds {
		loc = getLoc(seed, maps)
		fmt.Printf("Location for seed %d  = %d \n", seed, loc)
		if loc < mloc {
			mloc = loc
		}
	}
	fmt.Printf("Lowest Location %d \n", mloc)

}

func getLoc(seed int, maps []map[int]int) int {
	// last := seed
	current := seed
	for _, m := range maps {
		// fmt.Printf("Mapping lvl (%d) src: %d - dest %d  Map: %v  \n", i, current, m[current], m)
		tmp, exists := m[current]
		if exists {
			current = tmp
		}
	}
	return current
}

func splitUpFarmList(allMaps []string) ([]map[int]int, []int) {
	seedListStr := allMaps[0]
	seeds := getSeeds(seedListStr)

	maps := make([]map[int]int, 7)

	m := make(map[int]int)
	// header := ""
	mc := 0
	for i, line := range allMaps[2:] {
		fmt.Printf("Line(%d): %s \n", i, line)
		if line == "" {
			// fmt.Printf("Line [newline] (%d): %s \n", i, line)
			maps[mc] = m
			mc++
			m = make(map[int]int)
		} else if unicode.IsDigit(rune(line[0])) {
			addToMap(line, m)
		}
		//  else if strings.Contains(line, "-") {
		// 	header = line
		// 	fmt.Printf("Header: %s \n", header)
		// }
	}
	maps[mc] = m
	return maps, seeds
}

func addToMap(srcDestRangeStr string, m map[int]int) {
	srcDestRange := strsToNums(strings.Split(srcDestRangeStr, " "))
	d := srcDestRange[0] // Dest start
	s := srcDestRange[1] // Src Start
	r := srcDestRange[2] // Range
	i := 0
	for i < r {
		m[s] = d
		d++
		s++
		i++
	}
	// fmt.Printf("Created map: %v \n", m)
}

func getSeeds(seedStr string) []int {
	seedsStrArr := strings.Split(seedStr[len("seeds: "):], " ")
	seeds := strsToNums(seedsStrArr)
	fmt.Printf("Considering seeds: %v \n", seeds)
	return seeds
}

func strsToNums(numStr []string) []int {
	nums := make([]int, len(numStr))
	for i, n := range numStr {
		nums[i], _ = strconv.Atoi(n)
	}
	return nums
}
