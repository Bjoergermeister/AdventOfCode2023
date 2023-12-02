package main;

import "os";
import "bufio";
import "fmt";

func check(error error){

}

func Puzzle1(lines [][]byte) int {
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

func main(){
	file, _ := os.Open("input.txt");
	defer file.Close();

	scanner := bufio.NewScanner(file);

	var lines [][]byte;
	for scanner.Scan() {
		line := scanner.Text();
		lineAsBytes := []byte(line);
		lines = append(lines, lineAsBytes);
	}

	fmt.Printf("Puzzle 1: %d", Puzzle1(lines));
}