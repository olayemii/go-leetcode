/*
28. Implement strStr()

https://leetcode.com/problems/implement-strstr/


Implement strStr().

Return the index of the first occurrence of needle in haystack, or -1 if needle is not part of haystack.

Clarification:

What should we return when needle is an empty string? This is a great question to ask during an interview.

For the purpose of this problem, we will return 0 when needle is an empty string. This is consistent to C's strstr() and Java's indexOf().



Example 1:

Input: haystack = "hello", needle = "ll"
Output: 2
Example 2:

Input: haystack = "aaaaa", needle = "bba"
Output: -1
Example 3:

Input: haystack = "", needle = ""
Output: 0


Constraints:

0 <= haystack.length, needle.length <= 5 * 104
haystack and needle consist of only lower-case English characters.

*/
package main

func strStr(haystack string, needle string) int {
	if (needle == "") {
		return 0;
	}
    if (len(needle) > len(haystack)) {
        return -1;
    }
    possibleSubstring := haystack[0 : len(haystack)-len(needle)+1]
	for pos, char := range possibleSubstring {
		if byte(char) == needle[0] && haystack[pos: pos+len(needle)] == needle {
			return pos;
		}
	}
	return -1;
}

/*
	Space Complexity: 0(1)
	Time Complexity: 0(n)
*/