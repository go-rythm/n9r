package dnc

import "math"

// http://www.lintcode.com/problem/minimum-subtree/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var (
	minSum  int
	minRoot *TreeNode
)

/**
 * @param root: the root of binary tree
 * @return: the root of the minimum subtree
 */
func FindSubtree(root *TreeNode) *TreeNode {
	minSum = math.MaxInt
	minRoot = nil
	getTreeSum(root)
	return minRoot
}

func getTreeSum(root *TreeNode) int {
	if root == nil {
		return 0
	}

	leftWeight := getTreeSum(root.Left)
	rightWeight := getTreeSum(root.Right)
	rootWeight := leftWeight + rightWeight + root.Val

	if rootWeight < minSum {
		minSum = rootWeight
		minRoot = root
	}
	return rootWeight
}

func FindSubtree2(root *TreeNode) *TreeNode {
	_, subtree, _ := helper(root)
	return subtree
}

func helper(root *TreeNode) (int, *TreeNode, int) {
	if root == nil {
		return math.MaxInt, nil, 0
	}
	leftMin, leftSubtree, leftWeight := helper(root.Left)
	rightMin, rightSubtree, rightWeight := helper(root.Right)

	rootWeight := leftWeight + rightWeight + root.Val
	if leftMin < rightMin && leftMin < rootWeight {
		return leftMin, leftSubtree, rootWeight
	}
	if rightMin < leftMin && rightMin < rootWeight {
		return rightMin, rightSubtree, rootWeight
	}
	return rootWeight, root, rootWeight
}
