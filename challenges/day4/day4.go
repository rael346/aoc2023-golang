package day4

import (
	"math"
	"regexp"
	"slices"
	"strings"
)

func Part1(input []string) int {
	sum := 0
	cardRegex := regexp.MustCompile(`Card\s+\d+: ([\d ]+)\| ([\d ]+)`)

	for _, line := range input {
		parsed := cardRegex.FindStringSubmatch(line)
		ourNums := strings.Fields(parsed[1])
		winNums := strings.Fields(parsed[2])

		matches := 0
		for _, ourNum := range ourNums {
			if slices.Contains(winNums, ourNum) {
				matches += 1
			}
		}

		sum += int(math.Pow(2, float64(matches-1)))
	}

	return sum
}

func Part2(input []string) int {
	sum := 0
	cardRegex := regexp.MustCompile(`Card\s+\d+: ([\d ]+)\| ([\d ]+)`)
	cardCount := make([]int, len(input))

	for i, line := range input {
		parsed := cardRegex.FindStringSubmatch(line)
		ourNums := strings.Fields(parsed[1])
		winNums := strings.Fields(parsed[2])

		cardCount[i] += 1
		currCopyIndex := i + 1
		for _, ourNum := range ourNums {
			if slices.Contains(winNums, ourNum) {
				cardCount[currCopyIndex] += cardCount[i]
				currCopyIndex += 1
			}
		}

		sum += cardCount[i]
	}

	return sum
}
