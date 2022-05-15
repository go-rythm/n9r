## 用递归实现遍历法和分治法

### 摘要

- 递归Recursion，深搜DFS，回溯Backtracking的联系和区别
- 递归三要素是什么
- 通过二叉树学习DFS中的遍历法
- 什么是分治法，遍历法和分治法的区别
  - 通过两道实战真题理解分治法和遍历法的区别


### 递归、深搜和回溯法的区别

Recursion / DFS / Backtracking

### 递归 | Recursion

* 递归函数：程序的一种实现方式，即函数进行了自我调用
* 递归算法：即大问题的结果依赖于小问题的结果，于是先用递归函数求解小问题
* 一般我们说递归的时候，大部分时候都在说递归函数而不是递归算法

三要素

1. 定义
2. 出口
3. 拆解

### 深度优先搜索 | Depth First Search

* 可以使用递归函数来实现
* 也可以不用递归函数来实现，如自己通过一个手动创建的栈 Stack 进行操作
* 深度优先搜索通常是指在搜索的过程中优先搜索深度更深的点而不是按照宽度搜索同层节点

### 回溯 | Backtracking

* 回溯法：就是深度优先搜索算法
* 回溯操作：递归函数在回到上一层递归调用处的时候，一些参数需要改回到调用前的值，这个操作就是回溯，即让状态参数回到之前的值，递归调用前做了什么改动，递归调用之后都改回来

### 找点 vs 找路径

**[480](https://www.lintcode.com/problem/480/) 二叉树的所有路径**

是否需要手动“回溯”的判断标准：找路径一般需要手动回溯

```go
// 操作系统帮助回溯
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
// 遍历法
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

    // 手动回溯
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

### 遍历法 vs 分治法

* 遍历法
    * 一个小人拿着一个记事本走遍所有的节点
    * 通常会用到一个全局变量或者是共享参数

* 分治法
    * 分配小弟去做子任务，自己进行结果汇总
    * 通常将利用 return value 记录子问题结果<br/>
        二叉树上的分治法本质上也是在做后序遍历

### [480](https://www.lintcode.com/problem/480/) 二叉树的所有路径

分治法：整棵树路径 = 左子树路径 + 右子树路径

```go
// 分治法
func BinaryTreePaths(node *TreeNode) []string {
	paths := []string{}
	if node == nil {
		return paths
	}
	if node.Left == nil && node.Right == nil {
		paths = append(paths, ""+strconv.Itoa(node.Val))
		return paths
	}

	for _, leftPath := range BinaryTreePaths(node.Left) {
		paths = append(paths, strconv.Itoa(node.Val)+"->"+leftPath)
	}
	for _, rightPath := range BinaryTreePaths(node.Right) {
		paths = append(paths, strconv.Itoa(node.Val)+"->"+rightPath)
	}

	return paths
}
```

### 二叉树上的分治法模板

什么是分治法：

* 整个问题的规模可以用参数去描述
* 问题可以分割成只是参数变化但形式是一致的多个子问题，且子问题之间没有交集

```go
func divideAndConquer(root *TreeNode) *result {
	res := &result{}
	if root == nil {
		// 处理空树返回的结果
		return res
	}

	if root.Left == nil && root.Right == nil {
		// 处理叶子应该返回的结果
		// 如果叶子的返回结果可以通过两个空节点的返回结果得到
		// 就可以省略这一段代码
	}

	// 左子树返回的结果
	leftRes := divideAndConquer(root.Left)
	// 右子树返回的结果
	rightRes := divideAndConquer(root.Right)
	// 整棵树的结果 = 按照一定方法合并左右子树的结果
	rootRes := merge(leftRes, rightRes)

	return rootRes
}
```

### [93](https://www.lintcode.com/problem/balanced-binary-tree/) 平衡二叉树 (判断二叉树是否是平衡的)

```go
func IsBalanced(root *TreeNode) bool {
	isBalanced, _ := divideConquer(root)
	return isBalanced
}

// 定义: 判断 root 为根的二叉树是否是平衡树并且返回高度是多少
func divideConquer(root *TreeNode) (bool, int) {
	// 出口
	if root == nil {
		return true, 0
	}

	// 拆解
	isLeftBalanced, leftHeight := divideConquer(root.Left)
	isRightBalanced, rightHeight := divideConquer(root.Right)
	rootHeight := max(leftHeight, rightHeight) + 1

	if !isLeftBalanced || !isRightBalanced {
		return false, rootHeight
	}
	if abs(leftHeight, rightHeight) > 1 {
		return false, rootHeight
	}
	return true, rootHeight
}
```

