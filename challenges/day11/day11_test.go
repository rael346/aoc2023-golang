package day11

import (
	"testing"

	"github.com/rael346/aoc2023-golang/internal/utils"
)

var givenInput = `
...#......
.......#..
#.........
..........
......#...
.#........
.........#
..........
.......#..
#...#.....
`

func TestPart1Given(t *testing.T) {
	result := Part1(utils.ToInput(givenInput))
	utils.ExpectInt(t, result, 374)
}

func TestPart1Actual(t *testing.T) {
	result := Part1(utils.ReadInputFromFile("input.txt"))
	utils.ExpectInt(t, result, 10154062)
}

func TestPart2Given(t *testing.T) {
	result := Part2(utils.ToInput(givenInput))
	utils.ExpectInt(t, result, 82000210)
}

func TestPart2Actual(t *testing.T) {
	result := Part2(utils.ReadInputFromFile("input.txt"))
	utils.ExpectInt(t, result, 553083047914)
}
