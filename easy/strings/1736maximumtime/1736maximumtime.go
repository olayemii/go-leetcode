/*
1736. Latest Time by Replacing Hidden Digits

https://leetcode.com/problems/latest-time-by-replacing-hidden-digits/

You are given a string time in the form of hh:mm, where some of the digits in the string are hidden (represented by ?).

The valid times are those inclusively between 00:00 and 23:59.

Return the latest valid time you can get from time by replacing the hidden digits.

 

Example 1:

Input: time = "2?:?0"
Output: "23:50"
Explanation: The latest hour beginning with the digit '2' is 23 and the latest minute ending with the digit '0' is 50.
Example 2:

Input: time = "0?:3?"
Output: "09:39"
Example 3:

Input: time = "1?:22"
Output: "19:22"
 

Constraints:

time is in the format hh:mm.
It is guaranteed that you can produce a valid time from the given string.
*/
package main

func maximumTime(time string) string {
    timeRune := []rune(time)
    for pos, char := range timeRune {
        if char == '?' {
            if pos == 0 && (timeRune[1] <= '3' || timeRune[1] == '?') {
                timeRune[pos] = '2'
            }else
            if pos == 0 {
                timeRune[pos] = '1'
            }else
            if pos == 1 && timeRune[0] == '2' {
                timeRune[pos] = '3'
            }else
            if pos == 1 {
                timeRune[pos] = '9'
            }else
            if pos == 3 {
                timeRune[pos] = '5'
            }else if (pos == 4){
                timeRune[pos] = '9'
            }
        }
    }
    
    return string(timeRune)
}

// Space Complexity: O(n)
// Time Complexity: O(n)