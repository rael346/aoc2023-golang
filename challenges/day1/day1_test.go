package day1

import (
	"testing"

	"github.com/rael346/aoc2023-golang/internal/utils"
)

func TestPart1Given(t *testing.T) {
	result := Part1(utils.ToInput(`
1abc2
pqr3stu8vwx
a1b2c3d4e5f
treb7uchet
`))
	utils.ExpectInt(t, result, 142)
}

func TestPart1Actual(t *testing.T) {
	result := Part1(utils.ReadInputFromFile("input.txt"))
	utils.ExpectInt(t, result, 55607)
}

func TestPart2Given(t *testing.T) {
	result := Part2(utils.ToInput(`
two1nine
eightwothree
abcone2threexyz
xtwone3four
4nineeightseven2
zoneight234
7pqrstsixteen
`))
	utils.ExpectInt(t, result, 281)
}

func TestPart2Actual(t *testing.T) {
	result := Part2(utils.ReadInputFromFile("input.txt"))
	utils.ExpectInt(t, result, 55291)
}
