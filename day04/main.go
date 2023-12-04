package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func calculateScore(num_correct int) int {
	if num_correct < 2 {
		return num_correct
	}
	return int(math.Pow(2.0, float64(num_correct)-1.0))
}
func checkCorrectAnswers(winning_numbers map[string]int, drawn_numbers map[string]int) int {
	num_correct := 0
	for k, _ := range drawn_numbers {
		num_correct += winning_numbers[k]
	}
	return num_correct
}

func numbersStringToMap(numbers string) map[string]int {
	split := strings.Split(numbers, " ")
	result := make(map[string]int)
	for _, v := range split {
		if v != "" {
			result[v] = 1
		}
	}
	return result

}

func parseLine(text string) (int, map[string]int, map[string]int) {
	parts := strings.Split(text, ": ")
	game_nr, err := strconv.Atoi(strings.Split(parts[0], " ")[1])
	if err != nil {
		game_nr = -1
	}

	numbers := strings.Split(parts[1], " | ")
	winning_numbers := numbersStringToMap(numbers[0])
	drawn_numbers := numbersStringToMap(numbers[1])

	return game_nr, winning_numbers, drawn_numbers
}

func readFile(filename string) []string {
	lines := make([]string, 0, 1)
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

func main() {
	lines := readFile("input.txt")
	sum := 0

	for _, l := range lines {
		_, winning_numbers, drawn_numbers := parseLine(l)
		num_correct := checkCorrectAnswers(winning_numbers, drawn_numbers)
		score := calculateScore(num_correct)
		sum += score
	}
	fmt.Println("Scorecard points", sum)
}
