package bfs2

import "fmt"

// https://www.lintcode.com/problem/605

/**
 * @param org: a permutation of the integers from 1 to n
 * @param seqs: a list of sequences
 * @return: true if it can be reconstructed only one or false
 */
func SequenceReconstruction(org []int, seqs [][]int) bool {
	graph := map[int][]int{}
	inDegree := map[int]int{}
	for _, seq := range seqs {
		for i, node := range seq {
			if i == 0 {
				if _, ok := inDegree[node]; !ok {
					inDegree[node] = 0
				}
			} else {
				graph[seq[i-1]] = append(graph[seq[i-1]], node)
				inDegree[node]++
			}
		}
	}

	queue := []int{}
	for k, v := range inDegree {
		if v == 0 {
			queue = append(queue, k)
		}
	}

	topo := []int{}
	numChoose := 0
	for len(queue) > 0 {
		if len(queue) > 1 {
			return false
		}
		nowPos := queue[0]
		queue = queue[1:]
		topo = append(topo, nowPos)
		numChoose++

		for _, v := range graph[nowPos] {
			inDegree[v]--
			if inDegree[v] == 0 {
				queue = append(queue, v)
			}
		}
	}

	if numChoose != len(org) {
		return false
	}
	return fmt.Sprint(org) == fmt.Sprint(topo)
}
