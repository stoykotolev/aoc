package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/stoykotolev/aoc/2024/utils"
)

func main() {
	input := utils.ReadFile("./input.txt")
	groupOne := make([]int, len(input))
	groupTwo := make([]int, len(input))
	for _, el := range input {
		parts := strings.Fields(el)
		num1, _ := strconv.Atoi(parts[0])
		num2, _ := strconv.Atoi(parts[1])
		groupOne = append(groupOne, num1)
		groupTwo = append(groupTwo, num2)
	}

	sort.Ints(groupOne)
	sort.Ints(groupTwo)
	part1(groupOne, groupTwo)
	part2(groupOne, groupTwo)
}

func part1(groupOne, groupTwo []int) {
	totalDistance := 0
	for i := range groupOne {
		a := groupOne[i]
		b := groupTwo[i]
		if a > b {
			totalDistance += a - b
		} else {
			totalDistance += b - a
		}
	}

	fmt.Println(totalDistance)

}

func part2(groupOne, groupTwo []int) {

	simScore := 0
	for _, a := range groupOne {
		count := 0
		for _, b := range groupTwo {
			if a == b {
				count += 1
			}
		}
		simScore += a * count
	}

	fmt.Println(simScore)
}
