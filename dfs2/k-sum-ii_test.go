package dfs2_test

import (
	"testing"

	"github.com/go-rythm/n9r/dfs2"
)

func TestKSumII(t *testing.T) {
	res := dfs2.KSumII([]int{1, 3, 2, 4}, 2, 5)
	t.Log(res)
}
