package main

import (
	"fmt"
	"strconv"
	"unicode"
)

func main() {
	cals, err := Read("d1.txt")
	if err != nil {
		fmt.Print("Fail to read file!")
	}
	ans := 0
	// cals = []string{"abc", "xyz123"}
	for _, cal := range cals {
		total, err := calTotal(cal)
		if err != nil {
			fmt.Println("error calculating calibration total: ", total)
		}
		ans += total
		fmt.Println("Total: ", total)

	}

	fmt.Println("Ans: ", ans)
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
