package day7

import (
	"testing"

	"github.com/rael346/aoc2023-golang/internal/utils"
)

var givenInput = `
32T3K 765
T55J5 684
KK677 28
KTJJT 220
QQQJA 483
`

func TestPart1Given(t *testing.T) {
	result := Part1(utils.ToInput(givenInput))
	utils.ExpectInt(t, result, 6440)
}

func TestPart1Actual(t *testing.T) {
	result := Part1(utils.ReadInputFromFile("input.txt"))
	utils.ExpectInt(t, result, 246163188)
}

func TestPart2Given(t *testing.T) {
	result := Part2(utils.ToInput(givenInput))
	utils.ExpectInt(t, result, 5905)
}

func TestPart2Actual(t *testing.T) {
	result := Part2(utils.ReadInputFromFile("input.txt"))
	utils.ExpectInt(t, result, 245794069)
}

func TestClassifyHand(t *testing.T) {
	utils.ExpectInt(t, getHandTypePart2("KKKKK"), FIVE_OF_A_KIND)
	utils.ExpectInt(t, getHandTypePart2("KKKKJ"), FIVE_OF_A_KIND)
	utils.ExpectInt(t, getHandTypePart2("KKKJJ"), FIVE_OF_A_KIND)
	utils.ExpectInt(t, getHandTypePart2("KKJJJ"), FIVE_OF_A_KIND)
	utils.ExpectInt(t, getHandTypePart2("KJJJJ"), FIVE_OF_A_KIND)
	utils.ExpectInt(t, getHandTypePart2("JJJJJ"), FIVE_OF_A_KIND)

	utils.ExpectInt(t, getHandTypePart2("KKJ54"), THREE_OF_A_KIND)
	utils.ExpectInt(t, getHandTypePart2("A23JQ"), ONE_PAIR)
}
