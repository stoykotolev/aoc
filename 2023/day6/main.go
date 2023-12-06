package main

import (
	"fmt"
	"strconv"
	"strings"

	"stoykotolev/aoc-2023/utils"
)

func main() {
	inputContents := utils.ReadFile("./input.txt")
	part1(inputContents)
	part2(inputContents)
}

func part1(input []string) {
	totalVariants := 1
	duration := strings.Fields(input[0])[1:]
	distance := strings.Fields(input[1])[1:]
	groups := make([][]int, 0)

	for i := 0; i < len(distance); i++ {
		dur, _ := strconv.Atoi(duration[i])
		dist, _ := strconv.Atoi(distance[i])

		groups = append(groups, []int{dur, dist})
	}

	for _, group := range groups {
		currRecord := group[1]
		raceDuration := group[0]
		wins := 0
		for i := 1; i < raceDuration; i++ {
			dist := i * (raceDuration - i)
			if dist > currRecord {
				wins++
			}
		}

		totalVariants *= wins
	}

	fmt.Println(totalVariants)
}

func part2(input []string) {
	totalVariants := 0
	raceDuration := strings.Split(strings.ReplaceAll(input[0], " ", ""), ":")[1]
	currentRecord := strings.Split(strings.ReplaceAll(input[1], " ", "")[1:], ":")[1]
	raceDur, _ := strconv.Atoi(raceDuration)
	currRec, _ := strconv.Atoi(currentRecord)

	for i := 1; i < raceDur; i++ {
		dist := i * (raceDur - i)
		if dist > currRec {
			totalVariants++
		}
	}

	fmt.Println(totalVariants)

}
