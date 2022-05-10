package bfs

// https://www.lintcode.com/problem/69/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func LevelOrder1(root *TreeNode) [][]int {
	res := make([][]int, 0)
	if root == nil {
		return res
	}

	q := []*TreeNode{root}
	for len(q) > 0 {
		var level []int
		len := len(q)
		for i := 0; i < len; i++ {
			node := q[i]
			level = append(level, node.Val)
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}
		res = append(res, level)
		q = q[len:]
	}
	return res
}

func LevelOrder2(root *TreeNode) [][]int {
	res := make([][]int, 0)
	if root == nil {
		return res
	}

	q := []*TreeNode{root}
	for len(q) > 0 {
		var level []int
		var nq []*TreeNode
		for _, node := range q {
			level = append(level, node.Val)
			if node.Left != nil {
				nq = append(nq, node.Left)
			}
			if node.Right != nil {
				nq = append(nq, node.Right)
			}
		}
		res = append(res, level)
		q = nq
	}

	return res
}

func LevelOrder3(root *TreeNode) [][]int {
	res := make([][]int, 0)
	if root == nil {
		return res
	}

	q := []*TreeNode{root, nil}
	var level []int
	for len(q) > 0 {
		node := q[0]
		q = q[1:]
		if node == nil {
			res = append(res, level)
			level = nil
			if len(q) > 0 {
				q = append(q, nil)
			}
			continue
		}
		level = append(level, node.Val)
		if node.Left != nil {
			q = append(q, node.Left)
		}
		if node.Right != nil {
			q = append(q, node.Right)
		}
	}

	return res
}
