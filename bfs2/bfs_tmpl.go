package bfs2

type Node struct {
	Val       int
	Neighbors []*Node
}

func BfsTmpl(node *Node) {
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
}
