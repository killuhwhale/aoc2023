package main

import (
	"errors"
	"fmt"

	"strconv"
	"strings"
	"unicode"
)

func d1() {
	cals, err := Read("d1.txt")
	if err != nil {
		fmt.Print("Fail to read file!")
	}
	ans := 0
	// cals = []string{"abc", "xyz123"}
	for _, cal := range cals {
		// total, err := calTotal(cal)
		first, err := findFirst(cal)
		last, err := findLast(cal)
		fmt.Printf("Found first number %d in %s: \n", first, cal)
		fmt.Printf("Found last number %d in %s: \n", last, cal)

		if err != nil {
			fmt.Println("error calculating calibration total: ", err)
		}

		s, err := strconv.Atoi(fmt.Sprint(first) + fmt.Sprint(last))
		if err == nil {
			ans += s
		}
		fmt.Println("Total: ", first+last)

	}

	fmt.Println("Ans: ", ans)
}

func findFirst(cal string) (int, error) {

	digitChars := []string{"o", "t", "f", "s", "e", "n", "z"}

	for i, ch := range cal {
		// fmt.Printf("Checking char: %s \n", string(ch))
		// If the current char is an int
		if unicode.IsDigit(ch) {
			return strconv.Atoi(string(ch))
		}

		// if the char starts with something like our numbers
		letterIdx := FindIndexFunc(digitChars, func(letter string) bool { return letter == string(ch) })
		if letterIdx >= 0 {
			// The current char is the beginnining of a digit potentially
			// Lets call a func with this index to get the rest of the digit word
			restofword := cal[i:]
			// fmt.Printf("getDigitWord, checking: %v \n", restofword)
			first := getDigitWord(restofword)
			if first >= 0 {
				return first, nil
			}
		}

	}

	return -1, errors.New("Failed to find a number....")
}

func findLast(cal string) (int, error) {

	digitChars := []string{"o", "t", "f", "s", "e", "n", "z"}
	N := len(cal)
	for ii, _ := range cal {
		// fmt.Printf("Checking char: %s \n", string(ch))
		// If the current char is an int
		i := N - 1 - ii
		ch := rune(cal[i])
		if unicode.IsDigit(ch) {
			return strconv.Atoi(string(ch))
		}

		// if the char starts with something like our numbers
		letterIdx := FindIndexFunc(digitChars, func(letter string) bool { return letter == string(ch) })
		if letterIdx >= 0 {
			// The current char is the beginnining of a digit potentially
			// Lets call a func with this index to get the rest of the digit word
			restofword := cal[i:]
			// fmt.Printf("getDigitWord, checking: %v \n", restofword)
			first := getDigitWord(restofword)
			if first >= 0 {
				return first, nil
			}
		}

	}

	return -1, errors.New("Failed to find a number....")
}

func getDigitWord(cal string) int {
	digitWords := []string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

	for val, digitWord := range digitWords {
		if strings.HasPrefix(cal, digitWord) {
			return val
		}
	}
	return -1

}

func calTotal(cal string) (int, error) {
	N := len(cal) - 1
	i := 0
	j := N

	fstart := false
	fend := false
	first := ""
	last := ""

	for i <= N && j >= 0 {
		// Check each idx for a number.
		if !fstart && unicode.IsDigit(rune(cal[i])) {
			fstart = true
			first = string(cal[i])
		} else {
			i++
		}

		if !fend && unicode.IsDigit(rune(cal[j])) {
			fend = true
			last = string(cal[j])
		} else {
			j--
		}

		if fstart && fend {
			// fmt.Println("Breaking")
			break
		}

	}

	// fmt.Println("Found first - last: ", cal, first, " / ", last)
	firstInt, err := strconv.Atoi(first + last)

	if err != nil {
		fmt.Println("Failed converting number: ", err)
		return 0, err
	}
	return firstInt, nil

}
