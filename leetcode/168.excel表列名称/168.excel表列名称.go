

// @lc code=start
func convertToTitle(n int) string {
	res := ""
	for n != 0 {
		n--
		shang := n / 26
		yu := n % 26
		res = string('A'+yu) + res
		n = shang
	}
	return res
}
// @lc code=end

