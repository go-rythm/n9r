package bfs3

// https://www.lintcode.com/problem/binary-search-tree-iterator/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type BSTIterator struct {
	stack []*TreeNode
}

func NewBSTIterator(root *TreeNode) *BSTIterator {
	iter := new(BSTIterator)
	for root != nil {
		iter.stack = append(iter.stack, root)
		root = root.Left
	}
	return iter
}

func (iter *BSTIterator) HasNext() bool {
	return len(iter.stack) > 0
}

func (iter *BSTIterator) Next() *TreeNode {
	cur := iter.stack[len(iter.stack)-1]
	node := cur
	if node.Right != nil {
		node = node.Right
		for node != nil {
			iter.stack = append(iter.stack, node)
			node = node.Left
		}
	} else {
		iter.stack = iter.stack[:len(iter.stack)-1]
		for iter.HasNext() && iter.stack[len(iter.stack)-1].Right == node {
			node = iter.stack[len(iter.stack)-1]
			iter.stack = iter.stack[:len(iter.stack)-1]
		}
	}
	return cur
}
