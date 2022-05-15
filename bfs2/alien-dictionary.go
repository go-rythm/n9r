package bfs2

import (
	"container/heap"
	"sort"
)

// https://www.lintcode.com/problem/892/

/**
 * @param words: a list of words
 * @return: a string which is correct order
 */
func AlienOrder1(words []string) string {
	// 初始化图
	graph := map[byte][]byte{}
	for _, word := range words {
		for _, r := range word {
			graph[byte(r)] = []byte{}
		}
	}

	// 初始化图中的边
	for i := 0; i < len(words)-1; i++ {
		min := len(words[i])
		if len(words[i+1]) < min {
			min = len(words[i+1])
		}

		for j := 0; j < min; j++ {
			if words[i][j] != words[i+1][j] {
				graph[words[i][j]] = append(graph[words[i][j]], words[i+1][j])
				break
			}
			if j == min-1 {
				if len(words[i]) > len(words[j]) {
					return ""
				}
			}
		}
	}

	// 初始化入度
	inDegrees := map[byte]int{}
	for node := range graph {
		inDegrees[node] = 0
	}
	for _, neighbors := range graph {
		for _, neighbor := range neighbors {
			inDegrees[neighbor]++
		}
	}

	queue := []byte{}
	for node, inDegree := range inDegrees {
		if inDegree == 0 {
			queue = append(queue, node)
		}
	}

	topoOrder := []byte{}
	for len(queue) > 0 {
		sort.Slice(queue, func(i int, j int) bool { return queue[i] < queue[j] })

		curNode := queue[0]
		queue = queue[1:]
		topoOrder = append(topoOrder, curNode)

		for _, neighbor := range graph[curNode] {
			inDegrees[neighbor]--
			if inDegrees[neighbor] == 0 {
				queue = append(queue, neighbor)
			}
		}
	}

	if len(topoOrder) != len(graph) {
		return ""
	}

	return string(topoOrder)
}

func AlienOrder(words []string) string {
	// 初始化图
	graph := map[byte][]byte{}
	for _, word := range words {
		for _, r := range word {
			graph[byte(r)] = []byte{}
		}
	}

	// 初始化图中的边
	for i := 0; i < len(words)-1; i++ {
		min := len(words[i])
		if len(words[i+1]) < min {
			min = len(words[i+1])
		}

		for j := 0; j < min; j++ {
			if words[i][j] != words[i+1][j] {
				graph[words[i][j]] = append(graph[words[i][j]], words[i+1][j])
				break
			}
			if j == min-1 {
				if len(words[i]) > len(words[j]) {
					return ""
				}
			}
		}
	}

	// 初始化入度
	inDegrees := map[byte]int{}
	for node := range graph {
		inDegrees[node] = 0
	}
	for _, neighbors := range graph {
		for _, neighbor := range neighbors {
			inDegrees[neighbor]++
		}
	}

	queue := &ByteHeap{}
	heap.Init(queue)
	for node, inDegree := range inDegrees {
		if inDegree == 0 {
			heap.Push(queue, node)
		}
	}

	topoOrder := []byte{}
	for queue.Len() > 0 {
		curNode := heap.Pop(queue).(byte)
		topoOrder = append(topoOrder, curNode)

		for _, neighbor := range graph[curNode] {
			inDegrees[neighbor]--
			if inDegrees[neighbor] == 0 {
				heap.Push(queue, neighbor)
			}
		}
	}

	if len(topoOrder) != len(graph) {
		return ""
	}

	return string(topoOrder)
}

// A ByteHeap is a min-heap of bytes.
type ByteHeap []byte

func (h ByteHeap) Len() int           { return len(h) }
func (h ByteHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h ByteHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *ByteHeap) Push(x any) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(byte))
}

func (h *ByteHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
