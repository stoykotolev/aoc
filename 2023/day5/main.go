package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type DirectionMap struct {
	Dest   int
	Source int
	Range  int
}

func main() {
	file, _ := os.ReadFile("./input.txt")
	inputContents := strings.Split(string(file), "\n\n")

	part1(inputContents)
	part2(inputContents)
}

func part1(input []string) {
	seeds := strings.Fields(input[0])[1:]
	directionMap := generateRanges(input[1:])

	minLocation := math.MaxInt

	for _, seed := range seeds {
		source, _ := strconv.Atoi(seed)

		minLocationForSeed := getMinLocationForSeed(source, directionMap)

		if minLocationForSeed < minLocation {
			minLocation = minLocationForSeed
		}
	}

	fmt.Println(minLocation)
}

func part2(input []string) {
	seedsRanges := strings.Fields(strings.Split(input[0], ":")[1])
	directionMap := generateRanges(input[1:])

	minLocation := math.MaxInt

	for i := 0; i < len(seedsRanges); i += 2 {
		seed, _ := strconv.Atoi(seedsRanges[i])
		length, _ := strconv.Atoi(seedsRanges[i+1])

		rangeMinLocation := math.MaxInt

		for j := seed; j <= seed+length; j++ {
			currentSeed := j
			minLocationForSeed := getMinLocationForSeed(currentSeed, directionMap)

			if minLocationForSeed < rangeMinLocation {
				rangeMinLocation = minLocationForSeed
			}
		}

		if rangeMinLocation < minLocation {
			minLocation = rangeMinLocation
		}
	}

	fmt.Println(minLocation)
}

func getMinLocationForSeed(seed int, directionMaps map[int][]DirectionMap) int {
	source := seed
	for i := 0; i < len(directionMaps); i++ {
		for _, directionMap := range directionMaps[i] {
			sourceRange := directionMap.Source + directionMap.Range
			if directionMap.Source <= source && source < sourceRange {
				source = source - directionMap.Source + directionMap.Dest
				break
			}
		}
	}

	return source
}

func generateRanges(mapping []string) map[int][]DirectionMap {
	var directionMap = make(map[int][]DirectionMap)

	for i, row := range mapping {
		values := strings.Split(row, "\n")[1:]

		for _, value := range values {
			if len(value) == 0 {
				continue
			}
			input := strings.Fields(value)
			drs, _ := strconv.Atoi(input[0])
			srs, _ := strconv.Atoi(input[1])
			rl, _ := strconv.Atoi(input[2])
			directionMap[i] = append(directionMap[i], DirectionMap{drs, srs, rl})
		}
	}

	return directionMap
}
