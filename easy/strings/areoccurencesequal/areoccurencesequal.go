/*
1941. Check if All Characters Have Equal Number of Occurrences

https://leetcode.com/problems/check-if-all-characters-have-equal-number-of-occurrences/

Given a string s, return true if s is a good string, or false otherwise.

A string s is good if all the characters that appear in s have the same number of occurrences (i.e., the same frequency).



Example 1:

Input: s = "abacbc"
Output: true
Explanation: The characters that appear in s are 'a', 'b', and 'c'. All characters occur 2 times in s.
Example 2:

Input: s = "aaabb"
Output: false
Explanation: The characters that appear in s are 'a' and 'b'.
'a' occurs 3 times while 'b' occurs 2 times, which is not the same number of times.


Constraints:

1 <= s.length <= 1000
s consists of lowercase English letters.
*/
package main

func areOccurrencesEqual(s string) bool {
	hashMap := map[rune]int{}

	for _, char := range s {
		if _,ok := hashMap[char]; ok {
			hashMap[char]++
		}else{
			hashMap[char] = 1
		}
	}

	for _, value := range hashMap {
		if value != hashMap[rune(s[0])] {
			return false
		}
	}

	return true
}

// Space Complexity: O(1)
// Time Complexity: O(n)