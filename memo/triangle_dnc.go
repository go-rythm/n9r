package memo

// https://www.lintcode.com/problem/triangle/

/**
 * @param triangle: a list of lists of integers
 * @return: An integer, minimum path sum
 */

func MinimumTotal2(triangle [][]int) int {
	return dnc(triangle, 0, 0)
}

func dnc(triangle [][]int, x, y int) int {
	if x == len(triangle) {
		return 0
	}

	l := dnc(triangle, x+1, y)
	r := dnc(triangle, x+1, y+1)
	return min(l, r) + triangle[x][y]
}
