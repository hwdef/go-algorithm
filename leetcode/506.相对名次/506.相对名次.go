

// @lc code=start
func findRelativeRanks(nums []int) []string {
	nums备份 := []int{}
	for _, i := range nums {
		nums备份 = append(nums备份, i)
	}
	res := []string{}
	名次 := make(map[int]int)
	sort.Ints(nums)
	key := len(nums)
	for _, value := range nums {
		名次[value] = key
		key--
	}
	标记 := 0
	for _, i := range nums备份 {
		switch 名次[i] {
		case 1:
			res = append(res, "Gold Medal")
			标记 = 1
		case 2:
			res = append(res, "Silver Medal")
			标记 = 1
		case 3:
			res = append(res, "Bronze Medal")
			标记 = 1
		}
		if 标记 == 0 {
			res = append(res, strconv.Itoa(名次[i]))
		}
		标记 = 0
	}
	return res
}

// @lc code=end
