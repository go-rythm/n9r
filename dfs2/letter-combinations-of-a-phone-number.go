package dfs2

// https://www.lintcode.com/problem/425/

/**
 * @param digits: A digital string
 * @return: all possible letter combinations
 *          we will sort your return value in output
 */

func LetterCombinations(digits string) []string {
	paths := []string{}
	if len(digits) == 0 {
		return paths
	}
	dict := []string{"abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}
	dfs(dict, digits, 0, nil, &paths)
	return paths
}

func dfs(dict []string, digits string, idx int, path []rune, paths *[]string) {
	if idx == len(digits) {
		*paths = append(*paths, string(path))
		return
	}

	for _, letter := range dict[digits[idx]-'2'] {
		path = append(path, letter)
		dfs(dict, digits, idx+1, path, paths)
		path = path[:len(path)-1]
	}
}
