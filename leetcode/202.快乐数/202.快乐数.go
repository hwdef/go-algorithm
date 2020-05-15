

// @lc code=start
func isHappy(n int) bool {
	for {
		var res int
		m := strconv.Itoa(n)
		for _, i := range m {
			j, _ := strconv.Atoi(string(i))
			res = res + j*j
		}
		n = res
		if n == 4 {
			return false
		}
		if n == 1 {
			return true
		}
	}
}
// @lc code=end

