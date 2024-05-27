package day6

import (
	"strconv"
	"strings"
)

func Part1(input []string) int {
	time := strings.Fields(input[0])[1:]
	distance := strings.Fields(input[1])[1:]

	product := 1
	for i := 0; i < len(time); i++ {
		t, _ := strconv.Atoi(time[i])
		d, _ := strconv.Atoi(distance[i])

		count := 0
		for speed := 1; speed < t; speed++ {
			if (t-speed)*speed > d {
				count += 1
			}
		}
		product *= count
	}
	return product
}

func Part2(input []string) int {
	timeArr := strings.Fields(input[0])[1:]
	distanceArr := strings.Fields(input[1])[1:]
	time, _ := strconv.Atoi(strings.Join(timeArr, ""))
	distance, _ := strconv.Atoi(strings.Join(distanceArr, ""))

	count := 0
	for speed := 1; speed < time; speed++ {
		if (time-speed)*speed > distance {
			count += 1
		}
	}
	return count
}
