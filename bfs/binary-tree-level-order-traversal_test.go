package bfs_test

import (
	"testing"

	"github.com/ysmood/got"

	"github.com/go-rythm/n9r/bfs"
)

func TestLevelOrder3(t *testing.T) {
	lchild := &bfs.TreeNode{Val: 2}
	rchild := &bfs.TreeNode{Val: 3}
	root := &bfs.TreeNode{Val: 1, Left: lchild, Right: rchild}
	v := bfs.LevelOrder3(root)
	got.T(t).Eq(v, [][]int{{1}, {2, 3}})
}
