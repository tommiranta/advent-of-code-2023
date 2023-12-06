package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

const START_VALUE int = math.MaxInt32

type mapping struct {
	destination int
	source      int
	rng         int
}

type spec struct {
	name     string
	mappings []mapping
}

func convertList(lst []string) []int {
	ret := []int{}
	for _, s := range lst {
		i, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}
		ret = append(ret, i)
	}
	return ret
}

func readMapping(lines []string) spec {
	name := strings.Split(lines[0], " ")[0]
	mappings := []mapping{}
	for _, l := range lines[1:] {
		line_split := strings.Split(l, " ")
		m := convertList(line_split)
		mappings = append(mappings, mapping{destination: m[0], source: m[1], rng: m[2]})
	}
	return spec{name: name, mappings: mappings}

}
func parseLines(lines []string) ([]int, []spec) {
	maps := []spec{}
	seeds := []int{}

	line_split := strings.Split(lines[0], " ")[1:]
	seeds = convertList(line_split)
	startIdx := 2
	for i, l := range lines {
		if (l == "" || i == len(lines)-1) && i > startIdx {
			maps = append(maps, readMapping(lines[startIdx:i]))
			startIdx = i + 1
		}
	}
	return seeds, maps
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

func calculateDestination(src int, m mapping) (int, error) {
	if src >= m.source && src < m.source+m.rng {
		idx := src - m.source
		return m.destination + idx, nil
	}
	return -1, errors.New("Not found")
}

func findBest(src int, mappings []mapping) int {
	winning := START_VALUE
	for _, m := range mappings {
		dst, err := calculateDestination(src, m)
		if err == nil && dst < winning {
			winning = dst
		}
	}
	if winning == START_VALUE {
		winning = src
	}

	return winning
}
func main() {
	lines := readFile("input.txt")
	seeds, maps := parseLines(lines)
	winning_seed := START_VALUE
	winning_location := START_VALUE
	for _, seed := range seeds {
		cur := seed
		for _, s := range maps {
			cur = findBest(cur, s.mappings)
		}
		fmt.Println(seed, cur)
		if cur < winning_location {
			winning_location = cur
			winning_seed = seed
		}
	}
	fmt.Println("Winning seed:", winning_seed, "winning_location:", winning_location)
}
