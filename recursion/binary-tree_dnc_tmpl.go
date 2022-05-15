package recursion

type result struct{}

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

func merge(l, r *result) *result {
	return &result{}
}
