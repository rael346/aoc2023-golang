package day8

import (
	"testing"

	"github.com/rael346/aoc2023-golang/internal/utils"
)

var givenInput = `
RL

AAA = (BBB, CCC)
BBB = (DDD, EEE)
CCC = (ZZZ, GGG)
DDD = (DDD, DDD)
EEE = (EEE, EEE)
GGG = (GGG, GGG)
ZZZ = (ZZZ, ZZZ)
`

func TestPart1Given1(t *testing.T) {
	result := Part1(utils.ToInput(givenInput))
	utils.ExpectInt(t, result, 2)
}

func TestPart1Given2(t *testing.T) {
	result := Part1(utils.ToInput(`
LLR

AAA = (BBB, BBB)
BBB = (AAA, ZZZ)
ZZZ = (ZZZ, ZZZ)
`))
	utils.ExpectInt(t, result, 6)
}

func TestPart1Actual(t *testing.T) {
	result := Part1(utils.ReadInputFromFile("input.txt"))
	utils.ExpectInt(t, result, 13019)
}

func TestPart2Given(t *testing.T) {
	result := Part2(utils.ToInput(`
LR

11A = (11B, XXX)
11B = (XXX, 11Z)
11Z = (11B, XXX)
22A = (22B, XXX)
22B = (22C, 22C)
22C = (22Z, 22Z)
22Z = (22B, 22B)
XXX = (XXX, XXX)
`))
	utils.ExpectInt(t, result, 6)
}

func TestPart2Actual(t *testing.T) {
	result := Part2(utils.ReadInputFromFile("input.txt"))
	utils.ExpectInt(t, result, 13524038372771)
}
