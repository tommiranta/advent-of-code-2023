package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

const UNICODE_ZERO = 48
const UNICODE_NINE = UNICODE_ZERO + 9

var NUMBERS_AS_LETTERS = [10]string{
	"zero",
	"one",
	"two",
	"three",
	"four",
	"five",
	"six",
	"seven",
	"eight",
	"nine",
}

func findFirstDigit(text string) (int, int) {
	for idx, ch := range text {
		if ch >= UNICODE_ZERO && ch <= UNICODE_NINE {
			return int(ch - UNICODE_ZERO), idx
		}
	}
	return -1, -1
}

func findLastDigit(text string) (int, int) {
	for i := len(text) - 1; i >= 0; i-- {
		ch := text[i]
		if ch >= UNICODE_ZERO && ch <= UNICODE_NINE {
			return int(ch - UNICODE_ZERO), i
		}
	}
	return -1, -1
}

func findFirstDigitAsLetters(text string) (int, int) {
	digit_winning := -1
	idx_winning := -1

	for value, letters := range NUMBERS_AS_LETTERS {
		idx := strings.Index(text, letters)
		if idx != -1 && (idx_winning == -1 || idx < idx_winning) {
			digit_winning = value
			idx_winning = idx
		}
	}
	return digit_winning, idx_winning
}

func findLastDigitAsLetters(text string) (int, int) {
	digit_winning := -1
	idx_winning := -1

	for value, letters := range NUMBERS_AS_LETTERS {
		idx := strings.LastIndex(text, letters)
		if idx != -1 && (idx_winning == -1 || idx > idx_winning) {
			digit_winning = value
			idx_winning = idx
		}
	}
	return digit_winning, idx_winning
}

func digitsToNumber(double_digit int, single_digit int) int {
	return double_digit*10 + single_digit
}

func parseLineNoLetters(text string) int {
	double_digit, _ := findFirstDigit(text)
	single_digit, _ := findLastDigit(text)
	if single_digit != -1 && double_digit != -1 {
		return digitsToNumber(double_digit, single_digit)
	}
	return -1
}

func parseLine(text string) int {
	single_digit_winner := -1
	double_digit_winner := -1

	double_digit, idx_double_digit := findFirstDigit(text)
	single_digit, idx_single_digit := findLastDigit(text)

	double_digit_letters, idx_double_digit_letters := findFirstDigitAsLetters(text)
	single_digit_letters, idx_single_digit_letters := findLastDigitAsLetters(text)

	if idx_double_digit_letters != -1 && idx_double_digit_letters < idx_double_digit {
		double_digit_winner = double_digit_letters
	} else {
		double_digit_winner = double_digit
	}

	if idx_single_digit_letters != -1 && idx_single_digit_letters > idx_single_digit {
		single_digit_winner = single_digit_letters
	} else {
		single_digit_winner = single_digit
	}

	if single_digit_winner != -1 && double_digit_winner != -1 {
		return digitsToNumber(double_digit_winner, single_digit_winner)
	}
	return -1
}

func main() {
	sum_pt1 := 0
	sum_pt2 := 0

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		number := parseLineNoLetters(scanner.Text())
		if number == -1 {
			log.Fatal("Parsing digits failed")
			break
		}
		sum_pt1 += number

		number = parseLine(scanner.Text())
		if number == -1 {
			log.Fatal("Parsing digits failed")
			break
		}
		sum_pt2 += number
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Solution part 1:", sum_pt1)
	fmt.Println("Solution part 2:", sum_pt2)
}
