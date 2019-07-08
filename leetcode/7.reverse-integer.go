import (
	"math"
	"strconv"
)

/*
 * @lc app=leetcode id=7 lang=golang
 *
 * [7] Reverse Integer
 *
 * https://leetcode.com/problems/reverse-integer/description/
 *
 * algorithms
 * Easy (25.33%)
 * Likes:    2195
 * Dislikes: 3357
 * Total Accepted:    705.4K
 * Total Submissions: 2.8M
 * Testcase Example:  '123'
 *
 * Given a 32-bit signed integer, reverse digits of an integer.
 *
 * Example 1:
 *
 *
 * Input: 123
 * Output: 321
 *
 *
 * Example 2:
 *
 *
 * Input: -123
 * Output: -321
 *
 *
 * Example 3:
 *
 *
 * Input: 120
 * Output: 21
 *
 *
 * Note:
 * Assume we are dealing with an environment which could only store integers
 * within the 32-bit signed integer range: [−2^31,  2^31 − 1]. For the purpose
 * of this problem, assume that your function returns 0 when the reversed
 * integer overflows.
 *
 */
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
