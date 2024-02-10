package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

const (
	TOP    = iota
	RIGHT  = iota
	BOTTOM = iota
	LEFT   = iota
)

// These are essential bit fields where four bits are used to mark whether tiles are present
// in the four possible directions (up, right, down, left (from least to most significant bit)).
// For example, VERTICAL is 0101, which means that the top and bottom tiles are present.
// Similarly, HORIZONTAL is 1010, meaning that the left and right tiles are set
const (
	VERTICAL     = 1<<TOP + 1<<BOTTOM
	HORIZONTAL   = 1<<LEFT + 1<<RIGHT
	TOP_RIGHT    = 1<<TOP + 1<<RIGHT
	RIGHT_BOTTOM = 1<<RIGHT + 1<<BOTTOM
	BOTTOM_LEFT  = 1<<BOTTOM + 1<<LEFT
	LEFT_TOP     = 1<<LEFT + 1<<TOP
)

var visitedTiles []string

func getTileType(x int, y int) int {
	tileType := 0

	topCode := fmt.Sprintf("%d:%d", x, y-1)
	rightCode := fmt.Sprintf("%d:%d", x+1, y)
	bottomCode := fmt.Sprintf("%d:%d", x, y+1)
	leftCode := fmt.Sprintf("%d:%d", x-1, y)

	if slices.Contains(visitedTiles, topCode) {
		tileType += 1 << TOP
	}
	if slices.Contains(visitedTiles, rightCode) {
		tileType += 1 << RIGHT
	}
	if slices.Contains(visitedTiles, bottomCode) {
		tileType += 1 << BOTTOM
	}
	if slices.Contains(visitedTiles, leftCode) {
		tileType += 1 << LEFT
	}

	return tileType
}

func fillVertical(start int, end int, x int) {
	for i := start; i <= end; i++ {
		positionCode := fmt.Sprintf("%d:%d", x, i)
		visitedTiles = append(visitedTiles, positionCode)
	}
}

func fillHorizontal(start int, end int, y int) {
	for i := start; i <= end; i++ {
		positionCode := fmt.Sprintf("%d:%d", i, y)
		visitedTiles = append(visitedTiles, positionCode)
	}
}

func Puzzle1(lines []string) int {
	x := 0
	y := 0
	maxX := 0
	minX := 0
	maxY := 0
	minY := 0

	visitedTiles = append(visitedTiles, "0:0")

	// Walk along the lines and save all visited positions
	for _, line := range lines {
		parts := strings.Split(line, " ")
		direction := parts[0]
		distance, _ := strconv.Atoi(parts[1])

		if direction == "U" {
			y -= distance
			fillVertical(y, y+(distance-1), x)
		} else if direction == "R" {
			x += distance
			fillHorizontal(x-(distance-1), x, y)
		} else if direction == "D" {
			y += distance
			fillVertical(y-(distance-1), y, x)
		} else {
			x -= distance
			fillHorizontal(x, x+(distance-1), y)
		}

		minX = min(x, minX)
		minY = min(y, minY)
		maxX = max(x, maxX)
		maxY = max(y, maxY)
	}

	hitCount := 0
	missingTileType := 0
	for y := minY; y <= maxY; y++ {
		isInside := false
		for x := minX; x <= maxX; x++ {
			positionCode := fmt.Sprintf("%d:%d", x, y)
			if slices.Contains(visitedTiles, positionCode) {
				hitCount++

				// Since we have to count the tiles inside the trench, we have to know when we are
				// inside it and when we are not. We switch between being inside and outside whenever we cross
				// a vertical line or horizontal line between to vertical lines in the opposite direction.
				// So, when we find a horizontal line, we have to check whether the vertical lines before and after it
				// either go in opposite directions, because this is when we cross the line. If the vertical lines go
				// into the same direction (e.g. both up or down), we don't cross and therefore don't switch between
				// outside and inside.
				tileType := getTileType(x, y)
				if tileType == VERTICAL || tileType == missingTileType {
					isInside = !isInside
					missingTileType = 0
					continue
				}

				// if we find a tile that marks a vertical line eigher going up or down,
				// we set the missing tile to the opposite one so that we know we crossed the line
				// when we find such a tile next
				if tileType == RIGHT_BOTTOM {
					missingTileType = LEFT_TOP
				} else if tileType == TOP_RIGHT {
					missingTileType = BOTTOM_LEFT
				}
			} else if isInside {
				hitCount++
			}
		}
	}

	return hitCount
}

func main() {
	file, _ := os.Open("input.txt")

	scanner := bufio.NewScanner(file)
	var lines []string

	for scanner.Scan() {
		line := scanner.Text()

		lines = append(lines, line)
	}

	fmt.Printf("Puzzle 1: %d\n", Puzzle1(lines))
}
