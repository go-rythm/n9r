package recursion

// https://www.lintcode.com/problem/balanced-binary-tree/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

/**
 * @param root: The root of binary tree.
 * @return: True if this Binary tree is Balanced, or false.
 */
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

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func abs(a, b int) int {
	if a-b > 0 {
		return a - b
	} else {
		return -a + b
	}
}
