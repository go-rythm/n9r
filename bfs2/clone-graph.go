package bfs2

// https://www.lintcode.com/problem/137/

/**
 * @param node: A undirected graph node
 * @return: A undirected graph node
 */

type UndirectedGraphNode struct {
	Label     int
	Neighbors []*UndirectedGraphNode
}

func CloneGraph(node *UndirectedGraphNode) *UndirectedGraphNode {
	if node == nil {
		return nil
	}
	nodeSet := FindNodesByBfs(node)
	mapping := CopyNodes(nodeSet)
	CopyEdges(nodeSet, mapping)
	return mapping[node]
}

type nodeSet map[*UndirectedGraphNode]bool

func FindNodesByBfs(node *UndirectedGraphNode) nodeSet {
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
			visited[neighbor] = true
			queue = append(queue, neighbor)
		}
	}
	return visited
}

func FindNodesByBfsWithLevel(node *UndirectedGraphNode) nodeSet {
	queue := []*UndirectedGraphNode{node}
	visited := nodeSet{
		node: true,
	}
	for len(queue) > 0 {
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

func FindNodesByBfsWrong(node *UndirectedGraphNode) nodeSet {
	queue := []*UndirectedGraphNode{node}
	visited := nodeSet{}
	for len(queue) > 0 {
		curNode := queue[0]
		queue = queue[1:]
		visited[curNode] = true
		for _, neighbor := range curNode.Neighbors {
			if visited[neighbor] {
				continue
			}
			// visited[neighbor] = true
			queue = append(queue, neighbor)
		}
	}
	return visited
}

func CopyNodes(nodeSet nodeSet) map[*UndirectedGraphNode]*UndirectedGraphNode {
	mapping := make(map[*UndirectedGraphNode]*UndirectedGraphNode)
	for node := range nodeSet {
		mapping[node] = &UndirectedGraphNode{
			Label: node.Label,
		}
	}
	return mapping
}

func CopyEdges(nodeSet nodeSet, mapping map[*UndirectedGraphNode]*UndirectedGraphNode) {
	for node := range nodeSet {
		newNode := mapping[node]
		for _, neighbor := range node.Neighbors {
			newNeighbor := mapping[neighbor]
			newNode.Neighbors = append(newNode.Neighbors, newNeighbor)
		}
	}
}
