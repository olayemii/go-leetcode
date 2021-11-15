/*
884. Uncommon Words from Two Sentences

https://leetcode.com/problems/uncommon-words-from-two-sentences/

A sentence is a string of single-space separated words where each word consists only of lowercase letters.

A word is uncommon if it appears exactly once in one of the sentences, and does not appear in the other sentence.

Given two sentences s1 and s2, return a list of all the uncommon words. You may return the answer in any order.



Example 1:

Input: s1 = "this apple is sweet", s2 = "this apple is sour"
Output: ["sweet","sour"]
Example 2:

Input: s1 = "apple apple", s2 = "banana"
Output: ["banana"]


Constraints:

1 <= s1.length, s2.length <= 200
s1 and s2 consist of lowercase English letters and spaces.
s1 and s2 do not have leading or trailing spaces.
All the words in s1 and s2 are separated by a single space.
*/

package main

import "strings"

func uncommonFromSentences(s1 string, s2 string) []string {
	splitWord1 := strings.Split(s1+" "+s2, " ");
	hashMap, ans := map[string]int{}, []string{}

	for _, word := range splitWord1 {
		if _, ok := hashMap[word]; ok {
			hashMap[word]++;
		}else{
			hashMap[word] = 1;
		}
	}

	for word, repeats := range hashMap {
		if repeats == 1 {
			ans = append(ans, word)
		}
	} 

	return ans;
}

// Space Complexity: O(n)
// Time Complexity: O(n)
