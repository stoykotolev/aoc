package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"stoykotolev/aoc-2023/utils"
)

var targetColors = map[string]int{
	"red":   12,
	"green": 13,
	"blue":  14,
}

func main() {
	inputContents := utils.ReadFile("./input.txt")
	part1(inputContents)
	part2(inputContents)

}

func part1(games []string) {
	var sumGameIds int
	for _, game := range games {
		isValidGame := true
		gameSplit := strings.Split(strings.TrimSpace(game), ":")
		gameIndexStr := strings.Split(gameSplit[0], " ")[1]
		re := regexp.MustCompile(`(\d+)\s+(\w+)`)
		matches := re.FindAllString(game, -1)

		for _, match := range matches {
			colorSlice := strings.Split(strings.TrimSpace(match), " ")
			color := colorSlice[1]
			count, err := strconv.Atoi(colorSlice[0])
			utils.Check(err)
			if targetColors[color] < count {
				isValidGame = false
			}
		}
		if isValidGame {
			gameIndex, err := strconv.Atoi(gameIndexStr)
			utils.Check(err)
			sumGameIds += gameIndex
		}

	}
	fmt.Println(sumGameIds)
}

func part2(games []string) {
	var totalPowerSum int

	for _, game := range games {
		colorPower := 1
		colorCounts := make(map[string]int)
		re := regexp.MustCompile(`(\d+)\s+(\w+)`)
		matches := re.FindAllString(game, -1)

		for _, match := range matches {
			colorSlice := strings.Split(strings.TrimSpace(match), " ")
			color := colorSlice[1]
			count, err := strconv.Atoi(colorSlice[0])
			utils.Check(err)
			if colorCounts[color] < count {
				colorCounts[color] = count
			}
		}
		for _, value := range colorCounts {
			colorPower *= value
		}
		totalPowerSum += colorPower
	}
	fmt.Println(totalPowerSum)
}
