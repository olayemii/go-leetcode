/*
2094. Finding 3-Digit Even Numbers

https://leetcode.com/problems/finding-3-digit-even-numbers/

You are given an integer array digits, where each element is a digit. The array may contain duplicates.

You need to find all the unique integers that follow the given requirements:

The integer consists of the concatenation of three elements from digits in any arbitrary order.
The integer does not have leading zeros.
The integer is even.
For example, if the given digits were [1, 2, 3], integers 132 and 312 follow the requirements.

Return a sorted array of the unique integers.



Example 1:

Input: digits = [2,1,3,0]
Output: [102,120,130,132,210,230,302,310,312,320]
Explanation: All the possible integers that follow the requirements are in the output array.
Notice that there are no odd integers or integers with leading zeros.
Example 2:

Input: digits = [2,2,8,8,2]
Output: [222,228,282,288,822,828,882]
Explanation: The same digit can be used as many times as it appears in digits.
In this example, the digit 8 is used twice each time in 288, 828, and 882.
Example 3:

Input: digits = [3,7,5]
Output: []
Explanation: No even integers can be formed using the given digits.


Constraints:

3 <= digits.length <= 100
0 <= digits[i] <= 9
*/
package main

func findEvenNumbers(digits []int) []int {
	allowedMap := map[int]int{}
	ans := []int{}
	for _, num := range digits {
		allowedMap[num] += 1
	}
	for i := 100; i < 1000; i++ {
		if i%2 == 0 && containsDigits(i, allowedMap) {
			ans = append(ans, i)
		}
	}
	return ans
}

func containsDigits(digit int, allowedMap map[int]int) bool {
	flowMap := mapClone(allowedMap)
	for digit > 0 {
		div := digit / 10
		n := digit % 10
		if v, ok := flowMap[n]; !ok || v == 0 {
			return false
		} else {
			flowMap[n]--
		}
		digit = div
	}
	return true
}

func mapClone(src map[int]int) map[int]int {
	dst := make(map[int]int, len(src))
	for k, v := range src {
		dst[k] = v
	}
	return dst
}

// Space Complexity: O(n)
// Time Complexity: O(n)
