package dnc

// https://www.lintcode.com/problem/kth-smallest-element-in-a-bst/

/**
 * @param root: the given BST
 * @param k: the given k
 * @return: the kth smallest element in BST
 */
func KthSmallest(root *TreeNode, k int) int {
	var stack []*TreeNode

	for root != nil {
		stack = append(stack, root)
		root = root.Left
	}

	for i := 0; i < k-1; i++ {
		node := stack[len(stack)-1]

		if node.Right == nil {
			node = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			for len(stack) != 0 && stack[len(stack)-1].Right == node {
				node = stack[len(stack)-1]
				stack = stack[:len(stack)-1]
			}
		} else {
			node = node.Right
			for node != nil {
				stack = append(stack, node)
				node = node.Left
			}
		}
	}
	return stack[len(stack)-1].Val
}
