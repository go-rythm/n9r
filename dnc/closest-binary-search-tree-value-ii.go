package dnc

// https://www.lintcode.com/problem/closest-binary-search-tree-value-ii/

/**
 * @param root: the given BST
 * @param target: the given target
 * @param k: the given k
 * @return: k values in the BST that are closest to the target
 *          we will sort your return value in output
 */
func ClosestKValues(root *TreeNode, target float64, k int) []int {
	if root == nil || k == 0 {
		return nil
	}

	s := &sol{
		lowerStack: getStack(root, target),
		upperStack: getStack(root, target),
	}
	if target > float64(s.lowerStack[len(s.lowerStack)-1].Val) {
		s.moveUpper()
	} else {
		s.moveLower()
	}

	var res []int
	for i := 0; i < k; i++ {
		if s.isLowerCloser(target) {
			res = append(res, s.lowerStack[len(s.lowerStack)-1].Val)
			s.moveLower()
		} else {
			res = append(res, s.upperStack[len(s.upperStack)-1].Val)
			s.moveUpper()
		}
	}

	return res
}

func getStack(root *TreeNode, target float64) []*TreeNode {
	var stack []*TreeNode
	for root != nil {
		stack = append(stack, root)
		if target < float64(root.Val) {
			root = root.Left
		} else {
			root = root.Right
		}
	}
	return stack
}

type sol struct {
	lowerStack []*TreeNode
	upperStack []*TreeNode
}

func (s *sol) moveUpper() {
	node := s.upperStack[len(s.upperStack)-1]
	if node.Right != nil {
		node = node.Right
		for node != nil {
			s.upperStack = append(s.upperStack, node)
			node = node.Left
		}
	} else {
		s.upperStack = s.upperStack[:len(s.upperStack)-1]
		for len(s.upperStack) > 0 && s.upperStack[len(s.upperStack)-1].Right == node {
			node = s.upperStack[len(s.upperStack)-1]
			s.upperStack = s.upperStack[:len(s.upperStack)-1]
		}
	}
}

func (s *sol) moveLower() {
	node := s.lowerStack[len(s.lowerStack)-1]
	if node.Left != nil {
		node = node.Left
		for node != nil {
			s.lowerStack = append(s.lowerStack, node)
			node = node.Right
		}
	} else {
		s.lowerStack = s.lowerStack[:len(s.lowerStack)-1]
		for len(s.lowerStack) > 0 && s.lowerStack[len(s.lowerStack)-1].Left == node {
			node = s.lowerStack[len(s.lowerStack)-1]
			s.lowerStack = s.lowerStack[:len(s.lowerStack)-1]
		}
	}
}

func (s *sol) isLowerCloser(target float64) bool {
	if len(s.lowerStack) == 0 {
		return false
	}
	if len(s.upperStack) == 0 {
		return true
	}
	return target-float64(s.lowerStack[len(s.lowerStack)-1].Val) < float64(s.upperStack[len(s.upperStack)-1].Val)-target
}
