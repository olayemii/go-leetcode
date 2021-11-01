/*
1002. Find Common Characters

https://leetcode.com/problems/find-common-characters/

Given a string array words, return an array of all characters that show up in all strings within the words (including duplicates). You may return the answer in any order.



Example 1:

Input: words = ["bella","label","roller"]
Output: ["e","l","l"]
Example 2:

Input: words = ["cool","lock","cook"]
Output: ["c","o"]


Constraints:

1 <= words.length <= 100
1 <= words[i].length <= 100
words[i] consists of lowercase English letters.
*/

package main

import "math";

func commonChars(words []string) []string {
	minFreqArray := [26]int{}
	res := []string{};
    for i := 0; i < 26; i++ {
		minFreqArray[i] = math.MaxInt32;
	}

	for _, word := range words {
		freqArray := [26]int{}
		for _, char := range word {
			freqArray[char-'a']++
		}

		for pos, freq := range minFreqArray {
			minFreqArray[pos] =  int(math.Min(float64(freqArray[pos]), float64(freq)))
		}

	}

	for key := range minFreqArray {
		for minFreqArray[key] > 0 {
			res = append(res, string('a'+key))
			minFreqArray[key]--;
		}
	}
	return res;
}

// Space Complexity: O(n)
// Time Complexity: O(n*m)