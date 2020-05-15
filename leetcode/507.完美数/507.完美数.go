

// @lc code=start
func checkPerfectNumber(num int) bool {
	if num == 1 {
		return false
	}
	var nums []int
	nums = append(nums, 1)
	for i := 2; i <= int(math.Sqrt(float64(num))); i++ {
		if num%i == 0 {
			nums = append(nums, i)
			nums = append(nums, num/i)
		}
	}
	sum := 0
	for _, i := range nums {
		sum += i
	}
	if sum == num {
		return true
	} else {
		return false
	}
}

// @lc code=end

