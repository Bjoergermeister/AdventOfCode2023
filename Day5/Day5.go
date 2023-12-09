package main

import (
	"os"
	"bufio"
	"strings"
	"strconv"
	"fmt"
)

type Mapping struct {
	sourceStart int
	destinationStart int
	length int
}

var seeds []int
var seedToSoil []Mapping
var soilToFertilizer []Mapping
var fertilizerToWater []Mapping
var waterToLight []Mapping
var lightToTemperature []Mapping
var temperatureToHumidity []Mapping
var humidityToLocation []Mapping

func parseMaps(lines []string) {

	// Parse soils
	seedNumbers := strings.Split(lines[0][7:], " ")
	for _, seedNumber := range seedNumbers {
		number, _ := strconv.Atoi(seedNumber)
		seeds = append(seeds, number)
	}

	// Parse the maps
	var currentSection string
	for i := 1; i < len(lines); i++ {
		line := lines[i]
		if len(line) == 0 {
			continue
		}

		if line[0] < '0' || line[0] > '9' {
			currentSection = line
			continue
		}

		parts := strings.Split(line, " ")
		destinationStart, _ := strconv.Atoi(parts[0])
		sourceStart, _ := strconv.Atoi(parts[1])
		length, _ := strconv.Atoi(parts[2])
		mapping := Mapping{ sourceStart, destinationStart, length }
		
		switch (currentSection){
			case "seed-to-soil map:":
				seedToSoil = append(seedToSoil, mapping)
			case "soil-to-fertilizer map:":
				soilToFertilizer = append(soilToFertilizer, mapping)
			case "fertilizer-to-water map:":
				fertilizerToWater = append(fertilizerToWater, mapping)
			case "water-to-light map:":
				waterToLight = append(waterToLight, mapping)
			case  "light-to-temperature map:":
				lightToTemperature = append(lightToTemperature, mapping)
			case  "temperature-to-humidity map:":
				temperatureToHumidity = append(temperatureToHumidity, mapping)
			case "humidity-to-location map:":
				humidityToLocation = append(humidityToLocation, mapping)
		}
	}
}

func mapTo(mappings []Mapping, source int) int {
	for _, mapping := range mappings {
		if source >= mapping.sourceStart && source <= mapping.sourceStart + mapping.length {
			offset := source - mapping.sourceStart
			return mapping.destinationStart + offset
		}
	}

	return source
}

func Puzzle1() int {
	lowestPosition := -1
	for _, seed := range seeds {
		soil := mapTo(seedToSoil, seed)
		fertilizer := mapTo(soilToFertilizer, soil)
		water := mapTo(fertilizerToWater, fertilizer)
		light := mapTo(waterToLight, water)
		temperature := mapTo(lightToTemperature, light)
		humidity := mapTo(temperatureToHumidity, temperature)
		location := mapTo(humidityToLocation, humidity)
		
		if lowestPosition == -1 || location < lowestPosition {
			lowestPosition = location
		} 
	}

	return lowestPosition
}

func main(){
	file, _ := os.Open("input.txt")

	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	parseMaps(lines)

	fmt.Printf("Puzzle 1: %d\n", Puzzle1())
}