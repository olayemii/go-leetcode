/*
917. Reverse Only Letters

https://leetcode.com/problems/reverse-only-letters/

Given a string s, reverse the string according to the following rules:

All the characters that are not English letters remain in the same position.
All the English letters (lowercase or uppercase) should be reversed.
Return s after reversing it.



Example 1:

Input: s = "ab-cd"
Output: "dc-ba"
Example 2:

Input: s = "a-bC-dEf-ghIj"
Output: "j-Ih-gfE-dCba"
Example 3:

Input: s = "Test1ng-Leet=code-Q!"
Output: "Qedo1ct-eeLg=ntse-T!"


Constraints:

1 <= s.length <= 100
s consists of characters with ASCII values in the range [33, 122].
s does not contain '\"' or '\\'.
*/
package main

func isAlpha(curr byte) bool {
    return curr >= 'A' && curr <= 'Z' || curr >= 'a' && curr <= 'z'
}

func reverseOnlyLetters(s string) string {
    i, j := 0, len(s) - 1
    charArr := []byte(s)
    for i < j {
        for i < j && !isAlpha(s[i]) {
            i++
        }
        for i < j && !isAlpha(s[j]) {
            j--
        }
        
        charArr[i], charArr[j] = charArr[j], charArr[i]
        i++
        j--
    }
    
    return string(charArr)
}

// Space Complexity: O(n)
// Time Complexity: O(n)