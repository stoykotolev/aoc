package main

import (
	"fmt"
	"strings"

	"stoykotolev/aoc-2023/utils"
)

func main() {
	inputContents := utils.ReadFile("./input.txt")
	part1(inputContents)
	part2(inputContents)
}

func part1(games []string) {
	var totalPoints int
	for _, game := range games {
		var currentCardPoints int
		firstWin := true
		gameSplit := strings.Split(game, ":")
		numbersGroups := strings.Split(gameSplit[1], "|")
		winningNumbers := strings.Fields(strings.TrimSpace(numbersGroups[0]))
		myNumbers := strings.Fields(strings.TrimSpace(numbersGroups[1]))

		for _, myNumber := range myNumbers {
			for _, winningNumber := range winningNumbers {
				if myNumber == winningNumber {
					if firstWin {
						currentCardPoints = 1
						firstWin = false
					} else {
						currentCardPoints *= 2
					}
				}
			}
		}
		totalPoints += currentCardPoints
	}

	fmt.Println(totalPoints)
}

func part2(games []string) {
	totalWonScratchCards := make([]int, len(games))
	for i := range totalWonScratchCards {
		totalWonScratchCards[i] = 1
	}
	for gameIndx, game := range games {
		gameSplit := strings.Split(game, ":")
		numbersGroups := strings.Split(gameSplit[1], "|")
		winningNumbers := strings.Fields(strings.TrimSpace(numbersGroups[0]))
		myNumbers := strings.Fields(strings.TrimSpace(numbersGroups[1]))
		correctNumbers := []string{}

		for _, myNumber := range myNumbers {
			for _, winningNumber := range winningNumbers {
				if myNumber == winningNumber {
					correctNumbers = append(correctNumbers, myNumber)
				}
			}
		}

		wonLength := len(correctNumbers)
		for wonLength > 0 {
			totalWonScratchCards[gameIndx+wonLength] += totalWonScratchCards[gameIndx]
			wonLength--
		}
	}

	var totalPointsWon int
	for _, val := range totalWonScratchCards {
		totalPointsWon += val
	}

	fmt.Println(totalPointsWon)
}
