package dnc

import (
	"testing"
)

func TestKthSmallest(t *testing.T) {
	root := &TreeNode{Val: 1}
	root.Right = &TreeNode{Val: 2}

	want := 2
	got := KthSmallest(root, 2)
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
