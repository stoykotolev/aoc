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
	part2(levels)
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

func part2(levels [][]int) {
	sl := 0

	for _, l := range levels {

		if isSafeDampener(l) {
			sl += 1
		}

	}
	fmt.Println(sl)
}

func isSafeDampener(l []int) bool {
	if isSafe(l) {
		return true
	}

	for i := range l {
		temp := []int{}
		temp = append(temp, l[:i]...)
		temp = append(temp, l[i+1:]...)
		if isSafe(temp) {
			return true
		}
	}

	return false
}

func isSafe(l []int) bool {

	ls := 0
	for i := 0; i < len(l)-1; i++ {
		diff := l[i] - l[i+1]
		if isAsc(l) {
			if diff <= -1 && diff >= -3 {
				ls++
			}
		} else {
			if diff >= 1 && diff <= 3 {
				ls++
			}
		}
	}

	if ls == len(l)-1 {
		return true
	}
	return false

}

func isAsc(l []int) bool {
	for i := 0; i < len(l)-1; i++ {
		if l[i] > l[i+1] {
			return false
		}
	}

	return true
}
