package main

import (
	"fmt"
	"strconv"
	"strings"

	"stoykotolev/aoc-2023/utils"
)

func main() {
	inputContents := utils.ReadFile("./input.txt")
	var inputGroups [][]int
	for _, line := range inputContents {
		inputNumbers := strings.Fields(line)
		parsedNumbers := make([]int, 0)
		for _, num := range inputNumbers {
			paredNum, _ := strconv.Atoi(num)
			parsedNumbers = append(parsedNumbers, paredNum)
		}
		inputGroups = append(inputGroups, parsedNumbers)
	}
	part1(inputGroups)
	part2(inputGroups)
}

func part1(input [][]int) {
	extrapolatedValues := make([]int, 0)
	for _, group := range input {
		currentGroups := make([][]int, 0)
		currentGroups = append(currentGroups, group)

		for i := 0; i < len(currentGroups); i++ {
			currentGroup := currentGroups[i]
			newGroup := make([]int, 0)
			for value := 1; value < len(currentGroup); value++ {
				newValue := currentGroup[value] - currentGroup[value-1]
				newGroup = append(newGroup, newValue)
			}
			allZeroes := everyIsZero(newGroup)
			if allZeroes {
				newGroup = append(newGroup, 0)
				currentGroups = append(currentGroups, newGroup)
				break
			}
			currentGroups = append(currentGroups, newGroup)
		}
		lastValueOfPrevGroup := 0
		for i := len(currentGroups) - 2; i >= 0; i-- {
			currentGroup := currentGroups[i]
			currentGroupLastValue := currentGroup[len(currentGroup)-1]
			total := currentGroupLastValue + lastValueOfPrevGroup
			lastValueOfPrevGroup = total

			if i == 0 {
				extrapolatedValues = append(extrapolatedValues, lastValueOfPrevGroup)
			}
		}
	}

	total := 0
	for _, num := range extrapolatedValues {
		total += num
	}

	fmt.Println(total)

}
func part2(input [][]int) {
	extrapolatedValues := make([]int, 0)
	for _, group := range input {
		currentGroups := make([][]int, 0)
		currentGroups = append(currentGroups, group)

		for i := 0; i < len(currentGroups); i++ {
			currentGroup := currentGroups[i]
			newGroup := make([]int, 0)
			for value := 1; value < len(currentGroup); value++ {
				newValue := currentGroup[value] - currentGroup[value-1]
				newGroup = append(newGroup, newValue)
			}
			allZeroes := everyIsZero(newGroup)
			if allZeroes {
				newGroup = append(newGroup, 0)
				currentGroups = append(currentGroups, newGroup)
				break
			}
			currentGroups = append(currentGroups, newGroup)
		}
		firstValueOfPrevGroup := 0
		for i := len(currentGroups) - 2; i >= 0; i-- {
			currentGroup := currentGroups[i]
			currentGroupLastValue := currentGroup[0]
			total := currentGroupLastValue - firstValueOfPrevGroup
			firstValueOfPrevGroup = total

			if i == 0 {
				extrapolatedValues = append(extrapolatedValues, firstValueOfPrevGroup)
			}
		}
	}

	total := 0
	for _, num := range extrapolatedValues {
		total += num
	}

	fmt.Println(total)

}

func everyIsZero(numberGroup []int) bool {
	for _, num := range numberGroup {
		if num != 0 {
			return false
		}
	}

	return true
}
