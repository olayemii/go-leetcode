/*
1832. Check if the Sentence Is Pangram

https://leetcode.com/problems/check-if-the-sentence-is-pangram/

A pangram is a sentence where every letter of the English alphabet appears at least once.

Given a string sentence containing only lowercase English letters, return true if sentence is a pangram, or false otherwise.



Example 1:

Input: sentence = "thequickbrownfoxjumpsoverthelazydog"
Output: true
Explanation: sentence contains at least one of every letter of the English alphabet.
Example 2:

Input: sentence = "leetcode"
Output: false


Constraints:

1 <= sentence.length <= 1000
sentence consists of lowercase English letters.
*/
package main

func checkIfPangram(sentence string) bool {
	freqArr := [26]int{}

	for _, char := range sentence {
		freqArr[char - 'a']++
	}

	for _, val := range freqArr {
		if val == 0 {
			return false
		}
	}

	return true
}

// Space Complexity: O(1)
// Time Complexity: O(n)