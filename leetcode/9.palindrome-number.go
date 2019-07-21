/*
 * @lc app=leetcode id=9 lang=golang
 *
 * [9] Palindrome Number
 */
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	y:=strconv.Itoa(x)
	for i,j:=0,len(y)-1;j>0;i,j=i+1,j-1{
		if y[i]!=y[j]{
			return false
		}
	}
	return true
}

