package dfs

import (
	"sort"
)

// https://www.lintcode.com/problem/subsets

/**
 * @param nums: A set of numbers
 * @return: A list of lists
 *          we will sort your return value in output
 */
func Subsets(nums []int) [][]int {
	sorted := append([]int{}, nums...)
	sort.Ints(sorted)
	res := new([][]int)
	dfs(sorted, 0, []int{}, res)
	return *res
}

func dfs(nums []int, index int, subset []int, res *[][]int) {
	if index == len(nums) {
		*res = append(*res, append([]int{}, subset...))
		return
	}

	dfs(nums, index+1, subset, res)

	subset = append(subset, nums[index])
	dfs(nums, index+1, subset, res)
}

func Subsets2(nums []int) [][]int {
	sorted := append([]int{}, nums...)
	sort.Ints(sorted)
	res := new([][]int)
	dfs2(sorted, 0, []int{}, res)
	return *res
}

func dfs2(nums []int, index int, subset []int, res *[][]int) {
	*res = append(*res, append([]int{}, subset...))
	for i := index; i < len(nums); i++ {
		// dfs2(nums, i+1, append(subset, nums[i]), res)

		subset = append(subset, nums[i])
		dfs2(nums, i+1, subset, res)
		subset = subset[:len(subset)-1]
	}
}
