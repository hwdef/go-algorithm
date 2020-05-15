
func lengthOfLastWord(s string) int {
	if len(s) < 1 {
		return 0
	}
	str := ""
	back := len(s) - 1
	for i := back; i >= 0; i-- {
		if string(s[i]) != " " {
			break
		} else {
			back--
		}
	}
	for i := back; i >= 0; i-- {
		if string(s[i]) == " " {
			break
		}
		str = str + string(s[i])
	}
	return len(str)
}
