/*
1360. Number of Days Between Two Dates

https://leetcode.com/problems/number-of-days-between-two-dates/

Write a program to count the number of days between two dates.

The two dates are given as strings, their format is YYYY-MM-DD as shown in the examples.

 

Example 1:

Input: date1 = "2019-06-29", date2 = "2019-06-30"
Output: 1
Example 2:

Input: date1 = "2020-01-15", date2 = "2019-12-31"
Output: 15
 

Constraints:

The given dates are valid dates between the years 1971 and 2100.
*/
package main

import (
	"time"
	"math"
)

func daysBetweenDates(date1 string, date2 string) int {
    const layout = "2006-01-02"

    t1, _ := time.Parse(layout, date1)
    t2, _ := time.Parse(layout, date2)
    return int(math.Abs(t2.Sub(t1).Hours())) / 24
}

// Time COmplexity: O(1)
// Space COmplexity: O(1)