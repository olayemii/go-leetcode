/*
1154. Day of the Year

https://leetcode.com/problems/day-of-the-year/

Given a string date representing a Gregorian calendar date formatted as YYYY-MM-DD, return the day number of the year.



Example 1:

Input: date = "2019-01-09"
Output: 9
Explanation: Given date is the 9th day of the year in 2019.
Example 2:

Input: date = "2019-02-10"
Output: 41
Example 3:

Input: date = "2003-03-01"
Output: 60
Example 4:

Input: date = "2004-03-01"
Output: 61


Constraints:

date.length == 10
date[4] == date[7] == '-', and all other date[i]'s are digits
date represents a calendar date between Jan 1st, 1900 and Dec 31, 2019.
*/
package main

import (
	"strconv"
	"strings"
)

func dayOfYear(date string) int {
	ans := 0;
    days := strings.Split(date, "-")
    year, _ := strconv.Atoi(days[0])
    month, _ := strconv.Atoi(days[1])
	day, _ := strconv.Atoi(days[2])
	isLeap := (((year % 4 == 0) && (year % 100!= 0)) || (year%400 == 0))
	for i := 1; i <= month; i++ {
		if month != i {
			if i == 2 {
				ans += 28
			}else if i == 1 || i == 3 || i == 5 || i == 7 || i == 8 || i == 10 || i == 12 {
				ans += 31
			}else{
				ans += 30
			}
		}

	}
	ans += day
	if isLeap {
		ans++
	}

	return ans;
}

// Time Complexity: O(n)
// Space Complexity: O(1)