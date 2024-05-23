package day3

import (
	"strconv"
	"unicode"
)

// for each character,
// if it is a digit then add it to a queue
// also if that character is adjacent to a symbol (if so then that current number is a part number)
// when the current character is not a digit then empty the queue,
// if the current number is a part number then add to the sum
func Part1(input []string) int {
	sum := 0
	tileMap := [][]rune{}
	for _, line := range input {
		tileMap = append(tileMap, []rune(line))
	}

	queue := []rune{}
	isPartNum := false
	for r, line := range input {
		for c, char := range line {
			if unicode.IsDigit(char) {
				queue = append(queue, char)
				if isPartNumber(tileMap, r, c) {
					isPartNum = true
				}
				continue
			}

			if isPartNum {
				partNumber, _ := strconv.Atoi(string(queue))
				sum += partNumber
			}

			queue = []rune{}
			isPartNum = false
		}
	}

	return sum
}

var deltas = [][]int{
	{-1, -1},
	{-1, 0},
	{-1, 1},

	{0, -1},
	{0, 1},

	{1, -1},
	{1, 0},
	{1, 1},
}

// check if the surrounding tile contains a symbol
func isPartNumber(tileMap [][]rune, r int, c int) bool {
	numRow, numCol := len(tileMap), len(tileMap[0])
	for _, delta := range deltas {
		newRow, newCol := r+delta[0], c+delta[1]
		if newRow < 0 || newRow > numRow-1 || newCol < 0 || newCol > numCol-1 {
			continue
		}

		char := tileMap[newRow][newCol]
		if char == '.' || ('0' <= char && char <= '9') {
			continue
		}

		return true
	}

	return false
}

// similar approach but now also keeping track of the gear location and their numbers in a map
// the key is a string of the format "rowIndex | colIndex"
func Part2(input []string) int {
	tileMap := [][]rune{}
	for _, line := range input {
		tileMap = append(tileMap, []rune(line))
	}

	queue := []rune{}
	gearMap := map[string][]int{}
	isGearNum := false
	gearRow := -1
	gearCol := -1

	for r, line := range input {
		for c, char := range line {
			if unicode.IsDigit(char) {
				queue = append(queue, char)
				if isGear, newGearRow, newGearCol := isGearNumber(tileMap, r, c); isGear {
					isGearNum = true
					gearRow, gearCol = newGearRow, newGearCol
				}
				continue
			}

			if isGearNum {
				partNumber, _ := strconv.Atoi(string(queue))
				gearKey := strconv.Itoa(gearRow) + " | " + strconv.Itoa(gearCol)
				if _, ok := gearMap[gearKey]; !ok {
					gearMap[gearKey] = []int{}
				}
				gearMap[gearKey] = append(gearMap[gearKey], partNumber)
			}

			queue = []rune{}
			isGearNum = false
			gearRow = -1
			gearCol = -1
		}
	}

	sum := 0
	for _, gearNums := range gearMap {
		if len(gearNums) != 2 {
			continue
		}

		sum += gearNums[0] * gearNums[1]
	}

	return sum
}

// check if the surrounding tiles contain the gear (asterisk symbol),
// if so then return the coord for that gear
func isGearNumber(tileMap [][]rune, r int, c int) (bool, int, int) {
	numRow, numCol := len(tileMap), len(tileMap[0])
	for _, delta := range deltas {
		newRow, newCol := r+delta[0], c+delta[1]
		if newRow < 0 || newRow > numRow-1 || newCol < 0 || newCol > numCol-1 {
			continue
		}

		char := tileMap[newRow][newCol]
		if char == '*' {
			return true, newRow, newCol
		}
	}

	return false, -1, -1
}
