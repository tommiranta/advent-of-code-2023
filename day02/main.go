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
func parseGameNumber(text string) (int, string) {
	parts := strings.Split(text, ":")
	game_nr, err := strconv.Atoi(strings.Split(parts[0], " ")[1])
	if err != nil {
		game_nr = -1
	}

	return game_nr, parts[1]
}
func processLine(text string) int {
	game_nr, game_content := parseGameNumber(text)
	if hasValidGames(game_content) {
		return game_nr
	}
	return 0
}

func main() {
	sum := 0

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		res := processLine(scanner.Text())
		sum += res
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Sum of game numbers with valid games", sum)
}
