package main

import (
	"fmt"
	"math"
	"strings"

	"stoykotolev/aoc-2023/utils"
)

type InstructionsTuple struct {
	Left  string
	Right string
}

func main() {
	inputContents := utils.ReadFile("./input.txt")
	part1(inputContents)
	part2(inputContents)
}

func part1(input []string) {
	instructions := strings.Split(strings.TrimSpace(input[0]), "")

	instructionsMap := make(map[string]InstructionsTuple)
	for _, instruction := range input[2:] {
		fields := strings.Split(instruction, " = ")
		tupleStr := strings.TrimSuffix(strings.TrimPrefix(fields[1], "("), ")")
		tupleValues := strings.Split(tupleStr, ", ")
		instructionsMap[fields[0]] = InstructionsTuple{Left: tupleValues[0], Right: tupleValues[1]}
	}

	steps := walkInstructions(instructionsMap, instructions, "AAA")

	fmt.Println(steps)
}
func part2(input []string) {
	instructions := strings.Split(strings.TrimSpace(input[0]), "")

	instructionsMap := make(map[string]InstructionsTuple)
	for _, instruction := range input[2:] {
		fields := strings.Split(instruction, " = ")
		tupleStr := strings.TrimSuffix(strings.TrimPrefix(fields[1], "("), ")")
		tupleValues := strings.Split(tupleStr, ", ")
		instructionsMap[fields[0]] = InstructionsTuple{Left: tupleValues[0], Right: tupleValues[1]}
	}

	var startNodes []string
	for key := range instructionsMap {
		if strings.Contains(key, "A") {
			startNodes = append(startNodes, key)
		}
	}

	var steps []int
	for _, node := range startNodes {
		nodeStepsTillEnd := walkInstructions(instructionsMap, instructions, node)
		steps = append(steps, nodeStepsTillEnd)
	}

	lcm := lcmMultiple(steps)

	fmt.Println(lcm)
}

// Calculate LCM for a slice of numbers
func lcmMultiple(numbers []int) int {
	result := numbers[0]

	for i := 1; i < len(numbers); i++ {
		result = lcm(result, numbers[i])
	}

	return result
}

func walkInstructions(
	instructionsMap map[string]InstructionsTuple,
	instructions []string,
	startNode string,
) int {
	steps := 0
	currentInstructions := instructions

	nextStep := startNode
	for {
		if nextStep[len(nextStep)-1] == 'Z' {
			break
		}
		if currentInstructions[steps] == "L" {
			nextStep = instructionsMap[nextStep].Left
		} else {
			nextStep = instructionsMap[nextStep].Right
		}

		currentInstructions = append(currentInstructions, currentInstructions[steps])
		steps++
	}

	return steps
}

func allElementsContainZ(slice []string) bool {
	for _, element := range slice {
		if !strings.ContainsRune(element, 'Z') {
			return false
		}
	}
	return true
}

// Calculate GCD using Euclidean algorithm
func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

// Calculate LCM using the formula: LCM(a, b) = |a * b| / GCD(a, b)
func lcm(a, b int) int {
	if a == 0 || b == 0 {
		return 0 // LCM is not defined for zero
	}
	return int(math.Abs(float64(a*b)) / float64(gcd(a, b)))
}
