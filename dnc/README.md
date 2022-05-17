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

* 碰到二叉树的问题，就想想整棵树在该问题上的结果
* 和左右儿子在该问题上的结果之间的联系是什么

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

