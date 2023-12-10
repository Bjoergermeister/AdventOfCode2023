package main

import (
	"os"
	"bufio"
	"strings"
	"strconv"
	"fmt"
)

func lineToIntegerList(line string) []int {
	var numbers []int
	line = strings.Split(line, ":")[1]
	for _, numberAsString := range strings.Split(line, " "){
		number, error := strconv.Atoi(numberAsString)
		if error == nil {
			numbers = append(numbers, number)
		}
	}

	return numbers
}

func Puzzle1(times []int, records []int) int {
	answer := 1
	for timeIndex, time := range times {
		numberOfWinningsWays := 0
		for loadingTime := 0; loadingTime <= time; loadingTime++ {
			if (loadingTime * (time - loadingTime) > records[timeIndex]){
				numberOfWinningsWays++
			}
		}
		answer *= numberOfWinningsWays
	}

	return answer
}

func main(){
	file, _ := os.Open("input.txt")

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	games := lineToIntegerList(scanner.Text())
	scanner.Scan()
	records := lineToIntegerList(scanner.Text())

	fmt.Printf("Puzzle 1: %d", Puzzle1(games, records))
}