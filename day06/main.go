package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func calculateDistance(load_duration int, race_duration int) (int, error) {
	if load_duration >= race_duration {
		return 0, errors.New("Too long load")
	}
	if load_duration == 0 {
		return 0, errors.New("No load")
	}
	return (race_duration - load_duration) * load_duration, nil
}

func findNumBetterStrategies(race_duration int, record int) int {
	ret := 0
	i := 1
	for i < race_duration {
		distance, err := calculateDistance(i, race_duration)
		if err == nil && distance > record {
			ret += 1
		}
		i += 1
	}
	return ret
}

func convertList(lst []string) []int {
	ret := []int{}
	for _, s := range lst {
		if s != "" {
			i, err := strconv.Atoi(s)
			if err != nil {
				panic(err)
			}
			ret = append(ret, i)
		}
	}
	return ret
}

func parseLine(text string) []int {
	startFrom := strings.Index(text, ":") + 1
	numbers := strings.Split(strings.Trim(text[startFrom:], " "), " ")
	return convertList(numbers)
}

func parseLineKerning(text string) int {
	startFrom := strings.Index(text, ":") + 1
	numbers := strings.Split(strings.Trim(text[startFrom:], " "), " ")
	oneBigNum := strings.Join(numbers, "")
	ret, err := strconv.Atoi(oneBigNum)
	if err != nil {
		panic(err)
	}
	return ret
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

func partOne(lines []string) int {
	result := 1
	times := parseLine(lines[0])
	records := parseLine(lines[1])
	if len(times) != len(records) {
		panic("You need the same amount of race times and record distances.")
	}
	i := 0
	for i < len(times) {
		num_winning_strategies := findNumBetterStrategies(times[i], records[i])
		if num_winning_strategies > 0 {
			result *= num_winning_strategies
		}
		i += 1
	}
	return result
}

func partTwo(lines []string) int {
	time := parseLineKerning(lines[0])
	record := parseLineKerning(lines[1])
	num_winning_strategies := findNumBetterStrategies(time, record)
	return num_winning_strategies
}

func main() {
	lines := readFile("input.txt")
	result := partOne(lines)
	fmt.Println("Result of part one:", result)
	result = partTwo(lines)
	fmt.Println("Result of part two:", result)
}
