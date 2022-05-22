## 组合类DFS

### 摘要

* 组合类搜索的隐式图模型
* 组合类搜索的时间复杂度和适用条件
* 如何对搜索结果进行去重
    * 选代表 vs 哈希表

### [17](https://www.lintcode.com/problem/subsets) 子集

使用组合类搜索的专用深度优先搜索算法。 一层一层的决策每个数要不要放到最后的集合里。

每一层只有两个选择，放或不放

```go
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
```

更通用的解法

![subset](https://gitee.com/luxcgo/imgs4md/raw/master/img/20220522172431.jpeg)

```go
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
```

### [18](https://www.lintcode.com/problem/subsets-ii) 子集 II

列表中可能具有重复数字

```go
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
```

