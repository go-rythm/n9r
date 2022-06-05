package dfs2

import "sort"

// https://www.lintcode.com/problem/combination-sum/

/**
 * @param candidates: A list of integers
 * @param target: An integer
 * @return: A list of lists of integers
 *          we will sort your return value in output
 */
func CombinationSum(candidates []int, target int) [][]int {
	subsets := [][]int{}
	if len(candidates) == 0 {
		return subsets
	}
	dict := removeDuplicatesAndSort(candidates)
	dfsCombinationSum(dict, target, 0, nil, &subsets)
	return subsets
}

func dfsCombinationSum(dict []int, target int, idx int, subset []int, subsets *[][]int) {
	if target == 0 {
		*subsets = append(*subsets, append([]int{}, subset...))
	}
	for i := idx; i < len(dict); i++ {
		if target < dict[i] {
			break
		}
		subset = append(subset, dict[i])
		dfsCombinationSum(dict, target-dict[i], i, subset, subsets)
		subset = subset[:len(subset)-1]
	}
}

func removeDuplicatesAndSort(nums []int) []int {
	n := len(nums)
	if n == 0 {
		return nums
	}
	sort.Ints(nums)
	slow := 1
	for fast := 1; fast < n; fast++ {
		if nums[fast] != nums[fast-1] {
			nums[slow] = nums[fast]
			slow++
		}
	}
	return nums[:slow]
}
