
 func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	lengthmin := len(strs[0])
	var min string
	for _, i := range strs {
		if len(i) <= lengthmin {
			lengthmin = len(i)
			min = i
		}
	}
	var minstring []string

	for i := len(min); i > 0; i-- {
		minstring = append(minstring, min[0:i])
	}
	for _, i := range minstring {
		flag := 0
		for _, j := range strs {
			if i != j[0:len(i)] {
				flag = 1
				break
			}
		}
		if flag == 0 {
			return i
		}
	}
	return ""
}

