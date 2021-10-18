/*
709. To Lower Case

https://leetcode.com/problems/to-lower-case/

Given a string s, return the string after replacing every uppercase letter with the same lowercase letter.



Example 1:

Input: s = "Hello"
Output: "hello"
Example 2:

Input: s = "here"
Output: "here"
Example 3:

Input: s = "LOVELY"
Output: "lovely"


Constraints:

1 <= s.length <= 100
s consists of printable ASCII characters.
*/

package main

func toLowerCase(s string) string {
	strArr := []byte(s)
	for pos, char := range strArr {
		if char >= 'A' && char <= 'Z' {
			char += 32
		}

		strArr[pos] = char
	}
	return string(strArr)
}

// Time Complexity: O(n)
// Space Complexity: O(n)