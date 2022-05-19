package dnc

// https://www.lintcode.com/problem/578/

/*
* @param root: The root of the binary tree.
* @param A: A TreeNode
* @param B: A TreeNode
* @return: Return the LCA of the two nodes.
 */
func LowestCommonAncestor3(root, A, B *TreeNode) *TreeNode {
	a, b, lca := helper3(root, A, B)
	if a && b {
		return lca
	}
	return nil
}

func helper3(root, A, B *TreeNode) (bool, bool, *TreeNode) {
	if root == nil {
		return false, false, nil
	}

	leftA, leftB, leftNode := helper3(root.Left, A, B)
	rightA, rightB, rightNode := helper3(root.Right, A, B)

	a := leftA || rightA || root == A
	b := leftB || rightB || root == B

	if root == A || root == B {
		return a, b, root
	}

	if leftNode != nil && rightNode != nil {
		return a, b, root
	}
	if leftNode != nil {
		return a, b, leftNode
	}
	if rightNode != nil {
		return a, b, rightNode
	}
	return a, b, nil
}
