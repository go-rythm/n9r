package memo

// https://www.lintcode.com/problem/triangle/

/**
 * @param triangle: a list of lists of integers
 * @return: An integer, minimum path sum
 */

func MinimumTotal(triangle [][]int) int {
	return divideConquer(triangle, 0, 0, map[struct{ x, y int }]int{})
}

func divideConquer(triangle [][]int, x, y int, memo map[struct{ x, y int }]int) int {
	if x == len(triangle) {
		return 0
	}

	if v, ok := memo[struct {
		x int
		y int
	}{x: x, y: y}]; ok {
		return v
	}

	l := divideConquer(triangle, x+1, y, memo)
	r := divideConquer(triangle, x+1, y+1, memo)
	minimum := min(l, r) + triangle[x][y]
	memo[struct {
		x int
		y int
	}{x: x, y: y}] = minimum
	return minimum
}
