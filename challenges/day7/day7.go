package day7

import (
	"slices"
	"strconv"
	"strings"
)

type Entry struct {
	hand     string
	bid      int
	handType int
}

const (
	FIVE_OF_A_KIND  = 6
	FOUR_OF_A_KIND  = 5
	FULL_HOUSE      = 4
	THREE_OF_A_KIND = 3
	TWO_PAIR        = 2
	ONE_PAIR        = 1
	HIGH_CARD       = 0
)

var labelMap = map[rune]int{
	'A': 14,
	'K': 13,
	'Q': 12,
	'J': 11,
	'T': 10,
	'9': 9,
	'8': 8,
	'7': 7,
	'6': 6,
	'5': 5,
	'4': 4,
	'3': 3,
	'2': 2,
}

// sort the entries of [hand bid] by the handType
func Part1(input []string) int {
	entries := make([]Entry, 0, len(input))
	for _, line := range input {
		parseLine := strings.Fields(line)
		hand := parseLine[0]
		bid, _ := strconv.Atoi(parseLine[1])
		entries = append(entries, Entry{hand, bid, getHandTypePart1(hand)})
	}

	slices.SortStableFunc(entries, func(a, b Entry) int {
		handTypeDiff := a.handType - b.handType
		if handTypeDiff != 0 {
			return handTypeDiff
		}

		// compare which hand has the higher valued card
		handA, handB := []rune(a.hand), []rune(b.hand)
		for i := 0; i < len(handA); i++ {
			charA, charB := handA[i], handB[i]
			charDiff := labelMap[charA] - labelMap[charB]

			if charDiff != 0 {
				return charDiff
			}
		}
		return 0
	})

	sum := 0
	for i, entry := range entries {
		sum += entry.bid * (i + 1)
	}
	return sum
}

func getHandTypePart1(hand string) int {
	charCount := map[rune]int{}
	for _, c := range hand {
		charCount[c] += 1
	}

	tripleCount := 0
	pairCount := 0
	for _, count := range charCount {
		switch count {
		case 5:
			return FIVE_OF_A_KIND
		case 4:
			return FOUR_OF_A_KIND
		case 3:
			tripleCount += 1
		case 2:
			pairCount += 1
		}
	}

	switch {
	case tripleCount == 1 && pairCount == 1:
		return FULL_HOUSE
	case tripleCount == 1:
		return THREE_OF_A_KIND
	case pairCount == 2:
		return TWO_PAIR
	case pairCount == 1:
		return ONE_PAIR
	default:
		return HIGH_CARD
	}
}

// still sorting but change up the value of J and the getHandType function
func Part2(input []string) int {
	labelMap['J'] = 1

	entries := make([]Entry, 0, len(input))
	for _, line := range input {
		parseLine := strings.Fields(line)
		hand := parseLine[0]
		bid, _ := strconv.Atoi(parseLine[1])
		entries = append(entries, Entry{hand, bid, getHandTypePart2(hand)})
	}

	slices.SortStableFunc(entries, func(a, b Entry) int {
		handTypeDiff := a.handType - b.handType
		if handTypeDiff != 0 {
			return handTypeDiff
		}

		handA, handB := []rune(a.hand), []rune(b.hand)
		for i := 0; i < len(handA); i++ {
			charA, charB := handA[i], handB[i]
			charDiff := labelMap[charA] - labelMap[charB]

			if charDiff != 0 {
				return charDiff
			}
		}
		return 0
	})

	sum := 0
	for i, entry := range entries {
		sum += entry.bid * (i + 1)
	}
	return sum
}

// separate the J count from the regular char count
// to add the J count to the max count later on
// since J can morph into any other char
func getHandTypePart2(hand string) int {
	charCount := map[rune]int{}
	jCount := 0
	for _, c := range hand {
		if c == 'J' {
			jCount += 1
		} else {
			charCount[c] += 1
		}
	}

	tripleCount := 0
	pairCount := 0
	maxCount := 0
	for _, count := range charCount {
		maxCount = max(maxCount, count)
		switch count {
		case 3:
			tripleCount += 1
		case 2:
			pairCount += 1
		}
	}

	switch maxCount + jCount {
	case 5:
		return FIVE_OF_A_KIND
	case 4:
		return FOUR_OF_A_KIND
	case 3:
		if pairCount == 2 || (maxCount == 3 && pairCount == 1) {
			return FULL_HOUSE
		}

		return THREE_OF_A_KIND
	case 2:
		if pairCount == 2 {
			return TWO_PAIR
		}
		return ONE_PAIR
	default:
		return HIGH_CARD
	}
}
