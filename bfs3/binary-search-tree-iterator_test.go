package bfs3

import (
	"fmt"
	"testing"
)

// https://www.lintcode.com/problem/binary-search-tree-iterator/

func TestBSTIterator(t *testing.T) {
	root := &TreeNode{
		Val: 10,
		Left: &TreeNode{
			Val:   1,
			Right: &TreeNode{Val: 6},
		},
		Right: &TreeNode{
			Val:   11,
			Right: &TreeNode{Val: 12},
		},
	}
	iter := NewBSTIterator(root)

	var vals []int
	for iter.HasNext() {
		vals = append(vals, iter.Next().Val)
	}

	got := fmt.Sprint(vals)
	want := fmt.Sprint([]int{1, 6, 10, 11, 12})
	// t.Log(vals)
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
