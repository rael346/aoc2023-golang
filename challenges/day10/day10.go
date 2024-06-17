package day10

// row, column
type Point [2]int

var TOP = Point{-1, 0}
var BOTTOM = Point{1, 0}
var LEFT = Point{0, -1}
var RIGHT = Point{0, 1}

func (p1 Point) add(p2 Point) Point {
	return Point{p1[0] + p2[0], p1[1] + p2[1]}
}

func (p1 Point) isEqual(p2 Point) bool {
	return p1[0] == p2[0] && p1[1] == p2[1]
}

var pipeMap = map[rune][2]Point{
	'|': {TOP, BOTTOM},
	'-': {LEFT, RIGHT},
	'F': {BOTTOM, RIGHT},
	'L': {TOP, RIGHT},
	'7': {BOTTOM, LEFT},
	'J': {TOP, LEFT},
}

type Item struct {
	prev Point
	curr Point
	path []Point
}

func Part1(input []string) int {
	numRow, numCol := len(input), len(input[0])
	tileMap := [][]rune{}
	for range numRow {
		tileMap = append(tileMap, make([]rune, numCol))
	}

	start := Point{0, 0}
	for r, line := range input {
		for c, tile := range line {
			tileMap[r][c] = tile
			if tile == 'S' {
				start[0], start[1] = r, c
			}
		}
	}

	return len(findLoop(start, tileMap)) / 2
}

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

		if currItem.curr[0] < 0 || currItem.curr[0] > numRow-1 || currItem.curr[1] < 0 || currItem.curr[1] > numCol-1 {
			continue
		}

		currSymbol := tileMap[currItem.curr[0]][currItem.curr[1]]
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
	tileMap := [][]rune{}
	for range numRow {
		tileMap = append(tileMap, make([]rune, numCol))
	}

	start := Point{0, 0}
	for r, line := range input {
		for c, tile := range line {
			tileMap[r][c] = tile
			if tile == 'S' {
				start[0], start[1] = r, c
			}
		}
	}

	loop := findLoop(start, tileMap)
	chain := []Point{}
	for _, p := range loop {
		tile := tileMap[p[0]][p[1]]
		if tile == '-' || tile == '|' {
			continue
		}
		chain = append(chain, p)
	}

	return 0
}
