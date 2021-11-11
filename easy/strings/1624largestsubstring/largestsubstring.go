/*
1624. Largest Substring Between Two Equal Characters

https://leetcode.com/problems/largest-substring-between-two-equal-characters/

Given a string s, return the length of the longest substring between two equal characters, excluding the two characters. If there is no such substring return -1.

A substring is a contiguous sequence of characters within a string.



Example 1:

Input: s = "aa"
Output: 0
Explanation: The optimal substring here is an empty substring between the two 'a's.
Example 2:

Input: s = "abca"
Output: 2
Explanation: The optimal substring here is "bc".
Example 3:

Input: s = "cbzxy"
Output: -1
Explanation: There are no characters that appear twice in s.
Example 4:

Input: s = "cabbac"
Output: 4
Explanation: The optimal substring here is "abba". Other non-optimal substrings include "bb" and "".


Constraints:

1 <= s.length <= 300
s contains only lowercase English letters.
*/
package main

import (
	"math"
)

func maxLengthBetweenEqualCharacters(s string) int {
	max, i, j := -1, 0, len(s)-1;
	for i < len(s) {
        j = len(s) - 1
		for s[j] != s[i] && j > i {
			j--
		}

		count := j - i - 1

		max = int(math.Max(float64(count), float64(max)))
		i++
	}

	return max
}

// Space Complexity: O(1)
// Time Complexity: O(n)