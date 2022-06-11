package dfs2

// https://www.jiuzhang.com/problem/string-permutation-ii/

import "sort"

func StringPermutation(str string) []string {
	permutations := []string{}
	if str == "" {
		return permutations
	}
	chars := []rune(str)
	sort.Slice(chars, func(i, j int) bool {
		return chars[i] < chars[j]
	})

	visited := make([]bool, len(chars))
	stringPermutationDfs(chars, visited, []rune{}, &permutations)
	return permutations
}

func stringPermutationDfs(chars []rune, visited []bool, permutation []rune, permutations *[]string) {
	if len(permutation) == len(chars) {
		*permutations = append(*permutations, string(permutation))
	}

	for i, char := range chars {
		if visited[i] {
			continue
		}

		if i > 0 && char == chars[i-1] && !visited[i-1] {
			continue
		}

		visited[i] = true
		permutation = append(permutation, char)
		stringPermutationDfs(chars, visited, permutation, permutations)
		permutation = permutation[:len(permutation)-1]

		visited[i] = false
	}
}
