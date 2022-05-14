package bfs2

// https://www.lintcode.com/problem/120/

/**
 * @param start: a string
 * @param end: a string
 * @param dict: a set of string
 * @return: An integer
 */
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
