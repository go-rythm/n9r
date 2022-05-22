package dfs

import (
	"sort"
)

// https://www.lintcode.com/problem/subsets-ii

/**
 * @param nums: A set of numbers.
 * @return: A list of lists. All valid subsets.
 *          we will sort your return value in output
 */
func SubsetsWithDup(nums []int) [][]int {
	sorted := append([]int{}, nums...)
	sort.Ints(sorted)
	res := new([][]int)
	dfsWithDup(sorted, 0, []int{}, res)
	return *res
}

func dfsWithDup(nums []int, index int, subset []int, res *[][]int) {
	*res = append(*res, append([]int{}, subset...))
	for i := index; i < len(nums); i++ {
		if i != index && nums[i] == nums[i-1] {
			continue
		}
		subset = append(subset, nums[i])
		dfsWithDup(nums, i+1, subset, res)
		subset = subset[:len(subset)-1]
	}
}
