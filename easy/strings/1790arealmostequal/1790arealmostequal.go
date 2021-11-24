/*
1790. Check if One String Swap Can Make Strings Equal

https://leetcode.com/problems/check-if-one-string-swap-can-make-strings-equal/

You are given two strings s1 and s2 of equal length. A string swap is an operation where you choose two indices in a string (not necessarily different) and swap the characters at these indices.

Return true if it is possible to make both strings equal by performing at most one string swap on exactly one of the strings. Otherwise, return false.

 

Example 1:

Input: s1 = "bank", s2 = "kanb"
Output: true
Explanation: For example, swap the first character with the last character of s2 to make "bank".
Example 2:

Input: s1 = "attack", s2 = "defend"
Output: false
Explanation: It is impossible to make them equal with one string swap.
Example 3:

Input: s1 = "kelb", s2 = "kelb"
Output: true
Explanation: The two strings are already equal, so no string swap operation is required.
Example 4:

Input: s1 = "abcd", s2 = "dcba"
Output: false
 

Constraints:

1 <= s1.length, s2.length <= 100
s1.length == s2.length
s1 and s2 consist of only lowercase English letters.
*/

package main

func areAlmostEqual(s1 string, s2 string) bool {
    swaps, freqArr := 0, [26]int{}
    if s1 == s2 {
        return true
    }
    for pos := range s1 {
        freqArr[s1[pos]-'a']++
        freqArr[s2[pos]-'a']--
        if s1[pos] != s2[pos] {
            swaps++
        }
    }
    for _, v := range freqArr {
        if v != 0 {
            return false
        }
    }
    return swaps <= 2
}

// Space Complexity: O(1)
// Time Complexity: O(n)