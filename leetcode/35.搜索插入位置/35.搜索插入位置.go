
func searchInsert(nums []int, target int) int {
	if nums[0] >= target {
		return 0
	}
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] < target && nums[i+1] >= target {
			return i + 1
		}
	}
	if nums[len(nums)-1] <= target {
		return len(nums)
	}
	return 0
}
