package main

import (
	"os"
	"bufio"
	"strings"
	"strconv"
	"fmt"
	"math"
)

func getNumberOfDigitsInInteger(number int) int {
	numberOfDigits := 1
	for number != 0 {
		number /= 10
		numberOfDigits++
	}

	return numberOfDigits
}

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

func Puzzle2(times []int, records []int) int {
	// Combine times and records
	time := times[0]
	record := records[0]

	for i := 1; i < len(times); i++ {
		time = (time * int(math.Pow(10, float64(getNumberOfDigitsInInteger(times[i]) - 1)))) + times[i]
		record = (record * int(math.Pow(10, float64(getNumberOfDigitsInInteger(records[i]) - 1)))) + records[i]
	}
	
	// find lower and upper bounds
	lowerEnd := 0
	for loadingTime := 0; loadingTime <= time; loadingTime++ {
		if (loadingTime * (time - loadingTime) > record){
			lowerEnd = loadingTime
			break
		}
	}

	upperEnd := 0
	for loadingTime := time; loadingTime >= 0; loadingTime-- {
		if (loadingTime * (time - loadingTime) > record){
			upperEnd = loadingTime
			break
		}
	}

	return (upperEnd - lowerEnd) + 1
}

func main(){
	file, _ := os.Open("input.txt")

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	games := lineToIntegerList(scanner.Text())
	scanner.Scan()
	records := lineToIntegerList(scanner.Text())

	fmt.Printf("Puzzle 1: %d\n", Puzzle1(games, records))
	fmt.Printf("Puzzle 1: %d\n", Puzzle2(games, records))
}