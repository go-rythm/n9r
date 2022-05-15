## 宽度优先搜索与图论入门

### 摘要

* 宽度优先搜索的适用场景
* 宽度优先搜索的三种实现方法
    1. 两个队列的实现方法
    2. DummyNode的实现方法
    3. 一个队列的实现方法
* 无向图和有向图的存储方法

### 宽度优先搜索的适用场景

* 分层遍历
    * 一层一层的遍历一个图、树、矩阵
    * 简单图最短路径
        * 简单图的定义是，图中所有的边长都一样
* 连通块问题
         * 通过图中一个点找到其他所有连通的点
         * 找到所有方案问题的一种非递归实现方式
         

* 拓扑排序
    * 实现容易度远超过 DFS

### 宽度优先搜索的三种实现方法

以一个题目为例，[69 · 二叉树的层次遍历](https://www.lintcode.com/problem/69/)

1. 两个队列的实现方法

    ```go
    func LevelOrder(root *TreeNode) [][]int {
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
    ```

    使用两个队列来回交替

    第一个队列，先保存根节点，即为第一层

    第二个队列，遍历第一队列并保存所有节点的子节点，即为第二层

    第二个队列赋值给第一个队列，循环往复

2. DummyNode的实现方法

    ```go
    func LevelOrder(root *TreeNode) [][]int {
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
    ```

    使用一个`nil`指针标识一层的结束

    优点是只有一层循环

3. 一个队列的实现方法

    ```go
    func LevelOrder(root *TreeNode) [][]int {
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
    ```

    和两个队列差距不大，只是把`node`都保存在一个`queue`中，并适时截断

