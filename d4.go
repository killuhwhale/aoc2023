package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	// "strconv"
	// "unicode"
)

func d4() {
	fmt.Printf("Yo, Day 4! \n")

	games, _ := Read("d4.txt")
	N := len(games)
	cards := make([]int, N)

	for i := range cards {
		cards[i] = 1
	}

	// ans := 0
	// Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53
	for i, game := range games {
		// points := getPoints(game)
		numCards := getPoints(game)
		// ans += points

		fmt.Printf("Adding cards..\n")
		addCards(cards, i+1, numCards, cards[i])
	}

	// fmt.Printf("Ans: %d \n", ans)
	ans := 0
	for _, n := range cards {
		ans += n
	}
	fmt.Printf("Ans: %v %d \n", cards, ans)

}

func addCards(cards []int, start, numCards, amount int) {
	i := start
	for numCards > 0 {
		cards[i] += amount
		numCards--
		i++
	}
}

func inGameNums(gameNums map[string]bool, s string) bool {
	_, exists := gameNums[s]
	return exists
}

func getPoints(game string) int {

	re, _ := regexp.Compile(`^Card \d+: `)
	numbers := re.ReplaceAllString(game, "")
	allNums := strings.Split(numbers, "|")

	gameNums := make(map[string]bool) //allNums[0]
	playerNums := strings.Split(allNums[1], " ")
	nums := strings.Split(allNums[0], " ")
	// Add game nums to set
	for _, n := range nums {
		na, _ := strconv.Atoi(n)
		fmt.Printf("Adding to gamenum: %s %d \n", n, na)
		if n != " " && n != "" {
			gameNums[n] = true
		}
	}

	points := 0
	for idx, playerNum := range playerNums {
		if inGameNums(gameNums, playerNum) {
			fmt.Printf("%d Found player num: %s in game nums: %s ~~~ %s \n", idx, playerNum, nums, playerNums)
			// Part 1
			// if points == 0 {
			// 	points = 1
			// } else {
			// 	points *= 2
			// }

			// Part 2
			points++

		}
	}
	return points
}
