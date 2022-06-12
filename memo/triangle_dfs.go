package memo

import "math"

// https://www.lintcode.com/problem/triangle/

/**
 * @param triangle: a list of lists of integers
 * @return: An integer, minimum path sum
 */

var minimum int

func MinimumTotal1(triangle [][]int) int {
	minimum = math.MaxInt
	traverse(triangle, 0, 0, 0)
	return minimum
}

func traverse(triangle [][]int, x, y, pathSum int) {
	if x == len(triangle) {
		minimum = min(minimum, pathSum)
		return
	}

	traverse(triangle, x+1, y, pathSum+triangle[x][y])
	traverse(triangle, x+1, y+1, pathSum+triangle[x][y])
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
