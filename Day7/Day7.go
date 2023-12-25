package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

type Hand struct {
	handType  int
	lineIndex int
	bid       int
}

var hands []Hand
var lines []string
var handStrengths = []int{5, 41, 32, 311, 221, 2111, 11111}
var cardStrenghts = []byte{'A', 'K', 'Q', 'J', 'T', '9', '8', '7', '6', '5', '4', '3', '2'}

func getTypeOfHand(hand map[byte]int) int {
	handType := 0
	for i := 5; i > 0; i-- {
		for key := range hand {
			if hand[key] == i {
				handType = handType*10 + i
			}
		}
	}
	return handType
}

func sliceIndex[T byte | int](slice []T, element T) int {
	for i := 0; i < len(slice); i++ {
		if slice[i] == element {
			return i
		}
	}

	return -1
}

func compareFunction(i, j int) bool {
	if hands[i].handType != hands[j].handType {
		return hands[i].handType > hands[j].handType
	}

	for k := 0; k < 5; k++ {
		card1Strength := sliceIndex(cardStrenghts, lines[hands[i].lineIndex][k])
		card2Strength := sliceIndex(cardStrenghts, lines[hands[j].lineIndex][k])
		if card1Strength == card2Strength {
			continue
		}

		return card1Strength > card2Strength
	}

	return true
}

func Puzzle1(lines []string) int {
	for lineIndex, line := range lines {
		hand := Hand{}
		cards := make(map[byte]int)
		for i := 0; i < 5; i++ {
			if _, ok := cards[line[i]]; ok {
				cards[line[i]]++
			} else {
				cards[line[i]] = 1
			}
		}

		bid, _ := strconv.Atoi(line[6:len(line)])

		hand.bid = bid
		hand.lineIndex = lineIndex
		hand.handType = sliceIndex(handStrengths, getTypeOfHand(cards))
		hands = append(hands, hand)

		sort.SliceStable(hands, compareFunction)
	}

	total := 0
	for rank, hand := range hands {
		total += (rank + 1) * hand.bid
	}

	return total
}

func main() {
	file, _ := os.Open("input.txt")

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	fmt.Printf("Puzzle 1: %d\n", Puzzle1(lines))
}
