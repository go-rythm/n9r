## 性价比之王——宽度优先搜索

### BFS使用场景

* 连通块问题（Connected Component）
    * 通过一个点找到图中连通的所有点
    * 非递归的方式找所有方案

* 分层遍历（Level Order Traversal）
    * 图的层次遍历
    * 简单图最短路径（Simple Graph Shortest Path）

* 拓扑排序（Topological Sorting）
    * 求任意拓扑序
    * 求是否有拓扑序
    * 求字典序最小的拓扑序
    * 求是否唯一拓扑序


### 图的BFS

* 如果图中存在环,则同一个节点可能重复进入队列

    ```
         a
       /   \
      /     \
      b - -  c 
    ```

    第一层节点 a<br />
    第二层节点 b c<br />
    对于a来说，存在路径a-c-b，b也可以是第三层的，这样b就进了两次队列

* BFS中，为什么同一个节点不需要重复进入队列？

    * 对于连通块问题，不可能带来新的节点
    * 对于最短路问题，不可能带来最短的路径

* 解决方法：使用哈希表去重

* 树是没有环的图

### 问最短路径

| 简单图 | 复杂图                                          |
| ------ | ----------------------------------------------- |
| BFS    | Floyd<br />Dijkstra<br />Bellman-ford<br />SPFA |

* 什么是简单图
    * 没有方向（undirected）
    * 没有权重（unweighted）
    * 两点之间最多只有一条边（no multiple edges）
    * 一个点没有一条边直接连着自己（no graph loops，这里的graph loop指的是自己直接指向自己的loop）

### 问最长路径(不能使用BFS)

* 图可以分层：动态规划 Dynamic Programming 
* 图不可以分层：深度优先搜索 DFS 

> 分层的意思是：路径有一定方向性，不能绕圈，第i层的点只能走到第i+1层不能回到i-1层

### 最简洁的 BFS 算法的通用模板

```go
// step 1: 初始化
// 把初始节点放到 queue 里，如果有多个就都放进去
// 并标记初始节点的距离为 0，记录在 distance 的 hashmap 里
// distance 有两个作用，一是判断是否已经访问过，二是记录离起点的距离
queue := []*Node{node}
distance := map[*Node]int64{
	node: 0,
}

// step 2: 不断访问队列
// while 循环 + 每次 pop 队列中的一个点出来
for len(queue) > 0 {
	head := queue[0]
	queue = queue[1:]
	// step 3: 拓展相邻节点
	// pop 出的节点的相邻节点，加入队列并在 distance 中存储距离
	for _, neighbor := range head.Neighbors {
		if _, ok := distance[neighbor]; ok {
			continue
		}
		distance[neighbor] = distance[head] + 1
		queue = append(queue, neighbor)
	}
}
```

N个点，M条边，图上BFS时间复杂度 = O(N + M)，说是O(M)问题也不大，因为M一般都比N大<br/>
M最大是 O(N^2) 的级别(任意两个点之间都有边)， 所以最坏情况可能是 O(N^2)   

### [137](https://www.lintcode.com/problem/137/) Clone Graph 克隆图

代码分析 —— 低耦合的清晰代码(**劝分不劝和**)

将整个算法分解为三个步骤: 

1. 找到所有点
2. 复制所有点 
3. 复制所有边

以上三个步骤：寻点，复制点，复制边交错在一起也能跑，但可读性就差了很多

```go
func CloneGraph(node *UndirectedGraphNode) *UndirectedGraphNode {
	if node == nil {
		return nil
	}
	nodeSet := findNodesByBfs(node)
	mapping := copyNodes(nodeSet)
	copyEdges(nodeSet, mapping)
	return mapping[node]
}

type nodeSet map[*UndirectedGraphNode]bool

func findNodesByBfs(node *UndirectedGraphNode) nodeSet {
	queue := []*UndirectedGraphNode{node}
	visited := nodeSet{
		node: true,
	}
	for len(queue) > 0 {
		curNode := queue[0]
		queue = queue[1:]
		for _, neighbor := range curNode.Neighbors {
			if visited[neighbor] {
				continue
			}
			queue = append(queue, neighbor)
			visited[neighbor] = true
		}
	}
	return visited
}

func copyNodes(nodeSet nodeSet) map[*UndirectedGraphNode]*UndirectedGraphNode {
	mapping := make(map[*UndirectedGraphNode]*UndirectedGraphNode)
	for node := range nodeSet {
		mapping[node] = &UndirectedGraphNode{
			Label: node.Label,
		}
	}
	return mapping
}

func copyEdges(nodeSet nodeSet, mapping map[*UndirectedGraphNode]*UndirectedGraphNode) {
	for node := range nodeSet {
		newNode := mapping[node]
		for _, neighbor := range node.Neighbors {
			newNeighbor := mapping[neighbor]
			newNode.Neighbors = append(newNode.Neighbors, newNeighbor)
		}
	}
}
```

### 80% 的人都可能会写错的 BFS 算法

![bfs-contrast](https://gitee.com/luxcgo/imgs4md/raw/master/img/20220514003700.jpeg)

如图，左右两种写法，唯一的区别是 node 放入 visited 的时机不同<br/>
左边的是正确的，右边的是错误的，例子如下

 ![bfs-wrong-case](https://gitee.com/luxcgo/imgs4md/raw/master/img/20220514004714.png)

可以看到延迟放入 visited 会导致 queue 中元素重复，当节点数量更多的时候，可想而知这种重复的情况会呈指数级的增长<br/>
所以要牢记**访问过一个节点后，就要立即放入已访问的集合中**，这两步不能分开

### BFS 分层 vs 不分层

之前都是不分层的写法，一个 node 跟着一个 node<br/>
下面提供一种分层的写法

```go
func FindNodesByBfsWithLevel(node *UndirectedGraphNode) nodeSet {
	queue := []*UndirectedGraphNode{node}
	visited := nodeSet{
		node: true,
	}
	for len(queue) > 0 {
        // 每一次 range 就是一层
		for range queue {
			curNode := queue[0]
			queue = queue[1:]
			for _, neighbor := range curNode.Neighbors {
				if visited[neighbor] {
					continue
				}
				visited[neighbor] = true
				queue = append(queue, neighbor)
			}
		}
	}
	return visited
}
```

### [120](https://www.lintcode.com/problem/120/) Word Ladder 单词接龙

首先要生成指定单词的变换列表<br/>
从第一个单词开始，他的子节点就是它变换列表中的所有单词，以此类推，生成一棵树，求的就是目标节点是 end 的最短距离

getNextWords 的时间复杂度分析<br/>
两层循环 + 内部构造字符串 = O(26 * L * L) L为每个词的长度

```go
func LadderLength(start string, end string, dict map[string]struct{}) int {
	dict[end] = struct{}{}
	queue := []string{start}
	visited := map[string]bool{
		start: true,
	}

	length := 1
	for len(queue) > 0 {
		length++
		for range queue {
			word := queue[0]
			queue = queue[1:]
			for _, nextWord := range getNextWords(word, dict) {
				if visited[nextWord] {
					continue
				}
				if nextWord == end {
					return length
				}
				visited[nextWord] = true
				queue = append(queue, nextWord)
			}
		}
	}

	return 0
}

func getNextWords(word string, dict map[string]struct{}) []string {
	alphabet := []rune{}
	for r := 'a'; r <= 'z'; r++ {
		alphabet = append(alphabet, r)
	}

	nextWords := []string{}
	for i, v := range word {
		left, right := word[:i], word[i+1:]
		for _, r := range alphabet {
			if v == r {
				continue
			}
			newWord := left + string(r) + right
			if _, ok := dict[newWord]; ok {
				nextWords = append(nextWords, newWord)
			}
		}
	}
	return nextWords
}
```

