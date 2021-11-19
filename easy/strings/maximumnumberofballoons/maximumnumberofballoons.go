/*
1189. Maximum Number of Balloons

https://leetcode.com/problems/maximum-number-of-balloons/

Given a string text, you want to use the characters of text to form as many instances of the word "balloon" as possible.

You can use each character in text at most once. Return the maximum number of instances that can be formed.



Example 1:



Input: text = "nlaebolko"
Output: 1
Example 2:



Input: text = "loonbalxballpoon"
Output: 2
Example 3:

Input: text = "leetcode"
Output: 0


Constraints:

1 <= text.length <= 104
text consists of lower case English letters only.
*/

package main

import "math"

func maxNumberOfBalloons(text string) int {
    freqArr := [26]int{}

	for _, char := range text {
		freqArr[char-'a']++
	}

    arr := [5]int{
        freqArr['b'-'a'],
        freqArr['a'-'a'],
        freqArr['l'-'a']/2,
        freqArr['o'-'a']/2,
        freqArr['n'-'a'],
    }
    return findMin(arr)

}

func findMin(arr [5]int) int {
	min := math.MaxInt64
	for _, v := range arr {
		min = int(math.Min(float64(v), float64(min)))
	}
	return min;
}


// Space Complexity: O(1)
// Time Complexity: O(n)