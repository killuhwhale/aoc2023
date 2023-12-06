package main

import (
	"fmt"
	"unicode"
)

type Pair struct {
	x int
	y int
}

func main() {
	fmt.Printf("Helo day 3! \n")

	grid, _ := Read("d3.txt")
	var nodes []Pair

	for y, row := range grid {
		fmt.Printf("Row (%d): %s\n", y, row)

		for x, cell := range row {
			if string(cell) != "." && !unicode.IsDigit(cell) {
				fmt.Printf("We have a special character, lets look at all neighbors: %s (%d, %d) \n", string(cell), x, y)
				getNeighbors(x, y, grid, nodes)
			}
		}

	}

}

// Create a global the LIST to store all cells processed (i.e numbers counted towards sum already)
func getNeighbors(x, y int, grid []string, nodes []Pair) {
	// Start by looking at all diagonals and adjacent
	// If its a digit, add that cell to the LIST
	// Get the rest of the digits and add to the LIST
	// add number to total

	var dirs []Pair
	dirs = append(dirs, Pair{-1, 0})
	dirs = append(dirs, Pair{-1, 1})
	dirs = append(dirs, Pair{0, 1})
	dirs = append(dirs, Pair{1, 1})
	dirs = append(dirs, Pair{1, 0})
	dirs = append(dirs, Pair{1, -1})
	dirs = append(dirs, Pair{0, -1})
	dirs = append(dirs, Pair{-1, -1})

	for _, dir := range dirs {
		// Check cell
		checkX := dir.x + x
		checkY := dir.y + y
		if unicode.IsDigit(rune(grid[checkY][checkX])) {
			fmt.Printf("Check cell (%d, %d)", checkX, checkY)
		}

	}

}
