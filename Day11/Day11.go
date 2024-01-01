package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

var GRID_WIDTH int
var EXPANSION_SIZE int

var NON_EMPTY_ROWS []int
var NON_EMPTY_COLUMNS []int
var GALAXIES []int

func parseGrid(lines []string) {
	// Search GALAXIES and remember rows and colums where at least one galaxy is located
	for y, line := range lines {
		for x, column := range line {
			if column == '#' {
				if slices.Contains(NON_EMPTY_ROWS, y) == false {
					NON_EMPTY_ROWS = append(NON_EMPTY_ROWS, y)
				}

				if slices.Contains(NON_EMPTY_COLUMNS, x) == false {
					NON_EMPTY_COLUMNS = append(NON_EMPTY_COLUMNS, x)
				}

				GALAXIES = append(GALAXIES, y*GRID_WIDTH+x)
			}
		}
	}
}

func calculateDistance(galaxy1 int, galaxy2 int) int {
	x1 := galaxy1 % GRID_WIDTH
	y1 := galaxy1 / GRID_WIDTH
	x2 := galaxy2 % GRID_WIDTH
	y2 := galaxy2 / GRID_WIDTH

	xStart := min(x1, x2)
	xEnd := max(x1, x2)
	yStart := min(y1, y2)
	yEnd := max(y1, y2)

	xDistance := 0
	for x := xStart; x <= xEnd; x++ {
		if slices.Contains(NON_EMPTY_COLUMNS, x) {
			xDistance++
		} else {
			xDistance += EXPANSION_SIZE
		}
	}

	yDistance := 0
	for y := yStart; y <= yEnd; y++ {
		if slices.Contains(NON_EMPTY_ROWS, y) {
			yDistance++
		} else {
			yDistance += EXPANSION_SIZE
		}
	}

	return xDistance + yDistance - 2 // Don't count the position of the first galaxy and only count the corner twice
}

func run(lines []string) int {
	total := 0
	for i := 0; i < len(GALAXIES)-1; i++ {
		for j := i + 1; j < len(GALAXIES); j++ {
			galaxy1 := GALAXIES[i]
			galaxy2 := GALAXIES[j]
			distance := calculateDistance(galaxy1, galaxy2)
			total += distance
		}
	}

	return total
}

func main() {
	file, _ := os.Open("input.txt")

	scanner := bufio.NewScanner(file)

	var lines []string
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	GRID_WIDTH = len(lines[0])

	parseGrid(lines)

	EXPANSION_SIZE = 2
	fmt.Printf("Puzzle 1: %d\n", run(lines))
	EXPANSION_SIZE = 1000000
	fmt.Printf("Puzzle 2: %d\n", run(lines))
}
