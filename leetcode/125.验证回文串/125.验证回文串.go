

// @lc code=start
func isPalindrome(s string) bool {
	var str string
	for _, i := range s {
		if unicode.IsLetter(i) {
			str = str + string(i)
			continue
		}
		if unicode.IsDigit(i) {
			str = str + string(i)
			continue
		}
	}
	str = strings.ToLower(str)
	for i, j := 0, len(str)-1; i < len(str)/2; i, j = i+1, j-1 {
		if str[i] != str[j] {
			return false
		}
	}
	return true
}

// @lc code=end

