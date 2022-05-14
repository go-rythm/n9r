package bfs2

// https://www.lintcode.com/problem/611/

// Definition for a point.
type Point struct {
	X, Y int
}

func isValid(p Point, grid [][]bool, cellToDisMap map[Point]int) bool {
	n, m := len(grid), len(grid[0])
	if p.X < 0 || p.X >= n {
		return false
	}
	if p.Y < 0 || p.Y >= m {
		return false
	}
	if grid[p.X][p.Y] {
		return false
	}
	if _, ok := cellToDisMap[p]; ok {
		return false
	}
	return true
}

/**
 * @param grid: a chessboard included 0 (false) and 1 (true)
 * @param source: a point
 * @param destination: a point
 * @return: the shortest path
 */
func ShortestPath(grid [][]bool, source *Point, destination *Point) int {
	if *source == *destination {
		return 0
	}

	queue := []Point{*source}
	cellToDisMap := map[Point]int{
		*source: 0,
	}

	for len(queue) > 0 {
		curP := queue[0]
		queue = queue[1:]
		for _, offset := range Offsets {
			newP := Point{
				X: curP.X + offset.X,
				Y: curP.Y + offset.Y,
			}
			if !isValid(newP, grid, cellToDisMap) {
				continue
			}
			if newP == *destination {
				return cellToDisMap[curP] + 1
			}
			queue = append(queue, newP)
			cellToDisMap[newP] = cellToDisMap[curP] + 1
		}
	}
	return -1
}

var Offsets = []Point{
	{-2, -1},
	{-2, 1},
	{2, 1},
	{2, -1},
	{1, 2},
	{1, -2},
	{-1, 2},
	{-1, -2},
}

func ShortestPath1(grid [][]bool, source *Point, destination *Point) int {
	queue := []Point{*source}
	cellToDisMap := map[Point]int{
		*source: 0,
	}

	for len(queue) > 0 {
		curP := queue[0]
		queue = queue[1:]
		if curP == *destination {
			return cellToDisMap[curP]
		}
		for _, offset := range Offsets {
			newP := Point{
				X: curP.X + offset.X,
				Y: curP.Y + offset.Y,
			}
			if !isValid(newP, grid, cellToDisMap) {
				continue
			}
			queue = append(queue, newP)
			cellToDisMap[newP] = cellToDisMap[curP] + 1
		}
	}
	return -1
}
