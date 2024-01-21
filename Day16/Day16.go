package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

var RIGHT int = 0
var BOTTOM int = 1
var LEFT int = 2
var TOP int = 3

var GRID_WIDTH int
var GRID_HEIGHT int

type Beam struct {
	position  int
	direction int // 0 = Right, 1 = Bottom, 2 = Left, 3 = Top
}

func Puzzle1(lines []string) int {
	beams := make([]Beam, 0, 100)
	state := make(map[int][]int)

	beams = append(beams, Beam{
		position:  0,
		direction: RIGHT,
	})

	for len(beams) > 0 {
		for i := 0; i < len(beams); i++ {
			beam := &beams[i]
			x := beam.position % GRID_HEIGHT
			y := beam.position / GRID_WIDTH

			tile := lines[y][x]

			if tile == '/' || tile == '\\' {
				changeDirection(beam, tile)
			}

			if shouldSplit(beam, tile) {
				newBeam := split(beam, tile)

				isOutOfGrid := move(&newBeam, newBeam.position%GRID_WIDTH, newBeam.position/GRID_WIDTH)
				if isOutOfGrid == false {
					beams = append(beams, newBeam)
					state[newBeam.position] = append(state[newBeam.position], newBeam.direction)
				}
			}

			isOutOfGrid := move(beam, x, y)
			if isOutOfGrid || slices.Contains(state[beam.position], beam.direction) {
				for i := 0; i < len(beams); i++ {
					if beam.position == beams[i].position && beam.direction == beams[i].direction {
						beams[i] = beams[len(beams)-1]
						beams = beams[:len(beams)-1]
					}
				}
				continue
			}

			state[beam.position] = append(state[beam.position], beam.direction)
		}
	}

	return len(state) + 1
}

func shouldSplit(beam *Beam, tile byte) bool {
	vertical := tile == '|' && (beam.direction == LEFT || beam.direction == RIGHT)
	horizontal := tile == '-' && (beam.direction == TOP || beam.direction == BOTTOM)
	return vertical || horizontal
}

func move(beam *Beam, x int, y int) bool {
	if beam.direction == RIGHT {
		x++
	} else if beam.direction == BOTTOM {
		y++
	} else if beam.direction == LEFT {
		x--
	} else {
		y--
	}

	beam.position = y*GRID_WIDTH + x

	return x < 0 || y < 0 || x >= GRID_WIDTH || y >= GRID_HEIGHT
}

func changeDirection(beam *Beam, tile byte) {
	if (beam.direction == RIGHT && tile == '/') || (beam.direction == LEFT && tile == '\\') {
		beam.direction = TOP
		return
	}
	if (beam.direction == BOTTOM && tile == '/') || (beam.direction == TOP && tile == '\\') {
		beam.direction = LEFT
		return
	}
	if (beam.direction == LEFT && tile == '/') || (beam.direction == RIGHT && tile == '\\') {
		beam.direction = BOTTOM
		return
	}
	if (beam.direction == TOP && tile == '/') || (beam.direction == BOTTOM && tile == '\\') {
		beam.direction = RIGHT
	}
}

func split(beam *Beam, tile byte) Beam {
	newBeam := Beam{
		direction: beam.direction,
		position:  beam.position,
	}

	if tile == '-' {
		beam.direction = LEFT
		newBeam.direction = RIGHT
	} else {
		beam.direction = TOP
		newBeam.direction = BOTTOM
	}

	return newBeam
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
	GRID_HEIGHT = len(lines)

	fmt.Printf("Puzzle 1: %d\n", Puzzle1(lines))
}
