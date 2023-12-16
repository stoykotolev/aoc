package main

import (
	"fmt"
	"stoykotolev/aoc-2023/utils"
	"time"
)

const (
	VERTICAL   = '|'
	HORIZONTAL = '-'
	LEFT       = '/'
	RIGHT      = '\\'
	EMPTY      = '.'
)

const (
	E = iota
	S
	W
	N
)

var directions = []pos{
	E: {1, 0},  // E
	S: {0, 1},  // S
	W: {-1, 0}, // W
	N: {0, -1}, //N
}

type pos struct {
	col int
	row int
}

var MIRROR_A = []int{
	E: N,
	W: S,
	S: W,
	N: E,
}

var MIRROR_B = []int{
	E: S,
	S: E,
	W: N,
	N: W,
}

var seen = make(map[[3]int]bool)

func main() {
	input := utils.ReadFile("./input.txt")
	grid := constructGrid(input)
	part1(grid)
	part2(grid)
}

func part2(grid [][]byte) {
	start := time.Now()
	best := 0
	for row := range grid {
		clear(seen)
		firstEnergyGrid := createEnergyGrid(grid)
		energizeGrid(grid, firstEnergyGrid, pos{0, row}, E)

		firstTotal := totalEnergy(firstEnergyGrid)
		if firstTotal > best {
			best = firstTotal
		}

		clear(seen)
		secondEnergyGrid := createEnergyGrid(grid)
		energizeGrid(grid, secondEnergyGrid, pos{len(grid[0]) - 1, row}, W)
		secondTotal := totalEnergy(secondEnergyGrid)
		if secondTotal > best {
			best = secondTotal
		}
	}

	for col := range grid[0] {
		clear(seen)
		firstEnergyGrid := createEnergyGrid(grid)
		energizeGrid(grid, firstEnergyGrid, pos{col, 0}, S)
		firstTotal := totalEnergy(firstEnergyGrid)
		if firstTotal > best {
			best = firstTotal
		}

		clear(seen)
		secondEnergyGrid := createEnergyGrid(grid)
		energizeGrid(grid, secondEnergyGrid, pos{col, len(grid) - 1}, N)
		secondTotal := totalEnergy(secondEnergyGrid)
		if secondTotal > best {
			best = secondTotal
		}
	}

	fmt.Println(best)
	fmt.Println(time.Since(start))
}

func part1(grid [][]byte) {
	start := time.Now()

	energyGrid := createEnergyGrid(grid)
	energizeGrid(grid, energyGrid, pos{0, 0}, E)
	totalEnergy := totalEnergy(energyGrid)

	fmt.Println(totalEnergy)
	fmt.Println(time.Since(start))
}

func constructGrid(input []string) [][]byte {
	grid := make([][]byte, len(input))

	for i, row := range input {
		grid[i] = []byte(row)
	}

	return grid
}

func createEnergyGrid(grid [][]byte) [][]int {
	energyGrid := make([][]int, len(grid))
	for i, row := range grid {
		energyGrid[i] = make([]int, len(row))
	}
	return energyGrid
}

func totalEnergy(energyGrid [][]int) int {
	total := 0

	for _, row := range energyGrid {
		for _, col := range row {
			if col > 0 {
				total += 1
			}
		}
	}

	return total
}

func energizeGrid(grid [][]byte, energy [][]int, position pos, dir int) {
	for {
		hashKey := [3]int{position.row, position.col, dir}

		if _, ok := seen[hashKey]; ok {
			return
		}

		energy[position.row][position.col] = 1
		seen[hashKey] = true

		switch grid[position.row][position.col] {
		case LEFT:
			dir = MIRROR_A[dir]
			break
		case RIGHT:
			dir = MIRROR_B[dir]
		case VERTICAL:
			if dir == E || dir == W {
				energizeGrid(grid, energy, position, N)
				energizeGrid(grid, energy, position, S)
				return
			}
			break
		case HORIZONTAL:
			if dir == S || dir == N {
				energizeGrid(grid, energy, position, W)
				energizeGrid(grid, energy, position, E)
				return
			}
			break
		}
		position.row += directions[dir].row
		position.col += directions[dir].col

		if position.row < 0 || position.row >= len(grid) {
			return
		}

		if position.col < 0 || position.col >= len(grid[0]) {
			return
		}
	}
}
