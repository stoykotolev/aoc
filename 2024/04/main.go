package main

import (
	"fmt"

	"github.com/stoykotolev/aoc/2024/utils"
)

func main() {
	input := utils.ReadFile("input.txt")
	var puzzle [][]byte

	for _, line := range input {
		letters := []byte(line)
		puzzle = append(puzzle, letters)
	}
	p1(puzzle)
	p2(puzzle)
}

var dirs = [][]int{{0, 1}, {0, -1}, {-1, 0}, {1, 0}, {-1, -1}, {-1, 1}, {1, -1}, {1, 1}}

func p1(p [][]byte) {
	total := 0

	for r := range p {
		for c := range p[r] {
			for _, d := range dirs {
				if findWord(p, r, c, d[0], d[1]) {
					total++
				}
			}
		}
	}

	fmt.Println(total)
}

func p2(p [][]byte) {
	total := 0
	bv := byte('M') + byte('S')
	for r := range p {
		for c := range p[r] {
			if p[r][c] == byte('A') {
				if r-1 < 0 || r+1 > len(p)-1 || c-1 < 0 || c+1 > len(p[r])-1 {
					continue
				}
				d1 := p[r-1][c-1] + p[r+1][c+1]
				d2 := p[r-1][c+1] + p[r+1][c-1]

				if d1 == bv && d2 == bv {
					total++
				}
			}
		}
	}

	fmt.Println(total)
}

func findWord(p [][]byte, r, c, dx, dy int) bool {
	word := "XMAS"

	for l := range word {
		nx := r + l*dx
		ny := c + l*dy
		if nx < 0 || nx > len(p)-1 || ny < 0 || ny > len(p[r])-1 {
			return false
		}
		if p[nx][ny] != word[l] {
			return false
		}
	}

	return true
}
