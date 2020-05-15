
func plusOne(digits []int) []int {
	for long := len(digits) - 1; long >= 0; long-- {
		if digits[long]+1 < 10 {
			digits[long]++
			break
		} else {
			digits[long] = 0
		}
	}

	if digits[0] == 0 {
		digits = append([]int{1}, digits...)
		return digits
	} else {
		return digits
	}
}
