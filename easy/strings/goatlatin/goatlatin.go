/*
824. Goat Latin

https://leetcode.com/problems/goat-latin/

You are given a string sentence that consist of words separated by spaces. Each word consists of lowercase and uppercase letters only.

We would like to convert the sentence to "Goat Latin" (a made-up language similar to Pig Latin.) The rules of Goat Latin are as follows:

If a word begins with a vowel ('a', 'e', 'i', 'o', or 'u'), append "ma" to the end of the word.
For example, the word "apple" becomes "applema".
If a word begins with a consonant (i.e., not a vowel), remove the first letter and append it to the end, then add "ma".
For example, the word "goat" becomes "oatgma".
Add one letter 'a' to the end of each word per its word index in the sentence, starting with 1.
For example, the first word gets "a" added to the end, the second word gets "aa" added to the end, and so on.
Return the final sentence representing the conversion from sentence to Goat Latin.



Example 1:

Input: sentence = "I speak Goat Latin"
Output: "Imaa peaksmaaa oatGmaaaa atinLmaaaaa"
Example 2:

Input: sentence = "The quick brown fox jumped over the lazy dog"
Output: "heTmaa uickqmaaa rownbmaaaa oxfmaaaaa umpedjmaaaaaa overmaaaaaaa hetmaaaaaaaa azylmaaaaaaaaa ogdmaaaaaaaaaa"


Constraints:

1 <= sentence.length <= 150
sentence consists of English letters and spaces.
sentence has no leading or trailing spaces.
All the words in sentence are separated by a single space.

*/

package main

import (
	"strings"
)

func toGoatLatin(sentence string) string {
	substr, count, ans := []rune{}, 1, ""
    sentence = sentence + " "
	for pos, char := range sentence {
		if char == ' ' {
			if isVowel(substr[0]) {
				ans += string(substr)
			}else{
				ans += string(substr[1:])+string(substr[0])
			}	
			ans += "ma"+strings.Repeat("a", count)
			if pos != len(sentence)-1 {
				ans+= " "
			}
			count++
			substr = nil
        }else{
    		substr = append(substr, char)
        }
	}

	return ans
}

func isVowel(char rune) bool {
	if char >= 'a' && char <= 'z' {
		char-=32
	}
	return char == 'A' || char == 'E' || char == 'I' || char == 'O' || char == 'U'
}

// Space Complexity: O(n)
// Time Complexity: O(n)