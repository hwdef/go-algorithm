/*
 * @lc app=leetcode id=14 lang=golang
 *
 * [14] Longest Common Prefix
 *
 * https://leetcode.com/problems/longest-common-prefix/description/
 *
 * algorithms
 * Easy (33.82%)
 * Likes:    1462
 * Dislikes: 1385
 * Total Accepted:    500.3K
 * Total Submissions: 1.5M
 * Testcase Example:  '["flower","flow","flight"]'
 *
 * Write a function to find the longest common prefix string amongst an array
 * of strings.
 *
 * If there is no common prefix, return an empty string "".
 *
 * Example 1:
 *
 *
 * Input: ["flower","flow","flight"]
 * Output: "fl"
 *
 *
 * Example 2:
 *
 *
 * Input: ["dog","racecar","car"]
 * Output: ""
 * Explanation: There is no common prefix among the input strings.
 *
 *
 * Note:
 *
 * All given inputs are in lowercase letters a-z.
 *
 */
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

