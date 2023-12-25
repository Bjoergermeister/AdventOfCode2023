package main

import (
	"bufio"
	"fmt"
	"os"
)

func hash(code string) int {
	var hash int = 0
	for i := 0; i < len(code); i++ {
		hash *= 100
		hash += int(code[i] - '@')
	}
	return hash
}

func Puzzle1(lines []string) int {
	directions := make(map[int][]int, len(lines))
	for _, line := range lines[2:] {
		source := hash(line[0:3])
		left := hash(line[7:10])
		right := hash(line[12 : len(line)-1])
		directions[source] = []int{left, right}
	}

	end := hash("ZZZ")
	current := hash("AAA")
	instructions := lines[0]

	index := 0
	for current != end {
		instruction := instructions[index%len(instructions)]
		if instruction == 'L' {
			current = directions[current][0]
		} else {
			current = directions[current][1]
		}

		index++
	}

	return index
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
