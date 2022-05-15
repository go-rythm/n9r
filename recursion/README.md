## 用递归实现遍历法和分治法

### 递归、深搜和回溯法的区别

Recursion / DFS / Backtracking

### 递归 | Recursion

* 递归函数：程序的一种实现方式，即函数进行了自我调用
* 递归算法：即大问题的结果依赖于小问题的结果，于是先用递归函数求解小问题
* 一般我们说递归的时候，大部分时候都在说递归函数而不是递归算法

### 深度优先搜索 | Depth First Search

* 可以使用递归函数来实现
* 也可以不用递归函数来实现，如自己通过一个手动创建的栈 Stack 进行操作
* 深度优先搜索通常是指在搜索的过程中优先搜索深度更深的点而不是按照宽度搜索同层节点

### 回溯 | Backtracking

* 回溯法：就是深度优先搜索算法
* 回溯操作：递归函数在回到上一层递归调用处的时候，一些参数需要改回到调用前的值，这个操作就是回溯，即让状态参数回到之前的值，递归调用前做了什么改动，递归调用之后都改回来

### 找点 vs 找路径

[480](https://www.lintcode.com/problem/480/) 二叉树的所有路径

是否需要手动“回溯”的判断标准

```go
func findNodes(node *TreeNode, nodes *[]*TreeNode) {
	if node == nil {
		return
	}
	*nodes = append(*nodes, node)
	findNodes(node.Left, nodes)
	findNodes(node.Right, nodes)
}
```

```go
func findPaths(node *TreeNode, path *[]*TreeNode, paths *[]string) {
	if node == nil {
		return
	}
	if node.Left == nil && node.Right == nil {
		pathVal := make([]string, 0, len(*path))
		for _, v := range *path {
			pathVal = append(pathVal, strconv.Itoa(v.Val))
		}
		*paths = append(*paths, strings.Join(pathVal, "->"))
	}

	*path = append(*path, node.Left)
	findPaths(node.Left, path, paths)
	pathSlice := *path
	*path = pathSlice[:len(*path)-1]

	*path = append(*path, node.Right)
	findPaths(node.Right, path, paths)
	pathSlice = *path
	*path = pathSlice[:len(*path)-1]
}
```

