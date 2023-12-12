package main

import (
	"fmt"
	"stoykotolev/aoc-2023/utils"
	"strconv"
	"strings"
	"time"
)

func main() {
	input := utils.ReadFile("./input.txt")
	var springsMap []string
	var groups [][]uint8
	for i := range input {
		lineGroups := strings.Split(input[i], " ")
		springsMap = append(springsMap, lineGroups[0])
		var group []uint8
		for _, num := range strings.Split(lineGroups[1], ",") {
			group = append(group, uint8(ToInt(num)))
		}
		groups = append(groups, group)
	}
	part1(springsMap, groups)
	part2(springsMap, groups)
}

func part1(springsMap []string, groups [][]uint8) {
	start := time.Now()

	total := 0

	for i := range springsMap {
		currentSpringMap := springsMap[i]
		springMapConditionRecord := groups[i]

		total += countVariants(currentSpringMap, springMapConditionRecord)
	}

	fmt.Println(total)
	fmt.Println(time.Since(start))

}
func part2(springsMap []string, groups [][]uint8) {
	start := time.Now()
	total := 0

	for i := range springsMap {
		currentSpringMap := unfoldSpringMap(springsMap[i])
		springMapConditionRecord := unfoldRecordGroup(groups[i])
		currentSpringMapTotal := countVariants(currentSpringMap, springMapConditionRecord)

		total += currentSpringMapTotal
	}

	fmt.Println(total)
	fmt.Println(time.Since(start))
}

type state struct {
	configuration string
	records       string
}

var cache = make(map[state]int)

func countVariants(configuration string, records []uint8) int {

	if len(configuration) == 0 {
		if len(records) == 0 {
			return 1
		}
		return 0
	}

	if len(records) == 0 {
		if strings.Contains(configuration, "#") {
			return 0
		}
		return 1
	}

	if value, ok := cache[state{configuration, string(records)}]; ok {
		return value
	}

	result := 0

	if configuration[0] == '.' || configuration[0] == '?' {
		result += countVariants(configuration[1:], records)
	}

	confLength := uint8(len(configuration))
	if configuration[0] == '#' || configuration[0] == '?' {
		if records[0] <= confLength &&
			!strings.Contains(configuration[:records[0]], ".") &&
			(records[0] == confLength || configuration[records[0]] != '#') {
			if records[0]+1 < confLength {
				result += countVariants(configuration[records[0]+1:], records[1:])
			} else {
				result += countVariants("", records[1:])
			}
		}
	}

	cache[state{configuration, string(records)}] = result
	return result
}

func unfoldSpringMap(springMap string) string {
	originalSpringMap := springMap
	for i := 0; i < 4; i++ {
		springMap += "?" + originalSpringMap
	}

	return springMap
}

func unfoldRecordGroup(records []uint8) []uint8 {
	originalRecords := records
	for i := 0; i < 4; i++ {
		records = append(records, originalRecords...)
	}

	return records
}

func ToInt(s string) uint64 {
	res, _ := strconv.Atoi(s)
	return uint64(res)
}
