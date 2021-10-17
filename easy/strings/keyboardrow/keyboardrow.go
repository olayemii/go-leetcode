/*
https://leetcode.com/problems/keyboard-row/

500. Keyboard Row

Given an array of strings words, return the words that can be typed using letters of the alphabet on only one row of American keyboard like the image below.

In the American keyboard:

the first row consists of the characters "qwertyuiop",
the second row consists of the characters "asdfghjkl", and
the third row consists of the characters "zxcvbnm".



Example 1:

Input: words = ["Hello","Alaska","Dad","Peace"]
Output: ["Alaska","Dad"]
Example 2:

Input: words = ["omk"]
Output: []
Example 3:

Input: words = ["adsdf","sfd"]
Output: ["adsdf","sfd"]


Constraints:

1 <= words.length <= 20
1 <= words[i].length <= 100
words[i] consists of English letters (both lowercase and uppercase).
*/

package main

func findWords(words []string) []string {
	keyMap := map[byte]int {
		'q': 1, 'w': 1, 'e': 1, 'r': 1, 't': 1, 'y': 1, 'u': 1, 'i': 1, 'o': 1, 'p': 1,
		'a': 2, 's': 2, 'd': 2, 'f': 2, 'g': 2, 'h': 2, 'j': 2, 'k': 2, 'l': 2,
		'z': 3, 'x': 3, 'c': 3, 'v': 3, 'b': 3, 'n': 3, 'm': 3,
	}
	var currentKey int;
	ans := []string{}
	for _, word := range words {
		currentFirstLetter := word[0]
		if currentFirstLetter < 97 {
			currentFirstLetter+=32
		}
		currentKey = keyMap[currentFirstLetter]
		correctWords := 0
		for _, char := range word {
			lowerCharKey := byte(char)
			if lowerCharKey < 97 {
				lowerCharKey+=32
			}
			if keyMap[lowerCharKey] == currentKey {
				correctWords++
			}
		}

		if (correctWords == len(word)) {
			ans = append(ans, word)
		}
	}
	return ans
}

// Space Complexity: O(n)
//  Time Complexity: O(n^2)