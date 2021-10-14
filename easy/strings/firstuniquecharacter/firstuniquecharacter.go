/*
	387. First Unique Character in a String

	https://leetcode.com/problems/first-unique-character-in-a-string/

	Given a string s, find the first non-repeating character in it and return its index. If it does not exist, return -1.



	Example 1:

	Input: s = "leetcode"
	Output: 0
	Example 2:

	Input: s = "loveleetcode"
	Output: 2
	Example 3:

	Input: s = "aabb"
	Output: -1


	Constraints:

	1 <= s.length <= 105
	s consists of only lowercase English letters.
*/
package main

func firstUniqChar(s string) int {
    freqArr := [26]int{}

	for _, char := range s {
		freqArr[char-'a']++
	}

	for pos, char := range s {
		if (freqArr[char-'a'] == 1) {
			return pos
		}
	}
	return -1
}

// Space Complexity O(n)
// Space Complexity O(1)