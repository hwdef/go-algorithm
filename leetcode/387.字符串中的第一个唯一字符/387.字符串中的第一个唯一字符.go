

// @lc code=start
func firstUniqChar(s string) int {
	a := [26]int{0}
	res := []int{}
	for _, i := range s {
		a[i-97]++
	}
	for m, i := range a {
		if i == 1 {
			n := string(m + 97)
			res = append(res, strings.Index(s, n))
		}
	}
	if len(res) == 0 {
		return -1
	}
	sort.Ints(res)
	return res[0]
}


// @lc code=end

