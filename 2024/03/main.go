package main

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/stoykotolev/aoc/2024/utils"
)

func main() {
	input := utils.ReadFile("input.txt")
	part1(input)
	part2(input)
}

func part1(input []string) {
	pattern := `mul\((\d+),(\d+)\)`
	re := regexp.MustCompile(pattern)

	total := 0
	for _, line := range input {
		matches := re.FindAllStringSubmatch(line, -1)

		for _, match := range matches {
			a, _ := strconv.Atoi(match[1])
			b, _ := strconv.Atoi(match[2])
			total += a * b
		}
	}

	fmt.Println(total)
}

func part2(input []string) {
	total := 0
	pattern := `(?:don't\(\)|do\(\)|mul\((\d+),(\d+)\))`
	re := regexp.MustCompile(pattern)

	shouldCalc := true
	for _, line := range input {
		matches := re.FindAllStringSubmatch(line, -1)

		for _, match := range matches {
			if match[0] == "don't()" {
				shouldCalc = false
			} else if match[0] == "do()" {
				shouldCalc = true
			}
			if shouldCalc {

				a, _ := strconv.Atoi(match[1])
				b, _ := strconv.Atoi(match[2])
				total += a * b
			}
		}
	}

	fmt.Println(total)
}
