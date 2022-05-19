package dnc

// https://www.lintcode.com/problem/474

type ParentTreeNode struct {
	Parent *ParentTreeNode
	Left   *ParentTreeNode
	Right  *ParentTreeNode
}

/*
* @param root: The root of the binary tree.
* @param A: A TreeNode in a Binary.
* @param B: A TreeNode in a Binary.
* @return: Return the least common ancestor(LCA) of the two nodes.
 */
func LowestCommonAncestorII(root, A, B *ParentTreeNode) *ParentTreeNode {
	parentSet := map[*ParentTreeNode]bool{}
	cur := A
	for cur != nil {
		parentSet[cur] = true
		cur = cur.Parent
	}
	cur = B
	for cur != nil {
		if parentSet[cur] {
			return cur
		}
		cur = cur.Parent
	}
	return nil
}
