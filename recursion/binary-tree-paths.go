package recursion

import (
	"bytes"
	"encoding/json"
	"strconv"
	"strings"
)

// https://www.lintcode.com/problem/binary-tree-paths

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
 * @param root: the root of the binary tree
 * @return: all root-to-leaf paths
 *          we will sort your return value in output
 */
func BinaryTreePaths(root *TreeNode) slice {
	path := []*TreeNode{root}
	paths := []string{strconv.Itoa(root.Val)}
	findPaths(root, &path, &paths)
	return slice(paths)
}

// 遍历法
func findPaths(node *TreeNode, path *[]*TreeNode, paths *[]string) {
	if node == nil {
		return
	}
	if node.Left == nil && node.Right == nil {
		pathVal := make([]string, 0, len(*path))
		for _, v := range *path {
			pathVal = append(pathVal, strconv.Itoa(v.Val))
		}
		*paths = append(*paths, strings.Join(pathVal, "->"))
	}

	*path = append(*path, node.Left)
	findPaths(node.Left, path, paths)
	pathSlice := *path
	*path = pathSlice[:len(*path)-1]

	*path = append(*path, node.Right)
	findPaths(node.Right, path, paths)
	pathSlice = *path
	*path = pathSlice[:len(*path)-1]
}

type slice []string

func (this slice) Marshal() ([]byte, error) {
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(this)
	if err != nil {
		return nil, err
	}
	return buffer.Bytes(), nil
}

func findNodes(node *TreeNode, nodes *[]*TreeNode) {
	if node == nil {
		return
	}
	*nodes = append(*nodes, node)
	findNodes(node.Left, nodes)
	findNodes(node.Right, nodes)
}

// 分治法
func binaryTreePaths(node *TreeNode) []string {
	paths := []string{}
	if node == nil {
		return paths
	}
	if node.Left == nil && node.Right == nil {
		paths = append(paths, ""+strconv.Itoa(node.Val))
		return paths
	}

	for _, leftPath := range binaryTreePaths(node.Left) {
		paths = append(paths, strconv.Itoa(node.Val)+"->"+leftPath)
	}
	for _, rightPath := range binaryTreePaths(node.Right) {
		paths = append(paths, strconv.Itoa(node.Val)+"->"+rightPath)
	}

	return paths
}
