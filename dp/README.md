## 动态规划入门与动规四要素

### 摘要

* 动态规划DP与记忆化搜索的关系
* 多重循环与记忆化搜索实现动态规划的区别
* 什么时候使用动态规划
    * 动态规划四要素是什么
* 自底向上的动态规划
* 自顶向下的动态规划

### 动态规划 Dynamic Programming

> 简称动规或者`DP`，是一种算法思想，而不是一种具体的算法。

#### 核心思想：由大化小

动态规划的算法思想：大规模问题的依赖于小规模问题的计算结果<br/>
类似思想算法的还有：递归，分治法

#### 实现方法

1. 记忆化搜索 (使用递归实现)
2. 多重循环 (使用for循环实现)

### 自顶向下 vs 自底向上

|        | 自顶向下的动态规划 | 自底向上的动态规划 |
| ------ | ------------------ | ------------------ |
| 状态   | 坐标               | 坐标               |
| 方程   | 从哪儿来           | 到哪儿去           |
| 初始化 | 起点               | 终点               |
| 答案   | 终点               | 起点               |

两种方法都可以，爱用哪个用哪个<br/>
一个关心从哪儿来，一个关心到哪儿去

### [109](https://www.lintcode.com/problem/triangle/) 数字三角形

让我们用多重循环的重做一下这个题

**自顶向下**

```go
func MinimumTotal(triangle [][]int) int {
	n := len(triangle)
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, i+1)
	}

	dp[0][0] = triangle[0][0]
	for i := 1; i < n; i++ {
		dp[i][0] = dp[i-1][0] + triangle[i][0]
		dp[i][i] = dp[i-1][i-1] + triangle[i][i]
	}

	for i := 2; i < n; i++ {
		for j := 1; j < i; j++ {
			dp[i][j] = min(dp[i-1][j], dp[i-1][j-1]) + triangle[i][j]
		}
	}

	var res = dp[n-1][0]
	for i := 1; i < n; i++ {
		res = min(res, dp[n-1][i])
	}
	return res
}
```

### 动规四要素

状态，方程，初始化，答案

#### 递归四要素 vs 动规四要素

* 动规的状态 State —— 递归的定义
    * 用 `f[i]` 或者 `f[i][j]` 代表在某些特定条件下某个规模更小的问题的答案
    * 规模更小用参数 i,j 之类的来划定
* 动规的方程 Function —— 递归的拆解
    * 大问题如何拆解为小问题
    * `f[i][j]` = 通过规模更小的一些状态求 max / min / sum / or 来进行推导
* 动规的初始化 Initialize —— 递归的出口
    * 设定无法再拆解的极限小的状态下的值
    * 如 `f[i][0]` 或者 `f[0][i]`
* 动规的答案 Answer —— 递归的调用
    * 最后要求的答案是什么
    * 如 `f[n][m]` 或者 `max(f[n][0], f[n][1] ... f[n][m])`

> 递归四要素完全对应动规四要素
>
> 这也就是为什么动态规划可以使用 “递归”版本的记忆化搜索来解决的原因!

### [114 ](https://www.lintcode.com/problem/unique-paths/)不同的路径

求方案总数类 DP 题

```go
func UniquePaths(m int, n int) int {
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	for i := 0; i < n; i++ {
		dp[0][i] = 1
	}
	for i := 0; i < m; i++ {
		dp[i][0] = 1
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}
	return dp[m-1][n-1]
}
```

