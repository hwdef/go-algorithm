

// @lc code=start
func thirdMax(nums []int) int {
	const INT_MAX = int(^uint(0) >> 1)
	const INT_MIN = ^INT_MAX
	一, 二, 三 := INT_MIN, INT_MIN, INT_MIN
	for i := 0; i < len(nums); i++ {
		if nums[i] > 一 {
			三 = 二
			二 = 一
			一 = nums[i]
		} else if nums[i] > 二 && nums[i] != 一 {
			三 = 二
			二 = nums[i]
		} else if nums[i] > 三 && nums[i] != 二 && nums[i] != 一 {
			三 = nums[i]
		}
	}
	if 二 == INT_MIN || 三 == INT_MIN {
		return 一
	}
	fmt.Println(一, 二, 三)
	return 三
}
// @lc code=end

