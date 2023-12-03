package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func hasValidColorCount(reveal string) bool {
	max_counts := map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}
	reveal_parts := strings.Split(strings.Trim(reveal, " "), " ")
	count, err := strconv.Atoi(reveal_parts[0])
	if err != nil {
		return false
	}
	if count > max_counts[reveal_parts[1]] {
		return false
	}
	return true

}

func hasValidGames(text string) bool {
	games := strings.Split(text, ";")
	for _, game := range games {
		reveals := strings.Split(game, ",")
		for _, reveal := range reveals {
			isValid := hasValidColorCount(reveal)
			if !isValid {
				return false
			}
		}
	}
	return true

}

func parseReveal(reveal string) (string, int, error) {
	reveal_parts := strings.Split(strings.Trim(reveal, " "), " ")
	count, err := strconv.Atoi(reveal_parts[0])
	if err != nil {
		return "invalid", -1, err
	}
	return reveal_parts[1], count, nil

}

func calculatePower(cubes map[string]int) int {
	result := 1
	for _, v := range cubes {
		result *= v
	}
	return result
}

func processGame(game string) int {
	min_counts := make(map[string]int)
	rounds := strings.Split(game, ";")
	for _, round := range rounds {
		reveals := strings.Split(round, ",")
		for _, reveal := range reveals {
			color, count, err := parseReveal(reveal)
			if err != nil {
				continue
			}
			if count > min_counts[color] {
				min_counts[color] = count
			}
		}
	}
	return calculatePower(min_counts)

}

func parseGameNumber(text string) (int, string) {
	parts := strings.Split(text, ":")
	game_nr, err := strconv.Atoi(strings.Split(parts[0], " ")[1])
	if err != nil {
		game_nr = -1
	}

	return game_nr, parts[1]
}

func processLine(text string) (int, int) {
	game_valid_score := 0
	game_nr, game_content := parseGameNumber(text)
	if hasValidGames(game_content) {
		game_valid_score = game_nr
	}
	game_power := processGame(game_content)
	return game_valid_score, game_power
}

func main() {
	valid_sum := 0
	power_sum := 0

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		game_valid_score, game_power := processLine(scanner.Text())
		valid_sum += game_valid_score
		power_sum += game_power
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Sum of game numbers with valid games", valid_sum)
	fmt.Println("Sum of game powers", power_sum)
}
