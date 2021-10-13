/*
	242. Valid Anagram

	https://leetcode.com/problems/valid-anagram/

	Given two strings s and t, return true if t is an anagram of s, and false otherwise.



	Example 1:

	Input: s = "anagram", t = "nagaram"
	Output: true
	Example 2:

	Input: s = "rat", t = "car"
	Output: false


	Constraints:

	1 <= s.length, t.length <= 5 * 104
	s and t consist of lowercase English letters.


	Follow up: What if the inputs contain Unicode characters? How would you adapt your solution to such a case?
*/

package main

func isAnagram(s string, t string) bool {
    if (len(s) != len(t)) {
		return false
	}
	hashmap := map[rune]int{}
	for pos, char := range s {
		if _, ok := hashmap[char]; ok {
			hashmap[char] = hashmap[char] + 1
			if (hashmap[char] == 0) {
				delete(hashmap, char)
			}
		}else{
			hashmap[char] = 1
		}
		if _, ok := hashmap[rune(t[pos])]; ok {
			hashmap[rune(t[pos])] = hashmap[rune(t[pos])] - 1
			if (hashmap[rune(t[pos])] == 0) {
				delete(hashmap, rune(t[pos]))
			}
		}else{
			hashmap[rune(t[pos])] = -1
		}
	}
	return len(hashmap) == 0
}

/*
	Time Complexity: O(n)
	Space Complexity: O(n) 
*/
