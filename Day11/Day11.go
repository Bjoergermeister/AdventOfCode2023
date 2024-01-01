package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

var GRID_WIDTH int

func calculateDistance(galaxy1 int, galaxy2 int, nonEmptyRows []int, nonEmptyColumns []int) int {
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
		if slices.Contains(nonEmptyColumns, x) {
			xDistance++
		} else {
			xDistance += 2
		}
	}

	yDistance := 0
	for y := yStart; y <= yEnd; y++ {
		if slices.Contains(nonEmptyRows, y) {
			yDistance++
		} else {
			yDistance += 2
		}
	}

	return xDistance + yDistance - 2 // Don't count the position of the first galaxy and only count the corner twice
}

func Puzzle1(lines []string) int {
	var nonEmptyRows []int
	var nonEmptyColumns []int
	var galaxies []int

	// Search galaxies and remember rows and colums where at least one galaxy is located
	for y, line := range lines {
		for x, column := range line {
			if column == '#' {
				if slices.Contains(nonEmptyRows, y) == false {
					nonEmptyRows = append(nonEmptyRows, y)
				}

				if slices.Contains(nonEmptyColumns, x) == false {
					nonEmptyColumns = append(nonEmptyColumns, x)
				}

				galaxies = append(galaxies, y*GRID_WIDTH+x)
			}
		}
	}

	total := 0
	for i := 0; i < len(galaxies)-1; i++ {
		for j := i + 1; j < len(galaxies); j++ {
			galaxy1 := galaxies[i]
			galaxy2 := galaxies[j]
			distance := calculateDistance(galaxy1, galaxy2, nonEmptyRows, nonEmptyColumns)
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

	fmt.Printf("Puzzle 1: %d\n", Puzzle1(lines))
}
