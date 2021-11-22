/*
1417. Reformat The String

https://leetcode.com/problems/reformat-the-string/

Given alphanumeric string s. (Alphanumeric string is a string consisting of lowercase English letters and digits).

You have to find a permutation of the string where no letter is followed by another letter and no digit is followed by another digit. That is, no two adjacent characters have the same type.

Return the reformatted string or return an empty string if it is impossible to reformat the string.



Example 1:

Input: s = "a0b1c2"
Output: "0a1b2c"
Explanation: No two adjacent characters have the same type in "0a1b2c". "a0b1c2", "0a1b2c", "0c2a1b" are also valid permutations.
Example 2:

Input: s = "leetcode"
Output: ""
Explanation: "leetcode" has only characters so we cannot separate them by digits.
Example 3:

Input: s = "1229857369"
Output: ""
Explanation: "1229857369" has only digits so we cannot separate them by characters.
Example 4:

Input: s = "covid2019"
Output: "c2o0v1i9d"
Example 5:

Input: s = "ab123"
Output: "1a2b3"


Constraints:

1 <= s.length <= 500
s consists of only lowercase English letters and/or digits.
*/

package main

import (
	"math"
)

func reformat(s string) string {
    alphaArr, numArr := []rune{}, []rune{}
    
    for _, char := range s {
        if char >= 'a' && char <= 'z' {
            alphaArr = append(alphaArr, char)
        }else{
            numArr = append(numArr, char)            
        }
    }
    
    if len(s) == 1 {
        return s
    }
	
    if (len(alphaArr) == 0 || len(numArr) == 0) || math.Abs(float64(len(alphaArr)) - float64(len(numArr))) > 1 {
        return ""
    }
        
    if len(alphaArr) > len(numArr) {
        return permutateString(string(alphaArr), string(numArr))        
    }
    return permutateString(string(numArr), string(alphaArr))
}

func permutateString(string1 string, string2 string) string {
    i, j, ans := 0, 0, ""
    
    for i < len(string1) || j < len(string2) {
        if i < len(string1) {
            ans+=string(string1[i])
        }
        if j < len(string2) {
            ans+=string(string2[j])
        }
        
        i++
        j++
    }
    return ans
}

// Time Complexity: O(n)
// Space Complexity: O(n)