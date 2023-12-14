package main

import (
	"bytes"
	"fmt"
	"stoykotolev/aoc-2023/utils"
)

func main() {

	input := utils.ReadFile("./input.txt")

	grid := constructGrid(input)

	part1(grid)
	part2(grid)
}

func part1(grid [][]byte) {
	moveN(grid)
	total := countRocks(grid)

	fmt.Println(total)
}

func part2(grid [][]byte) {
	gridCache := make(map[string]int)
	totalCycles := 1000000000
	for i := 0; i < totalCycles; i++ {
		moveN(grid)
		moveW(grid)
		moveS(grid)
		moveE(grid)
		hash := bytes.Join(grid, []byte{})
		if _, ok := gridCache[string(hash)]; ok {
			i = totalCycles - (totalCycles-i)%(i-gridCache[string(hash)])
		}
		gridCache[string(hash)] = i
	}

	fmt.Println(countRocks(grid))
}

func constructGrid(input []string) [][]byte {
	grid := make([][]byte, len(input))

	for i, line := range input {
		grid[i] = []byte(line)
	}

	return grid
}

const (
	SOLID = byte('#')
	ROUND = byte('O')
	EMPTY = byte('.')
)

func moveN(grid [][]byte) {
	for x := 0; x < len(grid[0]); x++ {
		available := 0
		for y := 0; y < len(grid); y++ {
			switch grid[y][x] {
			case SOLID:
				available = y + 1
			case ROUND:
				if available < y {
					grid[available][x] = ROUND
					grid[y][x] = EMPTY
				}
				available++
			}
		}
	}
}

func moveS(grid [][]byte) {
	for x := 0; x < len(grid[0]); x++ {
		available := len(grid) - 1
		for y := len(grid) - 1; y >= 0; y-- {
			switch grid[y][x] {
			case SOLID:
				available = y - 1
			case ROUND:
				if available > y {
					grid[available][x] = ROUND
					grid[y][x] = EMPTY
				}
				available--
			}
		}
	}
}

func moveW(grid [][]byte) {
	for y := 0; y < len(grid); y++ {
		available := 0
		for x := 0; x < len(grid[y]); x++ {
			switch grid[y][x] {
			case SOLID:
				available = x + 1
			case ROUND:
				if available < x {
					grid[y][available] = ROUND
					grid[y][x] = EMPTY
				}
				available++
			}
		}
	}
}

func moveE(grid [][]byte) {
	for y := 0; y < len(grid); y++ {
		available := len(grid[y]) - 1
		for x := len(grid[y]) - 1; x >= 0; x-- {
			switch grid[y][x] {
			case SOLID:
				available = x - 1
			case ROUND:
				if available > x {
					grid[y][available] = ROUND
					grid[y][x] = EMPTY
				}
				available--
			}
		}
	}
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

func countRocks(grid [][]byte) int {
	total := 0
	gridLength := len(grid)
	for i, row := range grid {
		rollingRocksCount := countByte(row, 'O')
		total += rollingRocksCount * (gridLength - i)
	}

	return total
}
