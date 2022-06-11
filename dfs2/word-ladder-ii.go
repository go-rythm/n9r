package dfs2

// https://www.lintcode.com/problem/121/

/**
 * @param start: a string
 * @param end: a string
 * @param dict: a set of string
 * @return: a list of lists of string
 *          we will sort your return value in output
 */
func FindLadders(start string, end string, dict map[string]struct{}) [][]string {
	dict[start] = struct{}{}
	dict[end] = struct{}{}

	// 起点到某个点的最短路径长度
	dist := make(map[string]int)
	// 从某一个词开始，可以到达的不绕远的下一个词的集合
	fromToMap := make(map[string][]string)
	ladderBfs(fromToMap, dist, start, end, dict)

	var res [][]string
	ladderDfs(fromToMap, start, end, dist[end], &[]string{}, &res)
	return res
}

func ladderBfs(fromToMap map[string][]string, dist map[string]int, start string, end string, dict map[string]struct{}) {
	dist[start] = 0
	q := []string{start}

	for len(q) > 0 {
		curWord := q[0]
		q = q[1:]
		for _, nextWord := range getNextWords(curWord, dict) {
			_, ok := dist[nextWord]
			if !ok || dist[nextWord] == dist[curWord]+1 {
				fromToMap[curWord] = append(fromToMap[curWord], nextWord)
			}

			if !ok {
				dist[nextWord] = dist[curWord] + 1
				q = append(q, nextWord)
			}
		}
	}
}

func ladderDfs(fromToMap map[string][]string, curWord string, end string, minLen int, path *[]string, res *[][]string) {
	if len(*path) == minLen+1 {
		return
	}
	*path = append(*path, curWord)
	if curWord == end {
		*res = append(*res, append([]string{}, *path...))
	} else {
		for _, nextWord := range fromToMap[curWord] {
			ladderDfs(fromToMap, nextWord, end, minLen, path, res)
		}
	}
	*path = (*path)[:len(*path)-1]
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
