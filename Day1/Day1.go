package main;

import "os";
import "bufio";
import "fmt";
import "strings";

func check(error error){

}

func Puzzle1(lines []string) int {
	total := 0;
	for _, line := range lines {
		var onlyNumbers []byte;

		// filter all characters so only numbers remain
		for i := 0; i < len(line); i++ {
			if line[i] >= '0' && line[i] <= '9' {
				onlyNumbers = append(onlyNumbers, line[i] - '0');
			}
		}

		total += int(onlyNumbers[0] * 10 + onlyNumbers[len(onlyNumbers) - 1]);
	}
	return total;
}

func getNumberAtPosition(line string, position int) int {
	digitsAsWords := []string{ "one", "two", "three", "four", "five", "six", "seven", "eight", "nine" };

	// first check if the line starts with an actual digit
	if line[position] >= '0' && line[position] <= '9' {
		return int(line[position] - '0');
	}
	
	// first check if the line starts with one of the numbers
	for index, digit := range digitsAsWords {
		if strings.HasPrefix(line[position:], digit) {
			return index + 1;
		}	
	}

	return -1;
}

func Puzzle2(lines []string) int {
	total := 0;
	
	for _, line := range lines {
		
		firstNumber := -1;
		lastNumber := -1;

		// search for first number from the start
		for i := 0; i < len(line); i++ {
			number := getNumberAtPosition(line, i);
			if (number != -1) {
				firstNumber = number;
				break;
			}
		}

		// search for last number from the back
		for i := len(line) - 1; i >= 0; i-- {
			number := getNumberAtPosition(line, i);
			if (number != -1) {
				lastNumber = number;
				break;
			}
		}

		total += firstNumber * 10 + lastNumber;
	}

	return total;
}

func main(){
	file, _ := os.Open("input.txt");
	defer file.Close();

	scanner := bufio.NewScanner(file);

	var lines []string;
	for scanner.Scan() {
		line := scanner.Text();
		lineAsBytes := line;
		lines = append(lines, lineAsBytes);
	}

	fmt.Printf("Puzzle 1: %d\n", Puzzle1(lines));
	fmt.Printf("Puzzle 2: %d\n", Puzzle2(lines));
}