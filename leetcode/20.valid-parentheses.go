/*
 * @lc app=leetcode id=20 lang=golang
 *
 * [20] Valid Parentheses
 *
 * https://leetcode.com/problems/valid-parentheses/description/
 *
 * algorithms
 * Easy (36.86%)
 * Likes:    3134
 * Dislikes: 149
 * Total Accepted:    640.2K
 * Total Submissions: 1.7M
 * Testcase Example:  '"()"'
 *
 * Given a string containing just the characters '(', ')', '{', '}', '[' and
 * ']', determine if the input string is valid.
 * 
 * An input string is valid if:
 * 
 * 
 * Open brackets must be closed by the same type of brackets.
 * Open brackets must be closed in the correct order.
 * 
 * 
 * Note that an empty string is also considered valid.
 * 
 * Example 1:
 * 
 * 
 * Input: "()"
 * Output: true
 * 
 * 
 * Example 2:
 * 
 * 
 * Input: "()[]{}"
 * Output: true
 * 
 * 
 * Example 3:
 * 
 * 
 * Input: "(]"
 * Output: false
 * 
 * 
 * Example 4:
 * 
 * 
 * Input: "([)]"
 * Output: false
 * 
 * 
 * Example 5:
 * 
 * 
 * Input: "{[]}"
 * Output: true
 * 
 * 
 */
func isValid(s string) bool {
	if len(s)==0{
		return true
	}
	if string(s[0])==")" || string(s[0])=="}"||string(s[0])=="]"{
		return false
	}
	if len(s)%2!=0{
		return false
	}
	var a []string
	for _,i:=range s{
		if string(i) == "(" || string(i) == "{" || string(i) == "[" {
			a = append(a,string(i))
		} else{
			if a[len(a)-1]=="("{
				if string(i) != ")"{
					return false
				} else{
					a = a[:len(a)-1]
				}
			}else if a[len(a)-1]=="{"{
				if string(i)!="}"{
					return false
				} else{
					a = a[:len(a)-1]
				}
			}else if a[len(a)-1]=="["{
				if string(i)!="]"{
					return false
				} else{
					a = a[:len(a)-1]
				}
			}
		}
	}
	if len(a) != 0{
		return false
	} else{
		return true
	}
}
