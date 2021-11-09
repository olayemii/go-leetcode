/*
1047. Remove All Adjacent Duplicates In String

https://leetcode.com/problems/remove-all-adjacent-duplicates-in-string/

You are given a string s consisting of lowercase English letters. A duplicate removal consists of choosing two adjacent and equal letters and removing them.

We repeatedly make duplicate removals on s until we no longer can.

Return the final string after all such duplicate removals have been made. It can be proven that the answer is unique.



Example 1:

Input: s = "abbaca"
Output: "ca"
Explanation:
For example, in "abbaca" we could remove "bb" since the letters are adjacent and equal, and this is the only possible move.  The result of this move is that the string is "aaca", of which only "aa" is possible, so the final string is "ca".
Example 2:

Input: s = "azxxzy"
Output: "ay"


Constraints:

1 <= s.length <= 105
s consists of lowercase English letters.
*/

package main

func removeDuplicatesBad(s string) string {
    i := 0
    for i < len(s) {
        if i+1 < len(s) && s[i] == s[i+1] {
			s = s[0:i]+s[i+2:]
			i = 0
        }else{
			i++
        }
    }
	return string(s)
}

func removeDuplicates(s string) string{
	stack := []rune{};
	for _, char := range s {
		if len(stack) > 0 && stack[len(stack)-1] == char {
			stack = stack[0: len(stack) - 1]
		}else{
			stack = append(stack, char)
		}
	}
	return string(stack)
}
// Time Complexity: O(n)
// Space Complexity: O(n)