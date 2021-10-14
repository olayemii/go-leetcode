/*
	345. Reverse Vowels of a String

	https://leetcode.com/problems/reverse-vowels-of-a-string/

	Given a string s, reverse only all the vowels in the string and return it.

	The vowels are 'a', 'e', 'i', 'o', and 'u', and they can appear in both cases.



	Example 1:

	Input: s = "hello"
	Output: "holle"
	Example 2:

	Input: s = "leetcode"
	Output: "leotcede"


	Constraints:

	1 <= s.length <= 3 * 105
	s consist of printable ASCII characters.
*/
package main

import (
	"strings"
	"unicode"
)

func reverseVowels(s string) string {
	strArr := strings.Split(s, "")
	pointerA := 0;
	pointerB := len(s) - 1;
    
	for pointerA < pointerB {
		for pointerA < pointerB && !isVowel(strArr[pointerA]) {
			pointerA++;
		}
		for pointerA < pointerB && !isVowel(strArr[pointerB]) {
			pointerB--;
		}
		strArr[pointerA], strArr[pointerB] = strArr[pointerB], strArr[pointerA]

		pointerA++
		pointerB--
	}

	return strings.Join(strArr, "")
}

func isVowel(char string) bool {
	c := unicode.ToLower(rune(char[0]))
	if c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u' {
		return true
	}
	return false
}

// Space Complexity: O(n)
// Time Complexity: O(n)