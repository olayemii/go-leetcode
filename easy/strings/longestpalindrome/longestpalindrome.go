/*
	409. Longest Palindrome

	https://leetcode.com/problems/longest-palindrome/

	Given a string s which consists of lowercase or uppercase letters, return the length of the longest palindrome that can be built with those letters.

	Letters are case sensitive, for example, "Aa" is not considered a palindrome here.



	Example 1:

	Input: s = "abccccdd"
	Output: 7
	Explanation:
	One longest palindrome that can be built is "dccaccd", whose length is 7.
	Example 2:

	Input: s = "a"
	Output: 1
	Example 3:

	Input: s = "bb"
	Output: 2


	Constraints:

	1 <= s.length <= 2000
	s consists of lowercase and/or uppercase English letters only.
*/
package main

func longestPalindrome(s string) int {
	freqMap := map[rune]int{}
	longest := 0;
	for _, elem := range s {
		if _, ok := freqMap[elem]; ok {
			freqMap[elem]++
		}else{
			freqMap[elem] = 1;
		}
	}
	for _, val := range freqMap {
		longest+=(val/2) * 2
		if (val%2 == 1 && longest%2 ==0) {
			longest+=1;
		}
	}
	return longest;
}

// Space complexity: O(n)
// Time complexity: O(n)

