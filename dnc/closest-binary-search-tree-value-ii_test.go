package dnc

import "testing"

func TestClosestKValues(t *testing.T) {
	root := &TreeNode{Val: 2}
	root.Left = &TreeNode{Val: 1}
	root.Right = &TreeNode{Val: 3}

	got := ClosestKValues(root, 5.571429, 2)
	t.Log(got)
}
