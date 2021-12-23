/*
2108. Find First Palindromic String in the Array

https://leetcode.com/problems/find-first-palindromic-string-in-the-array/

Given an array of strings words, return the first palindromic string in the array. If there is no such string, return an empty string "".

A string is palindromic if it reads the same forward and backward.



Example 1:

Input: words = ["abc","car","ada","racecar","cool"]
Output: "ada"
Explanation: The first string that is palindromic is "ada".
Note that "racecar" is also palindromic, but it is not the first.
Example 2:

Input: words = ["notapalindrome","racecar"]
Output: "racecar"
Explanation: The first and only string that is palindromic is "racecar".
Example 3:

Input: words = ["def","ghi"]
Output: ""
Explanation: There are no palindromic strings, so the empty string is returned.


Constraints:

1 <= words.length <= 100
1 <= words[i].length <= 100
words[i] consists only of lowercase English letters.
*/
package firstpalindrome

func firstPalindrome(words []string) string {
    for _, word := range words {
        if (isPalindrome(word)) {
            return word
        }
    }
    
    return ""
}

func isPalindrome(word string) bool {
    i,j := 0, len(word) - 1
    
    for i < j {
        if word[i] != word[j] {
            return false
        }
        i++
        j--
    }
    
    return true
}

// Space Complexity: O(1)
// Time Complexity: O(n)