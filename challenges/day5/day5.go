package day5

import (
	"math"
	"strconv"
	"strings"
)

func Part1(input []string) int {
	seedStrs := strings.Fields(input[0])[1:]
	srcMap := map[int]int{}
	for _, seedStr := range seedStrs {
		seed, _ := strconv.Atoi(seedStr)
		srcMap[seed] = seed
	}

	for _, line := range input[2:] {
		if strings.Contains(line, "map") {
			continue
		}

		if line == "" {
			newSrcMap := map[int]int{}
			for _, dest := range srcMap {
				newSrcMap[dest] = dest
			}
			srcMap = newSrcMap
			continue
		}

		mapping := strings.Fields(line)
		destStart, _ := strconv.Atoi(mapping[0])
		srcStart, _ := strconv.Atoi(mapping[1])
		rangeLen, _ := strconv.Atoi(mapping[2])

		for src := range srcMap {
			if srcStart <= src && src < srcStart+rangeLen {
				srcMap[src] = src - srcStart + destStart
			}
		}
	}

	minLoc := math.MaxInt

	for _, dest := range srcMap {
		minLoc = min(minLoc, dest)
	}

	return minLoc
}

// similar to part 1, but instead of only mapping
// a single number, need to map a range
// For each range, get the intersect and the unmapped section (non-intersect)
// of the src
func Part2(input []string) int {
	seedStrs := strings.Fields(input[0])[1:]
	srcs := make([][2]int, len(seedStrs)/2)
	for i := 0; i < len(seedStrs); i += 2 {
		seed, _ := strconv.Atoi(seedStrs[i])
		rangeLen, _ := strconv.Atoi(seedStrs[i+1])
		srcs[i/2] = [2]int{seed, seed + rangeLen - 1}
	}

	unmappedSrcs := [][2]int{}
	dest := [][2]int{}
	for _, line := range input[2:] {
		if strings.Contains(line, "map") {
			continue
		}

		if line == "" {
			srcs = append(srcs, dest...)
			dest = [][2]int{}
			continue
		}

		mapping := strings.Fields(line)
		destStart, _ := strconv.Atoi(mapping[0])
		srcMapStart, _ := strconv.Atoi(mapping[1])
		srcMapRange, _ := strconv.Atoi(mapping[2])
		srcMapEnd := srcMapStart + srcMapRange - 1

		for _, src := range srcs {
			newSrc, unmapped := Intersect(src, [2]int{srcMapStart, srcMapEnd}, destStart)
			if newSrc[0] != -1 && newSrc[1] != -1 {
				dest = append(dest, newSrc)
			}
			unmappedSrcs = append(unmappedSrcs, unmapped...)
		}
		srcs = unmappedSrcs
		unmappedSrcs = [][2]int{}
	}

	minLoc := math.MaxInt
	for _, srcRange := range append(srcs, dest...) {
		minLoc = min(minLoc, srcRange[0])
	}

	return minLoc
}

// Calculate the intersect and the unmapped sections between the two range
func Intersect(src [2]int, srcMap [2]int, destStart int) ([2]int, [][2]int) {
	srcStart, srcEnd := src[0], src[1]
	srcMapStart, srcMapEnd := srcMap[0], srcMap[1]
	unmapped := make([][2]int, 0, 2)

	// if there is no overlap between the two range
	// the src range is unmapped
	if srcStart > srcMapEnd || srcMapStart > srcEnd {
		unmapped = append(unmapped, src)
		return [2]int{-1, -1}, unmapped
	}

	var newSrcStart int
	if srcMapStart <= srcStart {
		newSrcStart = srcStart - srcMapStart + destStart
	} else {
		newSrcStart = destStart
		unmapped = append(unmapped, [2]int{srcStart, srcMapStart - 1})
	}

	var newSrcEnd int
	if srcEnd <= srcMapEnd {
		newSrcEnd = srcEnd - srcMapStart + destStart
	} else {
		newSrcEnd = srcMapEnd - srcMapStart + destStart
		unmapped = append(unmapped, [2]int{srcMapEnd + 1, srcEnd})
	}

	return [2]int{newSrcStart, newSrcEnd}, unmapped
}
