package main

import (
	"fmt"
	"strconv"
	"strings"
)

func d2() {
	fmt.Printf("Welcome to day2!\n")
	// s := "Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green"
	// Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
	games, _ := Read("d2.txt")
	total := 0
	for _, game := range games {

		gameID, rounds := splitGame(game)
		// Part 1
		// isPossible := checkColors(rounds)
		// fmt.Printf("In game: %d, counted RGB %d %d %d ", gameID, r, g, b)
		// if isPossible {
		// 	fmt.Printf("Game %d is possible! \n", gameID)
		// 	total += gameID
		// }

		// Part 2
		red, green, blue := getMaxColors(rounds) // Part 1
		fmt.Printf("For game %d, min needed to play: %d %d %d \n", gameID, red, green, blue)
		power := red * green * blue
		total += power

	}

	fmt.Printf("Total %d \n", total)
}

func gameIsPossible(r, g, b int) bool {
	var rt, gt, bt = 12, 13, 14
	return r <= rt && g <= gt && b <= bt
}

func getMaxColors(rounds string) (int, int, int) {
	roundColors := strings.Split(rounds, ";")
	var red, green, blue int

	for _, roundColor := range roundColors {

		numColor := strings.Split(roundColor, ",")
		for _, numColor2 := range numColor {

			numColor3 := strings.Split(numColor2, " ")
			// space := numColor3[0]
			num := numColor3[1]
			color := numColor3[2]
			cnum, err := strconv.Atoi(num)
			if err != nil {
				fmt.Printf("Err converting number: %v", err)
			}

			// fmt.Printf("num: %s\n", num)
			// fmt.Printf("color: %s\n", color)

			// fmt.Printf("Splitting: %s %d\n", color, cnum)

			switch color {
			case "red":
				if cnum > red {
					red = cnum
				}
			case "green":
				if cnum > green {
					green = cnum
				}
			case "blue":
				if cnum > blue {
					blue = cnum
				}
			}

		}
	}

	return red, green, blue
}

func checkColors(rounds string) bool {
	roundColors := strings.Split(rounds, ";")
	// var red, green, blue int

	for _, roundColor := range roundColors {

		numColor := strings.Split(roundColor, ",")
		for _, numColor2 := range numColor {

			numColor3 := strings.Split(numColor2, " ")
			// space := numColor3[0]
			num := numColor3[1]
			color := numColor3[2]
			cnum, err := strconv.Atoi(num)
			if err != nil {
				fmt.Printf("Err converting number: %v", err)
			}

			// fmt.Printf("num: %s\n", num)
			// fmt.Printf("color: %s\n", color)

			// fmt.Printf("Splitting: %s %d\n", color, cnum)

			switch color {
			case "red":
				if !gameIsPossible(cnum, 0, 0) {
					return false
				}
			case "green":
				if !gameIsPossible(0, cnum, 0) {
					return false
				}
			case "blue":
				if !gameIsPossible(0, 0, cnum) {
					return false
				}
			}

		}
	}

	return true
}

func splitGame(s string) (int, string) {
	splitStr := strings.Split(s, ":")
	gameNum, _ := strconv.Atoi(splitStr[0][5:])
	return gameNum, splitStr[1]
}
