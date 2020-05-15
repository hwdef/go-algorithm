

// @lc code=start
func containsDuplicate(nums []int) bool {
	if len(nums) == 1 {
		return false
	}
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] == nums[j] {
				return true
			}
		}
	}
	return false
}

// @lc code=end
