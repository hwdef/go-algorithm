

// @lc code=start
func dailyTemperatures(T []int) []int {
	res := []int{}
	for k, v := range T {
		flag := 0
		for i := k + 1; i < len(T); i++ {
			if T[i] > v {
				res = append(res, i-k)
				flag = 1
				break
			}
		}
		if flag == 0 {
			res = append(res, 0)
		}
	}
	return res
}
// @lc code=end

