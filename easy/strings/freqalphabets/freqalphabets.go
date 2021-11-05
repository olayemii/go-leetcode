/*
1309. Decrypt String from Alphabet to Integer Mapping

https://leetcode.com/problems/decrypt-string-from-alphabet-to-integer-mapping/

Given a string s formed by digits ('0' - '9') and '#' . We want to map s to English lowercase characters as follows:

Characters ('a' to 'i') are represented by ('1' to '9') respectively.
Characters ('j' to 'z') are represented by ('10#' to '26#') respectively.
Return the string formed after mapping.

It's guaranteed that a unique mapping will always exist.



Example 1:

Input: s = "10#11#12"
Output: "jkab"
Explanation: "j" -> "10#" , "k" -> "11#" , "a" -> "1" , "b" -> "2".
Example 2:

Input: s = "1326#"
Output: "acz"
Example 3:

Input: s = "25#"
Output: "y"
Example 4:

Input: s = "12345678910#11#12#13#14#15#16#17#18#19#20#21#22#23#24#25#26#"
Output: "abcdefghijklmnopqrstuvwxyz"


Constraints:

1 <= s.length <= 1000
s[i] only contains digits letters ('0'-'9') and '#' letter.
s will be valid string such that mapping is always possible.
*/
package main

import (
	"strconv"
)

func freqAlphabets(s string) string {
    strLen, curr, ans := len(s)-1, "", []rune{}
    for strLen >= 0 {
        if s[strLen] == '#' {
			curr = s[strLen-2:strLen]
			strLen-=3
		}else{
			curr = string(s[strLen])
			strLen--;
		}
		charCode, _ := strconv.Atoi(curr)
		ans = append([]rune{rune('a' + (charCode-1))}, ans...)
    }

	return string(ans);
}

// Time Complexity: O(n)
// Space Complexity: O(n)