/*
	67. Add Binary

	https://leetcode.com/problems/add-binary/

	Given two binary strings a and b, return their sum as a binary string.



	Example 1:

	Input: a = "11", b = "1"
	Output: "100"
	Example 2:

	Input: a = "1010", b = "1011"
	Output: "10101"


	Constraints:

	1 <= a.length, b.length <= 104
	a and b consist only of '0' or '1' characters.
	Each string does not contain leading zeros except for the zero itself.
*/

package main

import (
	"strconv"
)

func addBinary(a string, b string) string {
	lenA, lenB, ans := len(a)-1, len(b)-1, ""
	sum, carry := 0,0
	for lenA > -1 || lenB > -1 {
		sum = carry
		if (lenA > -1) {
			num, _ := strconv.Atoi(string(a[lenA]))
			sum +=  num;
		}
		if (lenB > -1) {
			num, _ := strconv.Atoi(string(b[lenB]))
			sum +=  num;
		}

		carry = sum / 2
		sum = sum % 2
		ans = strconv.Itoa(sum) + ans
		lenA--
		lenB--
	}
	if (carry > 0) {
		ans = strconv.Itoa(carry) + ans;
	}
	return ans;
}

/*
	Time complexity: O(n)
	Space Complexity: O(n)
*/