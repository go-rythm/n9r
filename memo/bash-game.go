package memo

// https://www.lintcode.com/problem/bash-game/

/**
 * @param n: an integer
 * @return: whether you can win the game given the number of stones in the heap
 */
func CanWinBash(n int) bool {
	memo := make(map[int]bool)
	return memoSearch(n, memo)
}

func memoSearch(n int, memo map[int]bool) bool {
	if n <= 3 {
		return true
	}
	if v, ok := memo[n]; ok {
		return v
	}
	for i := 1; i < 4; i++ {
		if !memoSearch(n-i, memo) {
			memo[n] = true
			return true
		}
	}
	memo[n] = false
	return false
}
