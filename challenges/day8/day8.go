package day8

import (
	"regexp"
	"strings"

	"github.com/rael346/aoc2023-golang/internal/utils"
)

var instructMap = map[rune]int{
	'L': 0,
	'R': 1,
}

func Part1(input []string) int {
	nodeTable := map[string][2]string{}
	nodeRegex := regexp.MustCompile(`\w+`)
	for _, line := range input[2:] {
		node := nodeRegex.FindAllString(line, 3)
		nodeTable[node[0]] = [2]string{node[1], node[2]}
	}

	instructions := []rune(input[0])
	return getSteps(instructions, "AAA", nodeTable)
}

func Part2(input []string) int {
	nodeTable := map[string][2]string{}
	startNodes := []string{}
	nodeRegex := regexp.MustCompile(`\w+`)
	for _, line := range input[2:] {
		node := nodeRegex.FindAllString(line, 3)
		nodeTable[node[0]] = [2]string{node[1], node[2]}
		if strings.HasSuffix(node[0], "A") {
			startNodes = append(startNodes, node[0])
		}
	}

	instructions := []rune(input[0])
	steps := make([]int, len(startNodes))
	for i, startNode := range startNodes {
		steps[i] = getSteps(instructions, startNode, nodeTable)
	}

	return utils.LCM(steps[0], steps[1], steps...)
}

func getSteps(instructions []rune, startNode string, nodeTable map[string][2]string) int {
	steps := 0
	currentNode := startNode
	for !strings.HasSuffix(currentNode, "Z") {
		instruct := instructions[steps%len(instructions)]
		currentNode = nodeTable[currentNode][instructMap[instruct]]
		steps += 1
	}

	return steps
}
