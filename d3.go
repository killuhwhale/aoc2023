package main

import (
	"fmt"
	"strconv"
	"unicode"
)

type Pair struct {
	x int
	y int
}

func d3() {
	fmt.Printf("Helo day 3! \n")

	grid, _ := Read("d3.txt")

	var dirs []Pair
	dirs = append(dirs, Pair{-1, 0})
	dirs = append(dirs, Pair{-1, 1})
	dirs = append(dirs, Pair{0, 1})
	dirs = append(dirs, Pair{1, 1})
	dirs = append(dirs, Pair{1, 0})
	dirs = append(dirs, Pair{1, -1})
	dirs = append(dirs, Pair{0, -1})
	dirs = append(dirs, Pair{-1, -1})

	nodes := make(map[Pair]bool)
	ans := 0
	for y, row := range grid {
		// fmt.Printf("Row (%d): %s\n", y, row)

		for x, cell := range row {
			// part 1, we looked all special characters.
			// if string(cell) != "." && !unicode.IsDigit(cell) {

			// part 2, we only want to look at gears *
			if string(cell) == "*" {
				// fmt.Printf("We have a special character, lets look at all neighbors: %s (%d, %d) \n", string(cell), x, y)
				tmp := getNeighbors(x, y, grid, nodes, dirs, 0) // DFS
				ans += tmp
				// New numbers found after this iteration.

				// fmt.Printf("reported in round: %d \n", tmp)
			}
		}

	}

	fmt.Printf("Ans: %d \n", ans)

}

func inSet(nodes map[Pair]bool, x, y int) bool {
	p := Pair{x, y}
	_, exists := nodes[p]
	return exists
}

// Create a global the LIST to store all cells processed (i.e numbers counted towards sum already)
func getNeighbors(x, y int, grid []string, nodes map[Pair]bool, dirs []Pair, depth int) int {
	// Start by looking at all diagonals and adjacent
	// If its a digit, add that cell to the LIST
	// Get the rest of the digits and add to the LIST
	// add number to total
	p := Pair{x, y}
	ans := 0
	if x < 0 || y < 0 || x >= len(grid[0]) || y >= len(grid) || inSet(nodes, x, y) {
		return 0
	}
	nodes[p] = true
	// fmt.Printf("Current node: %s (d=%d)\n", string(grid[y][x]), depth)

	// On initial depth == 0 we have a special character and need to check surroundings
	if depth == 0 {
		a := getNeighbors(x+-1, y+0, grid, nodes, dirs, depth+1)
		b := getNeighbors(x+-1, y+1, grid, nodes, dirs, depth+1)
		c := getNeighbors(x+0, y+1, grid, nodes, dirs, depth+1)
		d := getNeighbors(x+1, y+1, grid, nodes, dirs, depth+1)
		e := getNeighbors(x+1, y+0, grid, nodes, dirs, depth+1)
		f := getNeighbors(x+1, y+-1, grid, nodes, dirs, depth+1)
		g := getNeighbors(x+0, y+-1, grid, nodes, dirs, depth+1)
		h := getNeighbors(x+-1, y+-1, grid, nodes, dirs, depth+1)

		lol := []int{a, b, c, d, e, f, g, h}
		count := 0
		total := 1
		for _, result := range lol {
			if result > 0 {
				count++
				total *= result
			}
		}

		if count == 2 {
			ans += total
		}

		fmt.Printf("After depth 0 search we found %d numbers: %v \n", count, lol)

	} else {
		// we are now passed the special character
		isDigit := unicode.IsDigit(rune(grid[y][x]))

		if isDigit {
			// Now we want to check left and right for a number, and add the cells to nodes visited.
			// fmt.Printf("isDigit: %v %s %d \n", rune(grid[y][x]), string(grid[y][x]), grid[y][x])
			// fmt.Printf("Checking left and right starting at cell: %s  (%d, %d) \n", string(grid[y][x]), x, y)
			num := getNum(grid, x, y, nodes)
			fmt.Printf("Found num: %d \n", num)

			return num
		}
	}
	return ans
}

func getNum(grid []string, x, y int, nodes map[Pair]bool) int {
	// Left seach
	i := x
	j := x
	for i >= 0 {
		if unicode.IsDigit(rune(grid[y][i])) {
			p := Pair{i, y}
			nodes[p] = true
			if i > 0 {
				i--
			} else {
				break
			}

		} else {
			break
		}
	}

	if !unicode.IsDigit(rune(grid[y][i])) {
		i++
	}

	for j < len(grid[y]) {
		if unicode.IsDigit(rune(grid[y][j])) {
			p := Pair{j, y}
			nodes[p] = true
			if j < len(grid[y]) {
				j++
			} else {
				break
			}
		} else {
			break
		}
	}

	num, _ := strconv.Atoi(string(grid[y][i:j]))
	return num
}
