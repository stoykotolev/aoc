package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/stoykotolev/aoc/2024/utils"
)

func main() {
	input := utils.ReadFile("test.txt")
	levels := [][]int{}

	for _, el := range input {
		splitEl := strings.Split(el, " ")
		level := []int{}
		for _, num := range splitEl {
			parsedEl, _ := strconv.Atoi(num)
			level = append(level, parsedEl)
		}
		levels = append(levels, level)
	}
	part1(levels)
}

func part1(levels [][]int) {
	safeLevels := 0

	for _, level := range levels {
		if isSafe(level) {
			safeLevels += 1
		}
	}

	fmt.Println(safeLevels)
}

func isSafe(level []int) bool {

	for i := 1; i < len(level); i++ {

		if !isLevelSingleType(level) {
			return false
		}

		diff := (level[i] - level[i-1])

		if diff < 0 {
			diff = diff * -1
		}

		if diff < 1 || diff > 3 {
			return false
		}
	}

	return true
}

func isLevelSingleType(level []int) bool {

	isAscending := true

	if level[0] > level[1] {
		isAscending = false
	}

	dampen := true
	for i := 1; i < len(level); i++ {
		if level[i] == level[i-1] {
			if dampen {
				dampen = false
				continue
			}
			return false
		}
		if isAscending {
			if level[i] < level[i-1] {
				if dampen {
					dampen = false
					continue
				}
				return false
			}
		} else {
			if level[i] > level[i-1] {
				if dampen {
					dampen = false
					continue
				}
				return false
			}
		}

	}

	return true
}

func hasRepeatedElements(level []int) bool {
	for i := 1; i < len(level); i++ {
		if level[i] == level[i-1] {
			return true
		}
	}

	return false
}
