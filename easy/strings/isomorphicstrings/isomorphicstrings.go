/*
	205. Isomorphic Strings

	https://leetcode.com/problems/isomorphic-strings/

	Given two strings s and t, determine if they are isomorphic.

	Two strings s and t are isomorphic if the characters in s can be replaced to get t.

	All occurrences of a character must be replaced with another character while preserving the order of characters. No two characters may map to the same character, but a character may map to itself.



	Example 1:

	Input: s = "egg", t = "add"
	Output: true
	Example 2:

	Input: s = "foo", t = "bar"
	Output: false
	Example 3:

	Input: s = "paper", t = "title"
	Output: true


	Constraints:

	1 <= s.length <= 5 * 104
	t.length == s.length
	s and t consist of any valid ascii character.
*/
package main

func isIsomorphic(s string, t string) bool {
	set := map[rune]byte{}
	values := map[byte]rune{}
	for idx, char := range s {
		if val, ok := set[char]; ok {
			if (val != t[idx]) {
				return false;
			}
		}else{
			if _, ok := values[t[idx]]; ok {
				return false
			}
			set[char] = t[idx]
			values[t[idx]] = char;
		}
	}
    return true;
}

/*
	Time Complexity: O(n)
	Space Complexity: O(n)
*/