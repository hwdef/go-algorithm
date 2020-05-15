
func climbStairs(n int) int {
	nums :=[]int{1,2,3}
	for i:=3;i<n;i++{
		nums = append(nums,nums[i-2]+nums[i-1])
	}
	return nums[n-1]
}

