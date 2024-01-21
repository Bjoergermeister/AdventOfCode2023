package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Rule struct {
	attribute byte
	operator  byte
	value     int
	result    string
}

type Element struct {
	x int
	m int
	a int
	s int
}

var workflows map[string][]Rule = make(map[string][]Rule)

func Puzzle1(elements []Element) int {
	total := 0
	for _, element := range elements {
		workflow := workflows["in"]
		ruleIndex := 0
		for ruleIndex < len(workflow) {
			matches := checkRule(workflow[ruleIndex], element)
			if matches {
				index := min(ruleIndex, len(workflow)-1)
				destination := workflow[index].result
				if destination == "R" {
					break
				}

				if destination == "A" {
					total += element.x + element.m + element.a + element.s
					break
				}

				workflow = workflows[destination]
				ruleIndex = 0
				continue
			}

			ruleIndex++
		}
	}
	return total
}

func checkRule(rule Rule, element Element) bool {
	if rule.attribute == 0 {
		return true
	}

	var value int
	switch rule.attribute {
	case 'x':
		value = element.x
		break
	case 'm':
		value = element.m
		break
	case 'a':
		value = element.a
		break
	case 's':
		value = element.s
	}

	if rule.operator == '<' {
		return value < rule.value
	} else {
		return value > rule.value
	}
}

func main() {
	file, _ := os.Open("input.txt")

	scanner := bufio.NewScanner(file)

	var elements []Element

	isElement := false
	for scanner.Scan() {
		line := scanner.Text()

		if len(line) == 0 {
			isElement = true
			continue
		}

		if isElement {
			parts := strings.Split(line[1:len(line)-1], ",")
			element := Element{
				x: convertToInteger(parts[0][2:]),
				m: convertToInteger(parts[1][2:]),
				a: convertToInteger(parts[2][2:]),
				s: convertToInteger(parts[3][2:]),
			}
			elements = append(elements, element)
		} else {
			divider := strings.Index(line, "{")
			name := line[:divider]
			rules := strings.Split(line[divider+1:len(line)-1], ",")

			var ruleList []Rule
			for i := 0; i < len(rules); i++ {
				var rule Rule
				index := strings.Index(rules[i], ":")
				if index == -1 {
					rule.result = string(rules[i])
				} else {
					rule.attribute = rules[i][0]
					rule.operator = rules[i][1]
					rule.value = convertToInteger(rules[i][2:index])
					rule.result = rules[i][index+1:]

				}
				ruleList = append(ruleList, rule)
			}

			workflows[name] = ruleList
		}
	}

	fmt.Printf("Puzzle 1: %d\n", Puzzle1(elements))
}

func convertToInteger(s string) int {
	value, _ := strconv.Atoi(s)
	return value
}
