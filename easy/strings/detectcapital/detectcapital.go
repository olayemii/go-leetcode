/*
520. Detect Capital

https://leetcode.com/problems/detect-capital/

We define the usage of capitals in a word to be right when one of the following cases holds:

All letters in this word are capitals, like "USA".
All letters in this word are not capitals, like "leetcode".
Only the first letter in this word is capital, like "Google".
Given a string word, return true if the usage of capitals in it is right.



Example 1:

Input: word = "USA"
Output: true
Example 2:

Input: word = "FlaG"
Output: false


Constraints:

1 <= word.length <= 100
word consists of lowercase and uppercase English letters.
*/

package main

import (
	"strings"
)

func detectCapitalUse(word string) bool {
	return strings.ToUpper(word) == word || strings.ToLower(word[1:]) == word[1:]
}
// Space Complexity: O(1)
// Time Complexity: O(n)