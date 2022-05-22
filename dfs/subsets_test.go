package dfs_test

import (
	"testing"

	"github.com/go-rythm/n9r/dfs"
)

// https://www.lintcode.com/problem/subsets

/**
 * @param nums: A set of numbers
 * @return: A list of lists
 *          we will sort your return value in output
 */
func TestSubsets(t *testing.T) {
	slice := []int{9, 0, 3, 5, 7}
	res := dfs.Subsets2(slice)
	t.Log(res)
}

func TestSubsets2(t *testing.T) {
	slice := []int{1, 2}
	res := dfs.Subsets2(slice)
	t.Log(res)
}
