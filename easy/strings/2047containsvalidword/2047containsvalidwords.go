/**
2047. Number of Valid Words in a Sentence

https://leetcode.com/problems/number-of-valid-words-in-a-sentence/

A sentence consists of lowercase letters ('a' to 'z'), digits ('0' to '9'), hyphens ('-'), punctuation marks ('!', '.', and ','), and spaces (' ') only. Each sentence can be broken down into one or more tokens separated by one or more spaces ' '.

A token is a valid word if all three of the following are true:

It only contains lowercase letters, hyphens, and/or punctuation (no digits).
There is at most one hyphen '-'. If present, it must be surrounded by lowercase characters ("a-b" is valid, but "-ab" and "ab-" are not valid).
There is at most one punctuation mark. If present, it must be at the end of the token ("ab,", "cd!", and "." are valid, but "a!b" and "c.," are not valid).
Examples of valid words include "a-b.", "afad", "ba-c", "a!", and "!".

Given a string sentence, return the number of valid words in sentence.



Example 1:

Input: sentence = "cat and  dog"
Output: 3
Explanation: The valid words in the sentence are "cat", "and", and "dog".
Example 2:

Input: sentence = "!this  1-s b8d!"
Output: 0
Explanation: There are no valid words in the sentence.
"!this" is invalid because it starts with a punctuation mark.
"1-s" and "b8d" are invalid because they contain digits.
Example 3:

Input: sentence = "alice and  bob are playing stone-game10"
Output: 5
Explanation: The valid words in the sentence are "alice", "and", "bob", "are", and "playing".
"stone-game10" is invalid because it contains digits.
Example 4:

Input: sentence = "he bought 2 pencils, 3 erasers, and 1  pencil-sharpener."
Output: 6
Explanation: The valid words in the sentence are "he", "bought", "pencils,", "erasers,", "and", and "pencil-sharpener.".


Constraints:

1 <= sentence.length <= 1000
sentence only contains lowercase English letters, digits, ' ', '-', '!', '.', and ','.
There will be at least 1 token.
*/
package main

import (
	"fmt"
	"strings"
)

func countValidWords(sentence string) int {
	splitted, count := strings.Split(sentence, " "), 0

	for _, word := range splitted {
		if word != "" && (containsValidChar(word) && containsValidPunctuation(word)){
			count++
			fmt.Println(word)
		}
	}
	return count;
}

func containsValidChar(token string) bool {
    for _, char := range token {
        if !(char >= 'a' && char <= 'z') && !(char == '-') && !(char == '!') && !(char == '.') && !(char == ',')  {
            return false
        }
    }
    
    return true
}

func containsValidPunctuation(token string) bool {
    hyphenCount, punctuationMarkCount, hyphenTrue, punctuationTrue := 0, 0, false, false;
    for _, char := range token {
        if char == '-' {
            hyphenCount++
        }else
        if char == '!' || char == ',' || char == '.' ||  char == ' ' {
            punctuationMarkCount++
        } 
    }
	lastItem := token[len(token)-1]
	if (hyphenCount == 0 || (hyphenCount == 1 && len(token) > 1 && (token[0] != '-' && lastItem != '-'))) {
		hyphenTrue = true
	}

	if (punctuationMarkCount == 0 || (punctuationMarkCount == 1 && (lastItem == '!' || lastItem == ',' || lastItem == '.'))){
		punctuationTrue = true
	}

	if len(token) > 2 && ((lastItem == '!' || lastItem == ',' || lastItem == '.') && token[len(token)-2] == '-') {
		return false;
	}

	return hyphenTrue && punctuationTrue

}


// Space Complexity: O(n)
// Time Complexity: O(n*m)