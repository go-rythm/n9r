## 使用非递归实现二叉树的遍历

### 二叉树三种遍历

* 先序遍历 Pre-order
* 中序遍历 In-order
* 后序遍历 Post-order(分治法)

### [86](https://www.lintcode.com/problem/binary-search-tree-iterator/) 二叉查找树迭代器

通过实现 hasNext 和 next 两个方法，从而实现二叉查找树的中序遍历迭代器

**实现要点**

* 递归 → 非递归，意味着自己需要控制原来由操作系统控制的**栈**的进进出出
* 如何找到最小的第一个点？最左边的点即是
* 如何求出一个二叉树节点在中序遍历中的下一个节点？
* 在 stack 中记录从根节点到当前节点的整条路径
* 下一个点 = 右子树最小点 or 路径中最近一个通过左子树包含当前点的点

```go
type BSTIterator struct {
	stack []*TreeNode
}

func NewBSTIterator(root *TreeNode) *BSTIterator {
	iter := new(BSTIterator)
	for root != nil {
		iter.stack = append(iter.stack, root)
		root = root.Left
	}
	return iter
}

func (iter *BSTIterator) HasNext() bool {
	return len(iter.stack) > 0
}

func (iter *BSTIterator) Next() *TreeNode {
	cur := iter.stack[len(iter.stack)-1]
	node := cur
	if node.Right != nil {
		node = node.Right
		for node != nil {
			iter.stack = append(iter.stack, node)
			node = node.Left
		}
	} else {
		iter.stack = iter.stack[:len(iter.stack)-1]
		for iter.HasNext() && iter.stack[len(iter.stack)-1].Right == node {
			node = iter.stack[len(iter.stack)-1]
			iter.stack = iter.stack[:len(iter.stack)-1]
		}
	}
	return cur
}
```

### 一种更简单的实现方式

在 stack 中不保存那些已经被 iterator 访问过的节点<br/>
即如果 iterate 到了这个节点，即便右子树还未完全遍历<br/>
也从 stack 里踢出

```go
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
```

