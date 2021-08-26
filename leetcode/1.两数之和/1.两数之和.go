
 func twoSum(nums []int, target int) []int {
	var a []int
	b := 0
	for key, i := range nums {
		for key2, j := range nums {
			if (target == (i + j)) && (key != key2) {
				a = append(a, key)
				a = append(a, key2)
				b = 1
				break
			}
		}
		if b == 1 {
			break
		}
	}
	return a
}

//test