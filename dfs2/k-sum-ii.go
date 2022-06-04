package dfs2

// https://www.lintcode.com/problem/k-sum-ii/

/**
 * @param a: an integer array
 * @param k: a postive integer <= length(A)
 * @param target: an integer
 * @return: A list of lists of integer
 *          we will sort your return value in output
 */
func KSumII(a []int, k int, target int) [][]int {
	// 这里需要sort(a)吗?不需要，本题无需按照字母序，也无重复
	// 排序所有字母，排序的意义:
	// 1.可以按照字母序得到结果
	// 2.相同的字母在一起，方便去重
	subsets := [][]int{}
	dfsKSumII(a, k, target, 0, nil, &subsets)
	return subsets
}

func dfsKSumII(a []int, k int, target int, idx int, subset []int, subsets *[][]int) {
	if k == 0 && target == 0 {
		// log.Println(subset)
		*subsets = append(*subsets, append([]int{}, subset...))
		return
	}

	if k == 0 || target <= 0 {
		return
	}

	for i := idx; i < len(a); i++ {
		// log.Printf("i = %d, a[i] = %d, k = %d, target = %d, idx = %d", i, a[i], k, target, idx)
		subset = append(subset, a[i])
		dfsKSumII(a, k-1, target-a[i], i+1, subset, subsets)
		subset = subset[:len(subset)-1]
	}
}
