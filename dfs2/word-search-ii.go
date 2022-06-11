package dfs2

// https://www.lintcode.com/problem/word-search-ii/

type set map[string]bool

// 上下左右
var Directions = []Coord{
	{-1, 0},
	{1, 0},
	{0, -1},
	{0, 1},
}

type Coord struct {
	x, y int
}

/**
 * @param board: A list of lists of character
 * @param words: A list of string
 * @return: A list of string
 *          we will sort your return value in output
 */
func WordSearchII(board [][]byte, words []string) []string {
	var res []string
	if len(board) == 0 || len(board[0]) == 0 {
		return res
	}

	visited := make([][]bool, len(board))
	for i := range visited {
		visited[i] = make([]bool, len(board[0]))
	}
	wordSet := make(set)
	prefixSet := make(set)
	resSet := make(set)

	for _, word := range words {
		wordSet[word] = true
		for i := 0; i < len(word); i++ {
			prefixSet[string(word[:i+1])] = true
		}
	}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			visited[i][j] = true
			wordSearchIIDfs(board, visited, i, j, string(board[i][j]), wordSet, prefixSet, resSet)
			visited[i][j] = false
		}
	}

	res = make([]string, 0, len(resSet))
	for k := range resSet {
		res = append(res, k)
	}
	return res
}

func wordSearchIIDfs(board [][]byte, visited [][]bool, x int, y int, word string, wordSet set, prefixSet set, resSet set) {
	if !prefixSet[word] {
		return
	}

	if wordSet[word] {
		resSet[word] = true
	}

	for i := 0; i < len(Directions); i++ {
		adjX := x + Directions[i].x
		adjY := y + Directions[i].y
		if !inside(board, adjX, adjY) || visited[adjX][adjY] {
			continue
		}
		visited[adjX][adjY] = true
		wordSearchIIDfs(board, visited, adjX, adjY, word+string(board[adjX][adjY]), wordSet, prefixSet, resSet)
		visited[adjX][adjY] = false
	}
}

func inside(board [][]byte, x, y int) bool {
	return x >= 0 && x < len(board) && y >= 0 && y < len(board[0])
}
