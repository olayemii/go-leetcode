/*
1796. Second Largest Digit in a String

https://leetcode.com/problems/second-largest-digit-in-a-string/

Given an alphanumeric string s, return the second largest numerical digit that appears in s, or -1 if it does not exist.

An alphanumeric string is a string consisting of lowercase English letters and digits.



Example 1:

Input: s = "dfa12321afd"
Output: 2
Explanation: The digits that appear in s are [1, 2, 3]. The second largest digit is 2.
Example 2:

Input: s = "abc1111"
Output: -1
Explanation: The digits that appear in s are [1]. There is no second largest digit.


Constraints:

1 <= s.length <= 500
s consists of only lowercase English letters and/or digits.
*/
package main

import (
	"math"
)

func secondHighest(s string) int {
    queue := [2]int{math.MinInt32, math.MinInt32}

	for _, char := range s {
		if char >= '0' && char <= '9' {
			intVal := int(char - '0');
			if intVal > queue[1] {
				queue[0] = queue[1]
				queue[1] = intVal
			}else
			if intVal > queue[0] && intVal < queue[1] {
				queue[0] = intVal
			}
		}
	}
	if queue[0] == math.MinInt32 || queue[1] == math.MinInt32 {
		return -1
	}

	return queue[0]
}

// Space Complexity: O(1)
// Time Complexity: O(n)