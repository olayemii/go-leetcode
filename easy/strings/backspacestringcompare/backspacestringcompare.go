/*
844. Backspace String Compare

https://leetcode.com/problems/backspace-string-compare/

Given two strings s and t, return true if they are equal when both are typed into empty text editors. '#' means a backspace character.

Note that after backspacing an empty text, the text will continue empty.

 

Example 1:

Input: s = "ab#c", t = "ad#c"
Output: true
Explanation: Both s and t become "ac".
Example 2:

Input: s = "ab##", t = "c#d#"
Output: true
Explanation: Both s and t become "".
Example 3:

Input: s = "a##c", t = "#a#c"
Output: true
Explanation: Both s and t become "c".
Example 4:

Input: s = "a#c", t = "b"
Output: false
Explanation: s becomes "c" while t becomes "b".
 

Constraints:

1 <= s.length, t.length <= 200
s and t only contain lowercase letters and '#' characters.
 

Follow up: Can you solve it in O(n) time and O(1) space?
*/

package main

/*Stack Solution*/

func backspaceCompare(s string, t string) bool {
    return stackify(s) == stackify(t)
}

func stackify(str string) string {
    stack := []rune{}
    
    for _, char := range str {
        if char == '#' && len(stack) > 0 {
            stack = stack[0: len(stack)-1 ]
        }else
        if len(stack) == 1 && stack[0] == '#' {
            stack[0] = char   
        }else{
            stack = append(stack, char)
        }
    }
    
    return string(stack)
}
/*End Stack Solution*/
