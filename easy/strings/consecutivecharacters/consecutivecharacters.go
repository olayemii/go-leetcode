/*
1446. Consecutive Characters

https://leetcode.com/problems/consecutive-characters/

The power of the string is the maximum length of a non-empty substring that contains only one unique character.

Given a string s, return the power of s.



Example 1:

Input: s = "leetcode"
Output: 2
Explanation: The substring "ee" is of length 2 with the character 'e' only.
Example 2:

Input: s = "abbcccddddeeeeedcba"
Output: 5
Explanation: The substring "eeeee" is of length 5 with the character 'e' only.
Example 3:

Input: s = "triplepillooooow"
Output: 5
Example 4:

Input: s = "hooraaaaaaaaaaay"
Output: 11
Example 5:

Input: s = "tourist"
Output: 1


Constraints:

1 <= s.length <= 500
s consists of only lowercase English letters.
*/
package main

import "math"

func maxPower(s string) int {
    curr := s[0]
    count := 0
    max := 0
    for pos := range s {
        if s[pos] == curr {
            count++
        }else{
            max = int(math.Max(float64(count), float64(max)))
            count = 1;
            curr = s[pos]
        }
    }
    
    return int(math.Max(float64(count), float64(max)))
}

// Space Complexity: O(1)
// Time Complexity: O(n)