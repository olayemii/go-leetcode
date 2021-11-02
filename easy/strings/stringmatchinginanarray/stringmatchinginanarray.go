/*
1408. String Matching in an Array

https://leetcode.com/problems/string-matching-in-an-array/

Given an array of string words. Return all strings in words which is substring of another word in any order.

String words[i] is substring of words[j], if can be obtained removing some characters to left and/or right side of words[j].



Example 1:

Input: words = ["mass","as","hero","superhero"]
Output: ["as","hero"]
Explanation: "as" is substring of "mass" and "hero" is substring of "superhero".
["hero","as"] is also a valid answer.
Example 2:

Input: words = ["leetcode","et","code"]
Output: ["et","code"]
Explanation: "et", "code" are substring of "leetcode".
Example 3:

Input: words = ["blue","green","bu"]
Output: []
*/
package main

import "strings"

func stringMatching(words []string) []string {
	res, appearedStrings := []string{}, map[string]bool{}
	for _, word := range words {
		for _, word2 := range words {
			if _, ok := appearedStrings[word2]; !ok && strings.Index(word, word2) >= 0  && word != word2 {
				res = append(res, word2)
				appearedStrings[word2] = true
			}
		}
	}

	return res;
}

// Time Complexity: O(n^2)
// Space Complexity: O(n)