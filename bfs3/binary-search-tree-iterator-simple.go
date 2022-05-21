package bfs3

// https://www.lintcode.com/problem/binary-search-tree-iterator/

type BSTIteratorSimple struct {
	stack []*TreeNode
}

func NewBSTIteratorSimple(root *TreeNode) *BSTIteratorSimple {
	iter := new(BSTIteratorSimple)
	iter.findMostLeft(root)
	return iter
}

func (iter *BSTIteratorSimple) findMostLeft(node *TreeNode) {
	for node != nil {
		iter.stack = append(iter.stack, node)
		node = node.Left
	}
}

func (iter *BSTIteratorSimple) HasNext() bool {
	return len(iter.stack) > 0
}

func (iter *BSTIteratorSimple) Next() *TreeNode {
	node := iter.stack[len(iter.stack)-1]
	iter.stack = iter.stack[:len(iter.stack)-1]
	if node.Right != nil {
		iter.findMostLeft(node.Right)
	}
	return node
}
