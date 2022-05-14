package bfs2

// https://www.lintcode.com/problem/616/

/**
 * @param numCourses: a total of n courses
 * @param prerequisites: a list of prerequisite pairs
 * @return: the course order
 */
func FindOrder(numCourses int, prerequisites [][]int) []int {
	// 构建图，代表（先修课->多个后修课）的映射
	graph := map[int][]int{}
	// 每个点的入度
	inDegree := map[int]int{}
	for _, v := range prerequisites {
		graph[v[1]] = append(graph[v[1]], v[0])
		inDegree[v[0]]++
	}
	queue := []int{}
	for i := 0; i < numCourses; i++ {
		if _, ok := inDegree[i]; !ok {
			queue = append(queue, i)
		}
	}

	numChoose := 0
	topoOrder := make([]int, 0, numCourses)
	for len(queue) > 0 {
		nowPos := queue[0]
		queue = queue[1:]
		topoOrder = append(topoOrder, nowPos)
		numChoose++
		for _, v := range graph[nowPos] {
			inDegree[v]--
			if inDegree[v] == 0 {
				queue = append(queue, v)
			}
		}
	}

	if numChoose != numCourses {
		return []int{}
	}
	return topoOrder
}

func FindOrder1(numCourses int, prerequisites [][]int) []int {
	// head 表示队头，tail表示队尾
	var head = 0
	var tail = -1

	// 构建图，代表（先修课->多个后修课）的映射
	graph := map[int][]int{}
	// 每个点的入度
	inDegree := map[int]int{}
	for _, v := range prerequisites {
		graph[v[1]] = append(graph[v[1]], v[0])
		inDegree[v[0]]++
	}
	queue := []int{}
	for i := 0; i < numCourses; i++ {
		if _, ok := inDegree[i]; !ok {
			queue = append(queue, i)
			tail++
		}
	}

	numChoose := 0
	topoOrder := make([]int, 0, numCourses)
	for head <= tail {
		nowPos := queue[head]
		head++
		topoOrder = append(topoOrder, nowPos)
		numChoose++
		for _, v := range graph[nowPos] {
			inDegree[v]--
			if inDegree[v] == 0 {
				queue = append(queue, v)
				tail++
			}
		}
	}

	if numChoose != numCourses {
		return []int{}
	}
	return topoOrder
}
