package day6

import (
	"testing"

	"github.com/rael346/aoc2023-golang/internal/utils"
)

var givenInput = ` Time:      7  15   30
Distance:  9  40  200`

func TestPart1Given(t *testing.T) {
	result := Part1(utils.ToInput(givenInput))
	utils.ExpectInt(t, result, 288)
}

func TestPart1Actual(t *testing.T) {
	result := Part1(utils.ReadInputFromFile("input.txt"))
	utils.ExpectInt(t, result, 3316275)
}

func TestPart2Given(t *testing.T) {
	result := Part2(utils.ToInput(givenInput))
	utils.ExpectInt(t, result, 71503)
}

func TestPart2Actual(t *testing.T) {
	result := Part2(utils.ReadInputFromFile("input.txt"))
	utils.ExpectInt(t, result, -1)
}
