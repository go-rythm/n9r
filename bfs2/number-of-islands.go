package bfs2

// https://www.lintcode.com/problem/433/

/**
 * @param grid: a boolean 2D matrix
 * @return: an integer
 */
func NumIslands(grid [][]bool) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	islands := 0
	visited := map[Coord]bool{}

	for i, row := range grid {
		for j, v := range row {
			p := Coord{i, j}
			if v && !visited[p] {
				islands++
				bfs(grid, p, visited)
			}
		}
	}

	return islands
}

func bfs(grid [][]bool, p Coord, visited map[Coord]bool) {
	queue := []Coord{p}
	visited[p] = true
	for len(queue) > 0 {
		curP := queue[0]
		queue = queue[1:]
		for _, d := range Directions {
			nextP := Coord{
				x: curP.x + d.x,
				y: curP.y + d.y,
			}
			if !nextP.isValid(grid, visited) {
				continue
			}
			queue = append(queue, nextP)
			visited[nextP] = true
		}
	}
}

// 上下左右
var Directions = []Coord{
	{-1, 0},
	{1, 0},
	{0, -1},
	{0, 1},
}

type Coord struct {
	x, y int
}

func (p Coord) isValid(grid [][]bool, visited map[Coord]bool) bool {
	n, m := len(grid), len(grid[0])
	if p.x < 0 || p.x >= n {
		return false
	}
	if p.y < 0 || p.y >= m {
		return false
	}
	if visited[p] {
		return false
	}
	return grid[p.x][p.y]
}
