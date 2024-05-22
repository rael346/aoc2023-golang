package day1

import (
	"strconv"
	"strings"
	"unicode"
)

func Part1(input []string) int {
	sum := 0

	for _, line := range input {
		nums := []string{}
		for _, c := range line {
			if unicode.IsDigit(c) {
				nums = append(nums, string(c))
			}
		}

		first, last := nums[0], nums[len(nums)-1]
		calibrated_val, _ := strconv.Atoi(first + last)
		sum += calibrated_val
	}

	return sum
}

var numString = []string{
	"zero",
	"one",
	"two",
	"three",
	"four",
	"five",
	"six",
	"seven",
	"eight",
	"nine",
}

func Part2(input []string) int {
	sum := 0

	for _, line := range input {
		nums := []string{}
		for i, c := range line {
			if unicode.IsDigit(c) {
				nums = append(nums, string(c))
			}

			if unicode.IsLetter(c) {
				for num, numStr := range numString {
					if strings.HasPrefix(line[i:], numStr) {
						nums = append(nums, strconv.Itoa(num))
					}
				}
			}
		}

		first, last := nums[0], nums[len(nums)-1]
		calibrated_val, _ := strconv.Atoi(first + last)
		sum += calibrated_val
	}

	return sum
}
