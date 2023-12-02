package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
)

func parseLine(text string) (int, error) {
	const zero = 48
	const nine = zero + 9
	var a, b int

	for _, ch := range text {
		if ch >= zero && ch <= nine {
			if a == 0 {
				a = int(ch)
				b = int(ch)
			} else {
				b = int(ch)
			}
		}
	}

	if a > 0 && b > 0 {
		return (a-zero)*10 + (b - zero), nil
	} else {
		return -1, errors.New("Did not find digits in string.")
	}

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
		number, err := parseLine(scanner.Text())
		if err != nil {
			break
		}
		sum += number
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(sum)

}
