package day9

import (
	"strconv"
	"strings"

	"github.com/rael346/aoc2023-golang/internal/utils"
)

// for each sequence predict the next value
func Part1(input []string) int {
	sum := 0
	for _, line := range input {
		sequenceStrs := strings.Fields(line)
		sequence := make([]int, len(sequenceStrs))
		for i, str := range sequenceStrs {
			parsedNum, _ := strconv.Atoi(str)
			sequence[i] = parsedNum
		}
		sum += predictNext(sequence)
	}
	return sum
}

// calculate the history and add the final value to the prediction
func predictNext(sequence []int) int {
	prediction := 0
	currentSeq := sequence
	for !utils.All(currentSeq, isZero) {
		prediction += currentSeq[len(currentSeq)-1]
		newSequence := make([]int, len(currentSeq)-1)
		for i := 0; i < len(currentSeq)-1; i++ {
			newSequence[i] = currentSeq[i+1] - currentSeq[i]
		}
		currentSeq = newSequence
	}
	return prediction
}

func isZero(num int) bool {
	return num == 0
}

// same approach but now predict backwards
func Part2(input []string) int {
	sum := 0
	for _, line := range input {
		sequenceStrs := strings.Fields(line)
		sequence := make([]int, len(sequenceStrs))
		for i, str := range sequenceStrs {
			parsedNum, _ := strconv.Atoi(str)
			sequence[i] = parsedNum
		}

		sum += predictPrev(sequence)
	}
	return sum
}

// the backward prediction is calculate by
// getting the first number in a history
// and alternating the sign for each history
func predictPrev(sequence []int) int {
	prediction := 0
	currentSeq := sequence
	historyCount := 0
	for !utils.All(currentSeq, isZero) {
		if historyCount%2 == 0 {
			prediction += currentSeq[0]
		} else {
			prediction -= currentSeq[0]
		}

		newSequence := make([]int, len(currentSeq)-1)
		for i := 0; i < len(currentSeq)-1; i++ {
			newSequence[i] = currentSeq[i+1] - currentSeq[i]
		}
		currentSeq = newSequence
		historyCount += 1
	}
	return prediction
}
