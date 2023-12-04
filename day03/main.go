package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

const UNICODE_ZERO = byte('0')
const UNICODE_NINE = byte('9')
const UNICODE_DOT = byte('.')

type position struct {
	x int
	y int
}

type partnum struct {
	x1 int
	x2 int
	y  int
}

func isCharNumeric(chr byte) bool {
	if chr >= UNICODE_ZERO && chr <= UNICODE_NINE {
		return true
	}
	return false
}
func isStartOfNumber(chr byte, part partnum) bool {
	if isCharNumeric(chr) && part.x1 == -1 {
		return true
	}
	return false
}

func isEndOfNumber(chr byte, part partnum) bool {
	if !isCharNumeric(chr) && part.x1 != -1 {
		return true
	}
	return false
}

func hasStartEndIdx(part partnum) bool {
	if part.x1 != -1 && part.x2 != -1 {
		return true
	}
	return false
}

func isLastCharOnLine(x int, line string) bool {
	if x == len(line)-1 {
		return true
	}
	return false
}

func getNewPosition(x int, y int, lines []string) (int, int) {
	if isLastCharOnLine(x, lines[y]) {
		return 0, y + 1
	}
	return x + 1, y
}

func getVerificationArea(part partnum, lines []string) (position, position) {
	x_min := part.x1 - 1
	if x_min < 0 {
		x_min = 0
	}
	x_max := part.x2
	y_min := part.y - 1
	if y_min < 0 {
		y_min = 0
	}
	y_max := part.y + 1
	if y_max >= len(lines) {
		y_max = len(lines) - 1
	}

	return position{x: x_min, y: y_min}, position{x: x_max, y: y_max}
}
func isInsidePartNum(part partnum, x int, y int, last_char bool) bool {
	if y == part.y && x >= part.x1 && x < part.x2 {
		return true
	}
	if last_char && y == part.y && x >= part.x1 && x <= part.x2 {
		return true
	}
	return false
}

func isAdjacentToSymbol(part partnum, lines []string) bool {
	start, end := getVerificationArea(part, lines)
	for y := start.y; y <= end.y; y++ {
		for x := start.x; x <= end.x; x++ {
			if !isInsidePartNum(part, x, y, isLastCharOnLine(x, lines[y])) &&
				lines[y][x] != UNICODE_DOT {
				return true
			}

		}
	}
	return false
}

func findNext(pos *position, lines []string) int {
	for y := pos.y; y < len(lines); y++ {
		part := partnum{y: y, x1: -1, x2: -1}
		for x := pos.x; x < len(lines[y]); x++ {
			if isStartOfNumber(lines[y][x], part) {
				part.x1 = x
			}
			if isEndOfNumber(lines[y][x], part) || isLastCharOnLine(x, lines[y]) {
				part.x2 = x
			}

			if hasStartEndIdx(part) {
				pos.x, pos.y = getNewPosition(x, y, lines)
				part_string := lines[part.y][part.x1:part.x2]
				if isCharNumeric(lines[y][part.x2]) {
					part_string = lines[part.y][part.x1:]
				}
				res, err := strconv.Atoi(part_string)
				if err != nil {
					fmt.Printf("Error parsing '%s'\n", part_string)
					return -1
				}
				if part_string == "936" {
					one, two := getVerificationArea(part, lines)
					fmt.Println(
						"pos:",
						part,
						isAdjacentToSymbol(part, lines),
					)
					for y := one.y; y <= two.y; y++ {
						fmt.Println(lines[y][one.x:two.x])
					}
				}
				if isAdjacentToSymbol(part, lines) {
					return res
				}
				// fmt.Println("skipped:", part_string, "pos:", part)
				part.x1 = -1
				part.x2 = -1
			}
		}
		pos.x = 0
	}
	return -1
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
	cur := position{x: 0, y: 0}
	sum := 0

	for {
		ret := findNext(&cur, lines)
		if ret == -1 {
			break
		}
		// fmt.Println(ret)
		sum += ret
	}
	fmt.Println(sum)
}
