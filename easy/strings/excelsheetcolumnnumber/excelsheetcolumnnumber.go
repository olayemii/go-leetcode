/*

	171. Excel Sheet Column Number


	https://leetcode.com/problems/excel-sheet-column-number/submissions/


	Given a string columnTitle that represents the column title as appear in an Excel sheet, return its corresponding column number.

	For example:

	A -> 1
	B -> 2
	C -> 3
	...
	Z -> 26
	AA -> 27
	AB -> 28
	...


	Example 1:

	Input: columnTitle = "A"
	Output: 1
	Example 2:

	Input: columnTitle = "AB"
	Output: 28
	Example 3:

	Input: columnTitle = "ZY"
	Output: 701
	Example 4:

	Input: columnTitle = "FXSHRXW"
	Output: 2147483647


	Constraints:

	1 <= columnTitle.length <= 7
	columnTitle consists only of uppercase English letters.
	columnTitle is in the range ["A", "FXSHRXW"].


*/

package main

import (
	"math"
)

func titleToNumber(columnTitle string) int {
	strLen := len(columnTitle) - 1;
	sum := 0;
	for pos, item := range columnTitle {
		sum += int(item-64) * int(math.Pow(26, float64(strLen - pos))) 
	} 
    return sum;
}


// Space Complexity O(1)
// Time Complexity O(n)