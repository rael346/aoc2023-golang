package day9

import (
	"testing"

	"github.com/rael346/aoc2023-golang/internal/utils"
)

var givenInput = `
0 3 6 9 12 15
1 3 6 10 15 21
10 13 16 21 30 45
`

func TestPart1Given1(t *testing.T) {
	result := Part1(utils.ToInput(givenInput))
	utils.ExpectInt(t, result, 114)
}

func TestPart1Actual(t *testing.T) {
	result := Part1(utils.ReadInputFromFile("input.txt"))
	utils.ExpectInt(t, result, 1930746032)
}

func TestPart2Given(t *testing.T) {
	result := Part2(utils.ToInput(givenInput))
	utils.ExpectInt(t, result, 2)
}

func TestPart2Actual(t *testing.T) {
	result := Part2(utils.ReadInputFromFile("input.txt"))
	utils.ExpectInt(t, result, 1154)
}
