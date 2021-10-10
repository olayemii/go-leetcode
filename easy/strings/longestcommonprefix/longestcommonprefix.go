package main

/*
14. Longest Common Prefix

https://leetcode.com/problems/longest-common-prefix/

Write a function to find the longest common prefix string amongst an array of strings.

If there is no common prefix, return an empty string "".



Example 1:

Input: strs = ["flower","flow","flight"]
Output: "fl"
Example 2:

Input: strs = ["dog","racecar","car"]
Output: ""
Explanation: There is no common prefix among the input strings.


Constraints:

1 <= strs.length <= 200
0 <= strs[i].length <= 200
strs[i] consists of only lower-case English letters.
*/

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return "";
	}
	longestPrefix := "";
	for idx, item := range strs[0] {
        for i := 1;  i < len(strs); i++ {
            if (len(strs[i]) <= idx || rune(strs[i][idx]) != item) {
				return longestPrefix;
			}
		}
		longestPrefix += string(item)
	}

	return longestPrefix
}

// Space complexity O(n)
// TIme complexity O(n * m) - n is number of strings in strs and m is length of longest character