

// @lc code=start
func isPerfectSquare(num int) bool {
	if num == 1 {
		return true
	}
	for i := 0; i <= num/2; i++ {
		if num == i*i {
			return true
		}
	}
	return false
}

// @lc code=end
