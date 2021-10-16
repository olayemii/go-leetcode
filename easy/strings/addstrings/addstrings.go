/*
	415. Add Strings

	https://leetcode.com/problems/add-strings/

	Given two non-negative integers, num1 and num2 represented as string, return the sum of num1 and num2 as a string.

	You must solve the problem without using any built-in library for handling large integers (such as BigInteger). You must also not convert the inputs to integers directly.



	Example 1:

	Input: num1 = "11", num2 = "123"
	Output: "134"
	Example 2:

	Input: num1 = "456", num2 = "77"
	Output: "533"
	Example 3:

	Input: num1 = "0", num2 = "0"
	Output: "0"


	Constraints:

	1 <= num1.length, num2.length <= 104
	num1 and num2 consist of only digits.
	num1 and num2 don't have any leading zeros except for the zero itself.
*/
package main

import (
	"strconv"
)

func addStrings(num1 string, num2 string) string {
	carry, sum, lenA, lenB := 0, "", len(num1) - 1, len(num2) - 1;
	for lenA >= 0 || lenB >= 0 || carry > 0 {
		currentSum := carry;
		if (lenA >= 0) {
			currentSum += int(num1[lenA] - '0');
		}
		if (lenB >= 0) {
			currentSum += int(num2[lenB] - '0');
		}

		carry = currentSum / 10
		sum =  strconv.Itoa(currentSum%10) + sum

		lenA--;
		lenB--;
	}
	return sum;
}

// Space complexity: 0(1)
// Time Complexity O(n)