## 从搜索到动规——记忆化搜索入门

### 摘要

* 什么时候可以使用记忆化搜索
* DFS和记忆化搜索的区别
* 如何用三行代码让DFS变成记忆化搜索
* 用记忆化搜索解决博弈型动态规划
* 记忆化搜索的缺陷

### [109](https://www.lintcode.com/problem/triangle/) 数字三角形

找出数字三角形中从上到下的一条最短路径

**数字三角形 vs 二叉树**

* n 层的数字三角形：O(n^2) 个节点
* n 层的满二叉树：O(2^n) 个节点

#### DFS: Traverse

O(2^n)

```go
var minimum int

func MinimumTotal(triangle [][]int) int {
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
```

#### DFS: Divide & Conquer

O(2^n)

```go
func MinimumTotal(triangle [][]int) int {
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
```

#### DFS: Divide & Conquer + Hash

O(n^2)

```go
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
```

### 记忆化搜索 Memoization Search

> 将函数的计算结果保存下来，下次通过同样的参数访问时，直接返回保存下来的结果就叫做记忆化搜索。

问：

1. 对这个函数有什么限制条件没有?
    * 需要有返回值

2. 和系统设计中的什么很像?
    * cache

记忆化搜索通常能够将指数级别的时间复杂度降低到多项式级别。

### 记忆化搜索与动态规划

* 记忆化搜索的本质：动态规划
    * 动态规划快的原因：避免了重复计算
* 记忆化搜索 = 动态规划(DP)
    * 记忆化搜索是动态规划的一种实现方式
    * 记忆化搜索是用搜索的方式实现了动态规划
    * 因此记忆化搜索，就是动态规划

### [1300](https://www.lintcode.com/problem/bash-game/) 巴什博弈

记忆化搜索非常适合博弈型动态规划

```go
func CanWinBash(n int) bool {
	memo := make(map[int]bool)
	return memoSearch(n, memo)
}

func memoSearch(n int, memo map[int]bool) bool {
	if n <= 3 {
		return true
	}
	if v, ok := memo[n]; ok {
		return v
	}
	for i := 1; i < 4; i++ {
		if !memoSearch(n-i, memo) {
			memo[n] = true
			return true
		}
	}
	memo[n] = false
	return false
}
```

#### 这份代码有什么问题?

**StackOverflow**

* n 可能很大，深度达到 O(n) 级别
* <span style="color:red">**如果时间复杂度和递归深度都是 O(n) 级别会栈溢出**</span>
* 如果时间复杂度是 O(n^2) 级别，递归深度是 O(n) 级别就不会溢出

#### 解决思路

```
通过记忆化搜索的代码得到 n=1,2,3,4,... 的小数据结果
找规律，发现 n % 4 == 0 是 false，其他是 true
进一步想到 n % 4 == 0 时，先手取 x，后手取 4-x 先手必败
反之先手取 n % 4 个石子就让对手面对必败局面
```

```go
func CanWinBash(n int) bool {
	return n%4 != 0
}
```

### 记忆化搜索的缺点

不适合解决 O(n) 时间复杂度的 DP 问题<br/>
因为会有 StackOverflow 的风险
