package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var maxRedCubes int = 12
var maxGreenCubes int = 13
var maxBlueCubes int = 14

type CubeSet struct {
	green int
	red   int
	blue  int
}

func parseSets(sets []string) []CubeSet {
	var cubeSets []CubeSet
	for _, set := range sets {
		cubes := strings.Split(set, ",")
		var cubeSet CubeSet
		for _, cube := range cubes {
			cube = strings.TrimSpace(cube)
			divider := strings.Index(cube, " ")
			amount, _ := strconv.Atoi(cube[0:divider])
			color := cube[divider+1:]

			switch color {
			case "red":
				cubeSet.red = amount
			case "green":
				cubeSet.green = amount
			case "blue":
				cubeSet.blue = amount
			}
		}
		cubeSets = append(cubeSets, cubeSet)
	}
	return cubeSets
}

func Puzzle1(games [][]CubeSet) int {
	total := 0
	for gameIndex, game := range games {
		gameIsValid := true
		for _, set := range game {
			if set.red > maxRedCubes || set.green > maxGreenCubes || set.blue > maxBlueCubes {
				gameIsValid = false
			}
		}

		if gameIsValid {
			total += (gameIndex + 1)
		}
	}
	return total
}
func Puzzle2(games [][]CubeSet) int {
	total := 0
	for _, game := range games {
		lowestRedCubes := 0
		lowestGreenCubes := 0
		lowestBlueCubes := 0

		for _, set := range game {
			lowestBlueCubes = max(lowestBlueCubes, set.blue)
			lowestRedCubes = max(lowestRedCubes, set.red)
			lowestGreenCubes = max(lowestGreenCubes, set.green)
		}

		total += lowestBlueCubes * lowestGreenCubes * lowestRedCubes
	}
	return total
}

func main() {
	file, _ := os.Open("input.txt")

	var games [][]CubeSet
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		sets := strings.Split(line, ":")[1]
		games = append(games, parseSets(strings.Split(sets, ";")))
	}

	fmt.Printf("Puzzle 1: %d\n", Puzzle1(games))
	fmt.Printf("Puzzle 1: %d\n", Puzzle2(games))
}
