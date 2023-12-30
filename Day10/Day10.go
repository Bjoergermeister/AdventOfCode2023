package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

var gridWidth int = 0

const TOP = 1
const BOTTOM = 2
const LEFT = 3
const RIGHT = 4

var CORRESPONDING_CORNER_TILES = map[byte]byte{
	'F': 'J',
	'L': '7',
}

var OPPOSITE_CORNER_TILES = map[byte]byte{
	'F': '7',
	'L': 'J',
}

var pathTiles []int

func findFirstTileAndDirection(lines []string, x int, y int) (int, int, int) {
	top := lines[y-1][x]
	bottom := lines[y+1][x]
	left := lines[y][x-1]
	right := lines[y][x+1]

	if top == 'F' || top == '|' || top == '7' {
		return x, y - 1, TOP
	}

	if bottom == 'L' || bottom == '|' || bottom == 'J' {
		return x, y + 1, BOTTOM
	}

	if left == '-' || left == 'L' || left == 'F' {
		return x - 1, y, LEFT
	}

	if right == '-' || right == '7' || right == 'J' {
		return x + 1, y, RIGHT
	}

	panic("No valid continuation found")
}

func replaceStartWithCorrectTile() byte {
	firstTile := pathTiles[0]
	lastTile := pathTiles[len(pathTiles)-2] // last tile in list is starting tile

	firstX := firstTile % gridWidth
	firstY := firstTile / gridWidth

	lastX := lastTile % gridWidth
	lastY := lastTile / gridWidth

	if firstX == lastX {
		return '|'
	}

	if firstY == lastY {
		return '-'
	}

	if (lastX > firstX && lastY < firstY) || (lastX < firstX && lastY > firstY) {
		return 'F'
	}
	if (lastX > firstX && lastY > firstY) || (lastX < firstX && lastY < firstY) {
		return 'L'
	}
	if (lastX < firstX && lastY < firstY) || (lastX > firstX && lastY > firstY) {
		return '7'
	}
	if (lastX < firstX && lastY > firstY) || (lastX > firstX && lastY < firstY) {
		return 'J'
	}

	return 'X'
}

func Puzzle1(lines []string) int {
	// Find starting position
	var x, y int
	var startY int
	for i, line := range lines {
		for j := 0; j < len(line); j++ {
			if line[j] == 'S' {
				y, startY = i, i
				x = j
			}
		}
	}

	// Walk grid
	x, y, direction := findFirstTileAndDirection(lines, x, y)
	tile := lines[y][x]
	pathTiles = append(pathTiles, y*gridWidth+x)

	length := 0
	for tile != 'S' {
		// |
		if tile == '|' && direction == TOP {
			y -= 1
		}
		if tile == '|' && direction == BOTTOM {
			y += 1
		}

		// -
		if tile == '-' && direction == LEFT {
			x -= 1
		}
		if tile == '-' && direction == RIGHT {
			x += 1
		}

		// L
		if tile == 'L' && direction == LEFT {
			y -= 1
			direction = TOP
		}
		if tile == 'L' && direction == BOTTOM {
			x += 1
			direction = RIGHT
		}

		// J
		if tile == 'J' && direction == BOTTOM {
			x -= 1
			direction = LEFT
		}
		if tile == 'J' && direction == RIGHT {
			y -= 1
			direction = TOP
		}

		// 7
		if tile == '7' && direction == RIGHT {
			y += 1
			direction = BOTTOM
		}
		if tile == '7' && direction == TOP {
			x -= 1
			direction = LEFT
		}

		// F
		if tile == 'F' && direction == LEFT {
			y += 1
			direction = BOTTOM
		}
		if tile == 'F' && direction == TOP {
			x += 1
			direction = RIGHT
		}

		tile = lines[y][x]
		pathTiles = append(pathTiles, y*gridWidth+x)
		length++
	}

	// Replace the 'S' with a letter for the correct tile
	correctTile := replaceStartWithCorrectTile()
	lines[startY] = strings.Replace(lines[startY], "S", string(correctTile), 1)

	// For a single tile to be the farthest, the path must contain an odd number of tiles.
	// Then the answer is just the median in a list from 1 to PATH_LENGTH
	return ((length - 1) / 2) + 1
}

func Puzzle2(lines []string) int {
	containedTiles := 0
	for y, line := range lines {
		isInside := false

		// Every time we hit a tile that is part of the path, our inside/outside state switches.
		// If we hit '|', we can simply flip the current state
		// If we hit a corner tile ('L', '7', 'J' or 'F'), we will "walk" alongside a wall for some time
		// To determine whether we crossed an inside/outside wall, we need to either find the corresponding corner
		// tile so that we know we fully crossed the wall (e.g. a 'F' followed by a 'J') or the opposite so we know
		// we did NOT cross the wall completely (e.g a 'F' followed by a '7').
		// The needed tile variable contains the tile we need to find to make sure we fully crossed a wall
		var neededTile byte = '0'

		for x := 0; x < len(line); x++ {
			position := y*gridWidth + x

			if slices.Contains(pathTiles, position) == false && isInside {
				containedTiles++
			}

			// First check if the current tile if a part of the path.
			// Then figure out whether we step inside or outside
			if slices.Contains(pathTiles, position) {
				if line[x] == '-' {
					continue
				}

				if line[x] == '|' {
					isInside = !isInside
					continue
				}

				if neededTile == '0' {
					neededTile = line[x]
					continue
				}

				if neededTile != '0' {
					if line[x] == CORRESPONDING_CORNER_TILES[neededTile] {
						isInside = !isInside
					}
					neededTile = '0'
				}
			}
		}
	}

	return containedTiles
}

func main() {

	file, _ := os.Open("input.txt")

	scanner := bufio.NewScanner(file)

	var lines []string
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	gridWidth = len(lines[0])

	fmt.Printf("Puzzle 1: %d\n", Puzzle1(lines))
	fmt.Printf("Puzzle 2: %d\n", Puzzle2(lines))
}
