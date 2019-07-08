/*
 * @lc app=leetcode id=1 lang=golang
 *
 * [1] Two Sum
 *
 * https://leetcode.com/problems/two-sum/description/
 *
 * algorithms
 * Easy (44.23%)
 * Likes:    10911
 * Dislikes: 369
 * Total Accepted:    1.8M
 * Total Submissions: 4.2M
 * Testcase Example:  '[2,7,11,15]\n9'
 *
 * Given an array of integers, return indices of the two numbers such that they
 * add up to a specific target.
 *
 * You may assume that each input would have exactly one solution, and you may
 * not use the same element twice.
 *
 * Example:
 *
 *
 * Given nums = [2, 7, 11, 15], target = 9,
 *
 * Because nums[0] + nums[1] = 2 + 7 = 9,
 * return [0, 1].
 *
 *
 */
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
