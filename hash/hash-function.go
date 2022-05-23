package hash

// https://www.lintcode.com/problem/hash-function

/**
 * @param key: A string you should hash
 * @param hASH_SIZE: An integer
 * @return: An integer
 */
func HashCode(key string, hASH_SIZE int) int {
	var ans int
	for _, r := range key {
		ans = (ans*33 + int(r)) % hASH_SIZE
	}
	return ans
}
