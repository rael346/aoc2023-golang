package day11

import "math"

type Point struct {
	x int
	y int
}

func (p1 Point) manhattanDist(p2 Point) int {
	return int(math.Abs(float64(p1.x-p2.x)) + math.Abs(float64(p1.y-p2.y)))
}

func Part1(input []string) int {
	return totalDist(input, 1, 1)
}

func totalDist(input []string, emptyRowDist int, emptyColDist int) int {
	numRow, numCol := len(input), len(input[0])
	tileMap := make([][]rune, 0, numRow)
	galaxies := []Point{}

	for r, line := range input {
		row := make([]rune, 0, numCol)
		for c, tile := range line {
			if tile == '#' {
				galaxies = append(galaxies, Point{c, r})
			}
			row = append(row, tile)
		}
		tileMap = append(tileMap, row)
	}

	yMap := make([]int, numRow)
	for y := range numRow {
		isEmpty := true
		for x := range numCol {
			if tileMap[y][x] == '#' {
				isEmpty = false
				break
			}
		}

		if isEmpty {
			yMap[y] = yMap[max(y-1, 0)] + emptyRowDist
		} else {
			yMap[y] = yMap[max(y-1, 0)]
		}
	}

	xMap := make([]int, numCol)
	for x := range numCol {
		isEmpty := true
		for y := range numRow {
			if tileMap[y][x] == '#' {
				isEmpty = false
				break
			}
		}

		if isEmpty {
			xMap[x] = xMap[max(x-1, 0)] + emptyColDist
		} else {
			xMap[x] = xMap[max(x-1, 0)]
		}
	}

	sum := 0
	for _, g1 := range galaxies {
		for _, g2 := range galaxies {
			mappedG1 := Point{g1.x + xMap[g1.x], g1.y + yMap[g1.y]}
			mappedG2 := Point{g2.x + xMap[g2.x], g2.y + yMap[g2.y]}
			sum += mappedG1.manhattanDist(mappedG2)
		}
	}
	return sum / 2
}

func Part2(input []string) int {
	return totalDist(input, 1000000-1, 1000000-1)
}
