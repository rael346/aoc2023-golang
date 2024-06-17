package day10

import (
	"testing"

	"github.com/rael346/aoc2023-golang/internal/utils"
)

var givenInput1 = `
.....
.S-7.
.|.|.
.L-J.
.....
`

var givenInput2 = `
..F7.
.FJ|.
SJ.L7
|F--J
LJ...
`

func TestPart1Given1(t *testing.T) {
	result := Part1(utils.ToInput(givenInput1))
	utils.ExpectInt(t, result, 4)
}

func TestPart1Given2(t *testing.T) {
	result := Part1(utils.ToInput(givenInput2))
	utils.ExpectInt(t, result, 8)
}

func TestPart1Actual(t *testing.T) {
	result := Part1(utils.ReadInputFromFile("input.txt"))
	utils.ExpectInt(t, result, 6717)
}

var part2GivenInputSimple1 = `
...........
.S-------7.
.|F-----7|.
.||.....||.
.||.....||.
.|L-7.F-J|.
.|..|.|..|.
.L--J.L--J.
...........
`

var part2GivenInputSimple2 = `
..........
.S------7.
.|F----7|.
.||....||.
.||....||.
.|L-7F-J|.
.|..||..|.
.L--JL--J.
..........
`

func TestPart2GivenSimple(t *testing.T) {
	result1 := Part2(utils.ToInput(part2GivenInputSimple1))
	utils.ExpectInt(t, result1, 4)

	result2 := Part2(utils.ToInput(part2GivenInputSimple2))
	utils.ExpectInt(t, result2, 4)
}

var part2GivenInputLarge1 = `
.F----7F7F7F7F-7....
.|F--7||||||||FJ....
.||.FJ||||||||L7....
FJL7L7LJLJ||LJ.L-7..
L--J.L7...LJS7F-7L7.
....F-J..F7FJ|L7L7L7
....L7.F7||L7|.L7L7|
.....|FJLJ|FJ|F7|.LJ
....FJL-7.||.||||...
....L---J.LJ.LJLJ...
`

var part2GivenInputLarge2 = `
FF7FSF7F7F7F7F7F---7
L|LJ||||||||||||F--J
FL-7LJLJ||||||LJL-77
F--JF--7||LJLJ7F7FJ-
L---JF-JLJ.||-FJLJJ7
|F|F-JF---7F7-L7L|7|
|FFJF7L7F-JF7|JL---7
7-L-JL7||F7|L7F-7F7|
L.L7LFJ|||||FJL7||LJ
L7JLJL-JLJLJL--JLJ.L
`

func TestPart2GivenLarge(t *testing.T) {
	result1 := Part2(utils.ToInput(part2GivenInputLarge1))
	utils.ExpectInt(t, result1, 8)

	result2 := Part2(utils.ToInput(part2GivenInputLarge2))
	utils.ExpectInt(t, result2, 10)
}

func TestPart2Actual(t *testing.T) {
	result := Part2(utils.ReadInputFromFile("input.txt"))
	utils.ExpectInt(t, result, 381)
}
