/*
	344. Reverse String

	https://leetcode.com/problems/reverse-string/

	Write a function that reverses a string. The input string is given as an array of characters s.



	Example 1:

	Input: s = ["h","e","l","l","o"]
	Output: ["o","l","l","e","h"]
	Example 2:

	Input: s = ["H","a","n","n","a","h"]
	Output: ["h","a","n","n","a","H"]


	Constraints:

	1 <= s.length <= 105
	s[i] is a printable ascii character.


	Follow up: Do not allocate extra space for another array. You must do this by modifying the input array in-place with O(1) extra memory.
*/

package main

import "fmt"

func reverseString(s []byte)  {
    pos := 0
    for pos < len(s) / 2 {
        lastIndex := len(s) - pos - 1
        s[pos], s[lastIndex] = s[lastIndex],s[pos]
        pos++
    } 
    
    fmt.Println(s);
}

/*
	Reverse is done in-place
	
	Space Complexity: O(1)
	Time Complexity: O(n) 
*/