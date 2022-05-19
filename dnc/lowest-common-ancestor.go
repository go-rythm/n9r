package dnc

// https://www.lintcode.com/problem/88

/*
* @param root: The root of the binary tree.
* @param A: A TreeNode in a Binary.
* @param B: A TreeNode in a Binary.
* @return: Return the least common ancestor(LCA) of the two nodes.
 */
func LowestCommonAncestor(root, A, B *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	// 如果root为A或B，立即返回，无需继续向下寻找
	if root == A || root == B {
		return root
	}
	// 分别去左右子树寻找A和B
	left := LowestCommonAncestor(root.Left, A, B)
	right := LowestCommonAncestor(root.Right, A, B)

	// 如果A、B分别存于两棵子树，root为LCA，返回root
	if left != nil && right != nil {
		return root
	}
	// 左子树有一个点或者左子树有LCA
	if left != nil {
		return left
	}
	// 右子树有一个点或者右子树有LCA
	if right != nil {
		return right
	}
	// 左右子树啥都没有
	return nil
}
