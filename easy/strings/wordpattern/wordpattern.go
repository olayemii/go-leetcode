/*
	290. Word Pattern

	https://leetcode.com/problems/word-pattern/

	Given a pattern and a string s, find if s follows the same pattern.

	Here follow means a full match, such that there is a bijection between a letter in pattern and a non-empty word in s.



	Example 1:

	Input: pattern = "abba", s = "dog cat cat dog"
	Output: true
	Example 2:

	Input: pattern = "abba", s = "dog cat cat fish"
	Output: false
	Example 3:

	Input: pattern = "aaaa", s = "dog cat cat dog"
	Output: false
	Example 4:

	Input: pattern = "abba", s = "dog dog dog dog"
	Output: false


	Constraints:

	1 <= pattern.length <= 300
	pattern contains only lower-case English letters.
	1 <= s.length <= 3000
	s contains only lower-case English letters and spaces ' '.
	s does not contain any leading or trailing spaces.
	All the words in s are separated by a single space.
*/

package main

import (
	"strings"
)

func wordPattern(pattern string, s string) bool {
	patternMap := map[rune]string{};
	valueMap := map[string]bool{};
	splitString := strings.Split(s, " ")
	if len(splitString) != len(pattern) {
		return false;
	}
	for pos, char := range pattern {
		if _, ok := patternMap[char]; ok {
			if patternMap[char] != splitString[pos] {
				return false
			}
		}else{
			if valueMap[splitString[pos]] {
				return false;
			}
			valueMap[splitString[pos]] = true;
			patternMap[char] = splitString[pos]
		}
	}

	return true;
}

/*
	Time Complexity O(n)
	Space Complexity O(n)
*/