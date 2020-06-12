
func twoSum(nums []int, target int) []int {
	hashmap := make(map[int]int, len(nums))
	for k, v := range nums {
		if x, y := hashmap[target-v]; y {
			return []int{x, k}
		}
		hashmap[v] = k
	}
	return []int{-1, -1}
}
