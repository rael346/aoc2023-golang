package day3

import (
	"testing"

	"github.com/rael346/aoc2023-golang/internal/utils"
)

var givenInput = `
467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..
`

func TestPart1Given(t *testing.T) {
	result := Part1(utils.ToInput(givenInput))
	utils.ExpectInt(t, result, 4361)
}

func TestPart1Actual(t *testing.T) {
	result := Part1(utils.ReadInputFromFile("input.txt"))
	utils.ExpectInt(t, result, 557705)
}

func TestPart2Given(t *testing.T) {
	result := Part2(utils.ToInput(givenInput))
	utils.ExpectInt(t, result, 467835)
}

func TestPart2Actual(t *testing.T) {
	result := Part2(utils.ReadInputFromFile("input.txt"))
	utils.ExpectInt(t, result, 84266818)
}
