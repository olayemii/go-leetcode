/*
2062. Count Vowel Substrings of a String

https://leetcode.com/problems/count-vowel-substrings-of-a-string/

A substring is a contiguous (non-empty) sequence of characters within a string.

A vowel substring is a substring that only consists of vowels ('a', 'e', 'i', 'o', and 'u') and has all five vowels present in it.

Given a string word, return the number of vowel substrings in word.



Example 1:

Input: word = "aeiouu"
Output: 2
Explanation: The vowel substrings of word are as follows (underlined):
- "aeiouu"
- "aeiouu"
Example 2:

Input: word = "unicornarihan"
Output: 0
Explanation: Not all 5 vowels are present, so there are no vowel substrings.
Example 3:

Input: word = "cuaieuouac"
Output: 7
Explanation: The vowel substrings of word are as follows (underlined):
- "cuaieuouac"
- "cuaieuouac"
- "cuaieuouac"
- "cuaieuouac"
- "cuaieuouac"
- "cuaieuouac"
- "cuaieuouac"
Example 4:

Input: word = "bbaeixoubb"
Output: 0
Explanation: The only substrings that contain all five vowels also contain consonants, so there are no vowel substrings.


Constraints:

1 <= word.length <= 100
word consists of lowercase English letters only.
*/
package main

import (
	"strings"
)

func countVowelSubstrings(word string) int {
    i, j, count := 0, 0, 0

    for i < len(word)  {

		if j < len(word) && (j < len(word)  && word[j] == 'a' || word[j] == 'e' || word[j] == 'i' || word[j] == 'o' || word[j] == 'u') {
			if containsVowels(word[i:j+1]) {
				count++;
			}
			j++
		}else{
			i++;
			j = i
		}
	}
	return count;
}

func containsVowels(word string) bool {
	return strings.Index(word, "a") > - 1 && strings.Index(word, "e") > - 1 && strings.Index(word, "i") > - 1 && strings.Index(word, "o") > - 1 && strings.Index(word, "u") > - 1;
}

// Time Complexity: O(n^2)
// Space Complexity: O(1)

