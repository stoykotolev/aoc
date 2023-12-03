package main

import (
	"fmt"
	"regexp"
	"strconv"

	"stoykotolev/aoc-2023/utils"
)

type gridPosition struct {
	row int
	col int
}

var neighbors = [][2]int{
	{-1, 0}, {1, 0}, {0, -1}, {0, 1}, {-1, -1}, {-1, 1}, {1, -1}, {1, 1},
}

var specialChars map[gridPosition]bool

func main() {
	inputContents := utils.ReadFile("./input.txt")
	specialChars = make(map[gridPosition]bool)

	for rIndx, row := range inputContents {
		for cIndx, c := range row {
			if string(c) == "." {
				continue
			} else if string(c) == "@" {
				specialChars[gridPosition{row: rIndx, col: cIndx}] = false
			} else if string(c) == "#" {
				specialChars[gridPosition{row: rIndx, col: cIndx}] = false
			} else if string(c) == "$" {
				specialChars[gridPosition{row: rIndx, col: cIndx}] = false
			} else if string(c) == "%" {
				specialChars[gridPosition{row: rIndx, col: cIndx}] = false
			} else if string(c) == "-" {
				specialChars[gridPosition{row: rIndx, col: cIndx}] = false
			} else if string(c) == "&" {
				specialChars[gridPosition{row: rIndx, col: cIndx}] = false
			} else if string(c) == "*" {
				specialChars[gridPosition{row: rIndx, col: cIndx}] = true
			} else if string(c) == "+" {
				specialChars[gridPosition{row: rIndx, col: cIndx}] = false
			} else if string(c) == "/" {
				specialChars[gridPosition{row: rIndx, col: cIndx}] = false
			} else if string(c) == "=" {
				specialChars[gridPosition{row: rIndx, col: cIndx}] = false
			}
		}
	}
	part1(inputContents)
	part2(inputContents)
}

func part1(inputContents []string) {
	gridLength := len(inputContents)
	var totalSum int

	for rowIndx, row := range inputContents {
		rowLength := len(row)
		re := regexp.MustCompile(`\d+`)
		nums := re.FindAllStringIndex(row, -1)

		for _, num := range nums {
			if hasSpecialCharInProxy(num[0]-1, num[1], rowIndx, gridLength, rowLength) {
				numVal, err := strconv.Atoi(row[num[0]:num[1]])
				utils.Check(err)
				totalSum += numVal
			}
		}
	}

	fmt.Println(totalSum)
}

func part2(inputContents []string) {
	var totalGearRation int

	for gridPoint, isStartChar := range specialChars {
		gearValues := []int{}
		if isStartChar {
			re := regexp.MustCompile(`\d+`)
			// check the above row
			if gridPoint.row >= 1 {
				numsAbove := re.FindAllStringIndex(inputContents[gridPoint.row-1], -1)
				for _, num := range numsAbove {
					if (gridPoint.col-num[0] >= -1 && gridPoint.col-num[0] <= 1) ||
						(gridPoint.col-num[1] >= -2 && gridPoint.col-num[1] <= 0) {
						num, err := strconv.Atoi(inputContents[gridPoint.row-1][num[0]:num[1]])
						utils.Check(err)
						gearValues = append(gearValues, num)
					}
				}
			}

			//curr row
			numberIndicesOnRow := re.FindAllStringIndex(inputContents[gridPoint.row], -1)

			for _, numberIndex := range numberIndicesOnRow {
				if gridPoint.col-numberIndex[1] == 0 || gridPoint.col-numberIndex[0] == -1 {
					num, err := strconv.Atoi(
						inputContents[gridPoint.row][numberIndex[0]:numberIndex[1]],
					)
					utils.Check(err)
					gearValues = append(gearValues, num)
				}
			}

			// next row
			if len(inputContents) > gridPoint.row+1 {
				numsBelow := re.FindAllStringIndex(inputContents[gridPoint.row+1], -1)
				for _, num := range numsBelow {
					if (gridPoint.col-num[0] >= -1 && gridPoint.col-num[0] <= 1) ||
						(gridPoint.col-num[1] >= -2 && gridPoint.col-num[1] <= 0) {
						num, err := strconv.Atoi(inputContents[gridPoint.row+1][num[0]:num[1]])
						utils.Check(err)
						gearValues = append(gearValues, num)
					}
				}
			}

			if len(gearValues) == 2 {
				totalGearRation += gearValues[0] * gearValues[1]
			}
		}
	}
	fmt.Println(totalGearRation)
}

func hasSpecialCharInProxy(start, end, row, numOfRows, rowLength int) bool {
	if start < 0 {
		start = 0
	}
	if end > rowLength {
		end = rowLength
	}

	// line above
	if row != 0 {
		for i := start; i <= end; i++ {
			_, ok := specialChars[gridPosition{row - 1, i}]
			if ok {
				return ok
			}
		}
	}

	_, hasOnSameStart := specialChars[gridPosition{row, start}]

	if hasOnSameStart {
		return hasOnSameStart
	}

	_, hasOnSameEnd := specialChars[gridPosition{row, end}]
	if hasOnSameEnd {
		return hasOnSameEnd
	}

	if row != numOfRows-1 {
		for i := start; i <= end; i++ {
			_, ok := specialChars[gridPosition{row + 1, i}]
			if ok {
				return ok
			}
		}

	}

	return false
}
