/*
20. Valid Parentheses

https://leetcode.com/problems/valid-parentheses/

Given a string s containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.

An input string is valid if:

Open brackets must be closed by the same type of brackets.
Open brackets must be closed in the correct order.


Example 1:

Input: s = "()"
Output: true
Example 2:

Input: s = "()[]{}"
Output: true
Example 3:

Input: s = "(]"
Output: false
Example 4:

Input: s = "([)]"
Output: false
Example 5:

Input: s = "{[]}"
Output: true


Constraints:

1 <= s.length <= 104
s consists of parentheses only '()[]{}'.
*/

package main

func isValid(s string) bool {
    stack := make([]rune, 0, len(s))

    for _, c := range s {
        if c == '(' || c == '[' || c == '{' {
            stack = append(stack, c)
            continue
        }



        if len(stack) == 0 || c == ')' && stack[len(stack)-1] != '(' ||
           c == ']' && stack[len(stack)-1] != '[' ||
           c == '}' && stack[len(stack)-1] != '{' {
                return false
        }
        stack = stack[:len(stack)-1]
    }

    return len(stack) == 0
}