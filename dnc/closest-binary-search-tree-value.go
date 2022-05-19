package dnc

import "math"

// https://www.lintcode.com/problem/900/

/**
 * @param root: the given BST
 * @param target: the given target
 * @return: the value in the BST that is closest to the target
 */
func ClosestValueRecursion(root *TreeNode, target float64) int {
	if root == nil {
		return 0
	}
	lowerNode := lowerBound(root, target)
	upperNode := upperBound(root, target)
	if lowerNode == nil {
		return upperNode.Val
	}
	if upperNode == nil {
		return lowerNode.Val
	}
	if target-float64(lowerNode.Val) > float64(upperNode.Val)-target {
		return upperNode.Val
	}
	return lowerNode.Val
}

func lowerBound(root *TreeNode, target float64) *TreeNode {
	if root == nil {
		return nil
	}

	if target < float64(root.Val) {
		return lowerBound(root.Left, target)
	}

	lowerNode := lowerBound(root.Right, target)
	if lowerNode != nil {
		return lowerNode
	}
	return root
}

func upperBound(root *TreeNode, target float64) *TreeNode {
	if root == nil {
		return nil
	}

	if target >= float64(root.Val) {
		return upperBound(root.Right, target)
	}

	upperNode := upperBound(root.Left, target)
	if upperNode != nil {
		return upperNode
	}
	return root
}

func ClosestValue(root *TreeNode, target float64) int {
	upper := root
	lower := root
	for root != nil {
		if target < float64(root.Val) {
			upper = root
			root = root.Left
		} else if target > float64(root.Val) {
			lower = root
			root = root.Right
		} else {
			return root.Val
		}
	}
	if math.Abs(float64(upper.Val)-target) < math.Abs(float64(lower.Val)-target) {
		return upper.Val
	} else {
		return lower.Val
	}
}
