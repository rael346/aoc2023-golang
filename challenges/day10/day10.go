package day10

import (
	"math"
)

type Point struct {
	x int
	y int
}

var TOP = Point{0, -1}
var BOTTOM = Point{0, 1}
var LEFT = Point{-1, 0}
var RIGHT = Point{1, 0}

func (p1 Point) add(p2 Point) Point {
	return Point{p1.x + p2.x, p1.y + p2.y}
}

func (p1 Point) isEqual(p2 Point) bool {
	return p1.x == p2.x && p1.y == p2.y
}

var pipeMap = map[rune][2]Point{
	'|': {TOP, BOTTOM},
	'-': {LEFT, RIGHT},
	'F': {BOTTOM, RIGHT},
	'L': {TOP, RIGHT},
	'7': {BOTTOM, LEFT},
	'J': {TOP, LEFT},
}

// stack item
type Item struct {
	// the previous point that create this item
	// mainly to check where the pipe input was
	prev Point
	// the current ponit this item is pointed to
	curr Point
	// the path that leads from the start point to the current point
	path []Point
}

func Part1(input []string) int {
	numRow, numCol := len(input), len(input[0])

	// build the grid (mainly for convenience)
	// bruteforce getting the start point
	tileMap := make([][]rune, numRow)
	start := Point{0, 0}
	for r, line := range input {
		row := make([]rune, numCol)
		for c, tile := range line {
			row[c] = tile

			if tile == 'S' {
				start.y, start.x = r, c
			}
		}
		tileMap[r] = row
	}

	return len(findLoop(start, tileMap)) / 2
}

// DFS from the start point,
// if a path reach the start point again then it is the loop
func findLoop(start Point, tileMap [][]rune) []Point {
	numRow, numCol := len(tileMap), len(tileMap[0])
	stack := []Item{
		{start, start.add(TOP), []Point{start}},
		{start, start.add(BOTTOM), []Point{start}},
		{start, start.add(LEFT), []Point{start}},
		{start, start.add(RIGHT), []Point{start}},
	}

	loop := []Point{}
	for len(stack) > 0 {
		currItem := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if currItem.curr.y < 0 || currItem.curr.y > numRow-1 || currItem.curr.x < 0 || currItem.curr.x > numCol-1 {
			continue
		}

		currSymbol := tileMap[currItem.curr.y][currItem.curr.x]
		if currSymbol == '.' {
			continue
		}

		if currSymbol == 'S' {
			loop = currItem.path
			break
		}

		pipeOut1 := currItem.curr.add(pipeMap[currSymbol][0])
		pipeOut2 := currItem.curr.add(pipeMap[currSymbol][1])

		if pipeOut1.isEqual(currItem.prev) {
			stack = append(stack, Item{currItem.curr, pipeOut2, append(currItem.path, currItem.curr)})
		}

		if pipeOut2.isEqual(currItem.prev) {
			stack = append(stack, Item{currItem.curr, pipeOut1, append(currItem.path, currItem.curr)})
		}
	}

	return loop
}

func Part2(input []string) int {
	numRow, numCol := len(input), len(input[0])

	// brute force find the start point
	tileMap := make([][]rune, numRow)
	start := Point{0, 0}
	for r, line := range input {
		row := make([]rune, numCol)
		for c, tile := range line {
			row[c] = tile

			if tile == 'S' {
				start.y, start.x = r, c
			}
		}
		tileMap[r] = row
	}

	// get the vertices list
	loop := findLoop(start, tileMap)
	vertices := []Point{}
	for _, p := range loop {
		tile := tileMap[p.y][p.x]
		if tile == '-' || tile == '|' {
			continue
		}
		vertices = append(vertices, p)
	}

	// Shoelace formula for area of a polygon
	vertices = append(vertices, start)
	twiceArea := 0
	for i := range len(vertices) - 1 {
		twiceArea += vertices[i].x*vertices[i+1].y - vertices[i].y*vertices[i+1].x
	}

	// Pick's Theorem for relating area of a int-vertices polygon,
	// number of integer points within the polygon and on the boundary
	return int(math.Abs(float64(twiceArea)))/2 + 1 - len(loop)/2
}
