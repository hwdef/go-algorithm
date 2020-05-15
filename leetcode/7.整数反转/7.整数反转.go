
 func reverse(x int) int {
	var x4 int
	if x > 0 {
		x1 := strconv.Itoa(x)
		x2 := []rune(x1)
		for i, j := 0, len(x2)-1; i < j; i, j = i+1, j-1 {
			x2[i], x2[j] = x2[j], x2[i]
		}
		x3 := string(x2)
		x4, _ = strconv.Atoi(x3)
	} else {
		x := 0 - x
		x1 := strconv.Itoa(x)
		x2 := []rune(x1)
		for i, j := 0, len(x2)-1; i < j; i, j = i+1, j-1 {
			x2[i], x2[j] = x2[j], x2[i]
		}
		x3 := string(x2)
		x4, _ = strconv.Atoi(x3)
		x4 = 0 - x4
	}
	if x4 >= int(math.Pow(2, 31)-1) || x4 <= int(math.Pow(-2, 31)) {
		return 0
	} else {
		return x4
	}
}

