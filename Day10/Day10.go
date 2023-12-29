package main

import (
	"bufio"
	"fmt"
	"os"
)

const TOP = 1
const BOTTOM = 2
const LEFT = 3
const RIGHT = 4

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

func Puzzle1(lines []string) int {
	// Find starting position
	var x, y int
	for i, line := range lines {
		for j := 0; j < len(line); j++ {
			if line[j] == 'S' {
				y = i
				x = j
			}
		}
	}

	// Walk grid
	x, y, direction := findFirstTileAndDirection(lines, x, y)
	tile := lines[y][x]

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
		length++
	}

	// For a single tile to be the farthest, the path must contain an odd number of tiles.
	// Then the answer is just the median in a list from 1 to PATH_LENGTH
	return ((length - 1) / 2) + 1
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
