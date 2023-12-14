package main

import (
	"fmt"
	"stoykotolev/aoc-2023/utils"
)

func main() {

	input := utils.ReadFile("./input.txt")

	grid := constructGrid(input)

	part1(grid)
	part2(input)
}

func part1(grid [][]byte) {
	northGrid := moveNorth(grid)
	total := 0
	gridLength := len(northGrid)

	for i, row := range northGrid {
		rollingRocksCount := countByte(row, 'O')
		total += rollingRocksCount * (gridLength - i)
	}

	fmt.Println(total)
}

func part2(input []string) {}

func constructGrid(input []string) [][]byte {
	grid := make([][]byte, len(input))

	for i, line := range input {
		grid[i] = []byte(line)
	}

	return grid
}

func moveNorth(grid [][]byte) [][]byte {

	// we need to on everything from the first row
	for row := 1; row < len(grid); row++ {
		// and we need to iterate over each column value
		for col := 0; col < len(grid[row]); col++ {
			// if the current value is a round rock
			if grid[row][col] == 'O' {
				// now we need to check each previous row in the same column and if that value is a ., replace the 2 values
				for backRow := row - 1; backRow >= 0; backRow-- {
					if grid[backRow][col] != '.' {
						break
					}
					grid[backRow][col] = 'O'
					grid[backRow+1][col] = '.'
				}
			}
		}
	}

	return grid
}

func countByte(byteSlice []byte, targetByte byte) int {
	count := 0
	for _, b := range byteSlice {
		if b == targetByte {
			count++
		}
	}
	return count
}
