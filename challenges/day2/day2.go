package day2

import (
	"regexp"
	"strconv"
	"strings"
)

var colorLimit = map[string]int{
	"red":   12,
	"green": 13,
	"blue":  14,
}

func Part1(input []string) int {
	sumID := 0
	for _, line := range input {
		res := regexp.MustCompile(`[;:]`).Split(line, -1)
		gameID, _ := strconv.Atoi(strings.Split(res[0], " ")[1])
		isPossible := true

	outer:
		for _, match := range res[1:] {
			for _, cubes := range strings.Split(match, ",") {
				cubesParsed := strings.Split(strings.TrimSpace(cubes), " ")
				count, _ := strconv.Atoi(cubesParsed[0])
				color := cubesParsed[1]

				if count > colorLimit[color] {
					isPossible = false
					break outer
				}
			}
		}

		if isPossible {
			sumID += gameID
		}
	}
	return sumID
}

func Part2(input []string) int {
	sum := 0
	for _, line := range input {
		res := regexp.MustCompile(`[;:]`).Split(line, -1)

		maxCube := map[string]int{
			"red":   0,
			"green": 0,
			"blue":  0,
		}

		for _, match := range res[1:] {
			for _, cubes := range strings.Split(match, ",") {
				cubesParsed := strings.Split(strings.TrimSpace(cubes), " ")
				count, _ := strconv.Atoi(cubesParsed[0])
				color := cubesParsed[1]

				maxCube[color] = max(maxCube[color], count)
			}
		}

		power := 1
		for _, count := range maxCube {
			power *= count
		}

		sum += power
	}
	return sum
}
