package main;

import "os";
import "fmt";
import "bufio";
import "strings";
import "strconv";

var maxRedCubes int = 12;
var maxGreenCubes int = 13;
var maxBlueCubes int = 14;

type CubeSet struct {
	green int;
	red int;
	blue int;
}

func parseSets(sets []string) []CubeSet {
	var cubeSets []CubeSet;
	for _, set := range sets {
		cubes := strings.Split(set, ",");
		var cubeSet CubeSet;
		for _, cube := range cubes {
			cube = strings.TrimSpace(cube);
			divider := strings.Index(cube, " ");
			amount, _ := strconv.Atoi(cube[0:divider]);
			color := cube[divider + 1:];

			switch color {
			case "red":
				cubeSet.red = amount;
			case "green":
				cubeSet.green = amount;
			case "blue":
				cubeSet.blue = amount;
			}
		}
		cubeSets = append(cubeSets, cubeSet);
	}
	return cubeSets;
}

func Puzzle1(lines []string) int {
	total := 0;
	for gameIndex, line := range lines {
		parts := strings.Split(line, ":");

		sets := parseSets(strings.Split(parts[1], ";"));

		gameIsValid := true;
		for _, set := range sets {
			if set.red > maxRedCubes || set.green > maxGreenCubes || set.blue > maxBlueCubes {
				gameIsValid = false;
			}
		}

		if (gameIsValid){
			total += (gameIndex + 1);
		}
	}
	return total;
}

func main(){
	file, _ := os.Open("input.txt");

	var lines []string;
	scanner := bufio.NewScanner(file);
	for scanner.Scan(){
		line := scanner.Text();
		lines = append(lines, line);
	}

	fmt.Printf("Puzzle 1: %d\n", Puzzle1(lines));
}