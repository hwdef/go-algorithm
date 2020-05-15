

// @lc code=start
func isAnagram(s string, t string) bool {
	计数 := make(map[rune]int)
	for _, i := range s {
		计数[i]++
	}
	for _, i := range t {
		计数[i]--
	}
	for _, j := range 计数 {
		if j != 0 {
			return false
		}
	}
	return true
}

// @lc code=end

