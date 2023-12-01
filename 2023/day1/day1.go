package main

import (
	"errors"
	"fmt"
	"strings"
	"unicode"

	"stoykotolev/aoc-2023/utils"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

var spelledNumbers = []string{
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

func main() {

	inputContents := utils.ReadFile("./input.txt")
	part1(inputContents)
	part2(inputContents)

}

func part1(contents []string) {

	total := 0

	for _, line := range contents {
		var digits []int
		for _, char := range line {
			if unicode.IsDigit(char) {
				digits = append(digits, int(char-'0'))
			}
		}
		total = total + (digits[0]*10 + digits[len(digits)-1])
	}
	fmt.Printf("Part 1 %d\n", total)
}

func part2(contents []string) {
	total := 0
	for _, line := range contents {
		first, err := findFirst(line)
		check(err)
		num, err := findLast(line)
		check(err)
		curr := first*10 + num
		total = total + curr
	}
	fmt.Printf("Part 2 %d\n", total)
}

func findFirst(input string) (int, error) {
	var builder strings.Builder

	for _, char := range input {
		if unicode.IsDigit(char) {
			return int(char - '0'), nil
		}

		builder.WriteString(string(char))
		word := builder.String()

		for index, number := range spelledNumbers {
			if strings.Contains(word, number) {
				return index, nil
			}
		}
	}

	return 0, errors.New("fail")
}

func findLast(input string) (int, error) {
	word := ""

	runes := []rune(input)

	for i := len(runes) - 1; i >= 0; i-- {
		if unicode.IsDigit(runes[i]) {
			return int(runes[i] - '0'), nil
		}
		word = string(runes[i]) + word

		for index, number := range spelledNumbers {
			if strings.Contains(word, number) {
				return index, nil
			}
		}
	}

	return 0, errors.New("fail")
}
