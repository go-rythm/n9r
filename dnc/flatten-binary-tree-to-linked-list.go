package dnc

// http://www.lintcode.com/problem/flatten-binary-tree-to-linked-list/

/**
 * @param root: a TreeNode, the root of the binary tree
 * @return: nothing
 */
func Flatten(root *TreeNode) {
	flattenAndReturnLastNode(root)
}

func flattenAndReturnLastNode(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	leftLast := flattenAndReturnLastNode(root.Left)
	rightLast := flattenAndReturnLastNode(root.Right)

	if leftLast != nil {
		leftLast.Right = root.Right
		root.Right = root.Left
		root.Left = nil
	}

	if rightLast != nil {
		return rightLast
	}
	if leftLast != nil {
		return leftLast
	}
	return root
}
