package main

import (
	"fmt"
	"stoykotolev/aoc-2023/utils"
	"strings"
	"time"
)

func main() {
	input := utils.ReadFile("./input.txt")

	part1(input)

}

func part1(input []string) {
	start := time.Now()
	sequence := strings.Split(strings.ReplaceAll(input[0], "\n", ""), ",")

	total := 0

	for _, step := range sequence {
		values := []byte(step)
		currentValue := 0
		for _, val := range values {
			currentValue += int(val)
			currentValue *= 17
			currentValue = currentValue % 256
		}
		total += currentValue
	}

	fmt.Println(total)
	fmt.Println(time.Since(start))
}
