## 解决99%二叉树问题的算法——分治法

### 摘要

* 用分治法解决二叉树求值求路径的问题
    * 理解什么是搜索中的回溯
  
* 用分治法解决二叉树形态变换的问题
    * 全局变量在代码中的危害

### 分治法 Divide & Conquer

将大规模问题拆分为若干个小规模的同类型问题去处理的算法

分治法和二分法(Binary Search)有什么区别?

* 二分法会每次丢弃掉一半
* 分治法分完以后两边都要处理

### 什么样的数据结构适合分治法?

* 数组：一个大数组可以拆分为若干个不相交的子数组<br/>
归并排序，快速排序，都是基于数组的分治法
* 二叉树：整棵树的左子树和右子树都是二叉树<br/>
二叉树的大部分题都可以使用分治法解决

### 独孤九剑 —— 破枪式

**碰到二叉树的问题，就想想整棵树在该问题上的结果和左右儿子在该问题上的结果之间的联系是什么**

### 二叉树考点剖析

* 高度：最坏 O(n) 最好 O(logn) 一般用 O(h) 来表示更合适

1. 第一类：Maximum / Minimum / Average / Sum / Paths
    * 考察形态：二叉树上求值，求路径 
    * 代表例题：http://www.lintcode.com/problem/subtree-with-maximum-average/ 
    * 考点本质：深度优先搜索(Depth First Search)
2. 第二类
    * 考察形态：二叉树结构变化 
    * 代表例题：http://www.lintcode.com/problem/invert-binary-tree/ 
    * 考点本质：深度优先搜索(Depth First Search)
3. 第三类
    * 考察形态：二叉查找树(Binary Search Tree) 
    * 代表例题：http://www.lintcode.com/problem/validate-binary-search-tree/ 
    * 考点本质：深度优先搜索(Depth First Search)

总结：Tree-based Depth First Search

* 不管二叉树的题型如何变化，考点都是基于树的深度优先搜索

### 一张图搞明白:递归，DFS，回溯，遍历，分治，迭代

![5F95BDD5-4B52-4F9E-8EFA-1668E43EE7B3](https://gitee.com/luxcgo/imgs4md/raw/master/img/20220517230556.jpeg)

将递归和非递归理解为算法的**一种实现方式**而不是算法

### [596](http://www.lintcode.com/problem/minimum-subtree/) 最小子树（第一类）

一棵有n个节点的二叉树有多少棵子树？  
n棵，每个节点都可以作为子树的根节点

#### 使用了全局变量的分治法

**全局变量的坏处**

* 函数不“纯粹”，容易出 Bug 
* 不利于多线程化，对共享变量加锁带来效率下降

```go
var (
	minSum  int
	minRoot *TreeNode
)

/**
 * @param root: the root of binary tree
 * @return: the root of the minimum subtree
 */
func FindSubtree(root *TreeNode) *TreeNode {
	minSum = math.MaxInt
	minRoot = nil
	getTreeSum(root)
	return minRoot
}

func getTreeSum(root *TreeNode) int {
	if root == nil {
		return 0
	}

	leftWeight := getTreeSum(root.Left)
	rightWeight := getTreeSum(root.Right)
	rootWeight := leftWeight + rightWeight + root.Val

	if rootWeight < minSum {
		minSum = rootWeight
		minRoot = root
	}
	return rootWeight
}
```

#### 纯分治的方法

```go
func FindSubtree2(root *TreeNode) *TreeNode {
	_, subtree, _ := helper(root)
	return subtree
}

func helper(root *TreeNode) (int, *TreeNode, int) {
	if root == nil {
		return math.MaxInt, nil, 0
	}
	leftMin, leftSubtree, leftWeight := helper(root.Left)
	rightMin, rightSubtree, rightWeight := helper(root.Right)

	rootWeight := leftWeight + rightWeight + root.Val
	if leftMin < rightMin && leftMin < rootWeight {
		return leftMin, leftSubtree, rootWeight
	}
	if rightMin < leftMin && rightMin < rootWeight {
		return rightMin, rightSubtree, rootWeight
	}
	return rootWeight, root, rootWeight
}
```

### [474](https://www.lintcode.com/problem/474) 最近公共祖先 II（第一类）

问法1：如果有父指针

使用 HashSet 记录从 A 到根的所有点，访问从 B 到根的所有点，第一个出现在 HashSet 中的就是

```go
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
```

### [88](https://www.lintcode.com/problem/88) 最近公共祖先（第一类）

问法2：两个节点都在树里

给你 root, A, B 三个点的信息，A和B保证都在 root 的下面

定义返回值:

* A,B 都存在 -> return LCA(A,B) 
* 只有A --> return A 
* 只有B --> return B
* A,B 都不存在 --> return nil

> 递归时间复杂度：递归一次 * 次数
>
> 递归空间复杂度：递归一次 + 递归深度

```go
func LowestCommonAncestor(root, A, B *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	// 如果root为A或B，立即返回，无需继续向下寻找
	if root == A || root == B {
		return root
	}
	// 分别去左右子树寻找A和B
	left := LowestCommonAncestor(root.Left, A, B)
	right := LowestCommonAncestor(root.Right, A, B)

	// 如果A、B分别存于两棵子树，root为LCA，返回root
	if left != nil && right != nil {
		return root
	}
	// 左子树有一个点或者左子树有LCA
	if left != nil {
		return left
	}
	// 右子树有一个点或者右子树有LCA
	if right != nil {
		return right
	}
	// 左右子树啥都没有
	return nil
}
```

### [578](http://www.lintcode.com/problem/lowest-common-ancestor-iii/) 最近公共祖先 III（第一类）

问法3：两个节点不一定都在树里

root, p, q，但是不保证 root 里一定有 p 和 q

```go
func LowestCommonAncestor3(root, A, B *TreeNode) *TreeNode {
	a, b, lca := helper3(root, A, B)
	if a && b {
		return lca
	}
	return nil
}

func helper3(root, A, B *TreeNode) (bool, bool, *TreeNode) {
	if root == nil {
		return false, false, nil
	}

	leftA, leftB, leftNode := helper3(root.Left, A, B)
	rightA, rightB, rightNode := helper3(root.Right, A, B)

	a := leftA || rightA || root == A
	b := leftB || rightB || root == B

	if root == A || root == B {
		return a, b, root
	}

	if leftNode != nil && rightNode != nil {
		return a, b, root
	}
	if leftNode != nil {
		return a, b, leftNode
	}
	if rightNode != nil {
		return a, b, rightNode
	}
	return a, b, nil
}
```

### [453](http://www.lintcode.com/problem/flatten-binary-tree-to-linked-list/) 将二叉树拆成链表（第二类）

将二叉树拆成链表。进行前序遍历，将上一个节点的右指针指向当前节点。

![flatten](https://gitee.com/luxcgo/imgs4md/raw/master/img/20220519225139.png)

```go
func Flatten(root *TreeNode) {
	flattenAndReturnLastNode(root)
}

func flattenAndReturnLastNode(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	leftLast := flattenAndReturnLastNode(root.Left)
	rightLast := flattenAndReturnLastNode(root.Right)

	if leftLast != nil {
		leftLast.Right = root.Right
		root.Right = root.Left
		root.Left = nil
	}

	if rightLast != nil {
		return rightLast
	}
	if leftLast != nil {
		return leftLast
	}
	return root
}
```

### 二叉搜索树 | Binary Search Tree

**BST 基本性质**

* 从定义出发：
    * 左子树都比根节点小
    * 右子树都不小于根节点

* 从效果出发
    * 中序遍历in-order traversal是“不下降”序列

* 性质：
    * 如果一棵二叉树的中序遍历不是“不下降”序列，则一定不是BST
    
    * 如果一棵二叉树的中序遍历是不下降，也未必是BST
        * 比如下面这棵树就不是 BST，但是它的中序遍历是不下降序列。
    
        ```sh
            1
          /   \
         1     1
        ```
    
* 高度：
  
    * 同样是最坏 O(n) 最好 O(logn)，用 O(h) 表示更合适
    * 只有 Balanced Binary Tree (平衡二叉树) 才是 O(logn)


### 红黑树 | Red-black Tree

红黑树是一种 Balanced BST

* O(LogN) 的时间内实现增删查改
* O(LogN) 的时间内实现找最大找最小
* O(LogN) 的时间内实现找比某个数小的最大值(upperBound)和比某个数大的最小值(lowerBound)

### [902](https://www.lintcode.com/problem/kth-smallest-element-in-a-bst/) BST中第K小的元素

时间复杂度分析：**O(k + h)**

当 k 是 1 的时候 => O(h)

当 k 是 n 的时候 => O(n) 

k和h两者取大值

```go
func KthSmallest(root *TreeNode, k int) int {
	var stack []*TreeNode

	for root != nil {
		stack = append(stack, root)
		root = root.Left
	}

	for i := 0; i < k-1; i++ {
		node := stack[len(stack)-1]

		if node.Right == nil {
			node = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			for len(stack) != 0 && stack[len(stack)-1].Right == node {
				node = stack[len(stack)-1]
				stack = stack[:len(stack)-1]
			}
		} else {
			node = node.Right
			for node != nil {
				stack = append(stack, node)
				node = node.Left
			}
		}
	}
	return stack[len(stack)-1].Val
}
```

