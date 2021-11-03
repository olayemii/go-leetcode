/*
2053. Kth Distinct String in an Array

https://leetcode.com/problems/kth-distinct-string-in-an-array/submissions/

A distinct string is a string that is present only once in an array.

Given an array of strings arr, and an integer k, return the kth distinct string present in arr. If there are fewer than k distinct strings, return an empty string "".

Note that the strings are considered in the order in which they appear in the array.



Example 1:

Input: arr = ["d","b","c","b","c","a"], k = 2
Output: "a"
Explanation:
The only distinct strings in arr are "d" and "a".
"d" appears 1st, so it is the 1st distinct string.
"a" appears 2nd, so it is the 2nd distinct string.
Since k == 2, "a" is returned.
Example 2:

Input: arr = ["aaa","aa","a"], k = 1
Output: "aaa"
Explanation:
All strings in arr are distinct, so the 1st string "aaa" is returned.
Example 3:

Input: arr = ["a","b","a"], k = 3
Output: ""
Explanation:
The only distinct string is "b". Since there are fewer than 3 distinct strings, we return an empty string "".


Constraints:

1 <= k <= arr.length <= 1000
1 <= arr[i].length <= 5
arr[i] consists of lowercase English letters.
*/

package main

func kthDistinct(arr []string, k int) string {
	hashMap, kTh := map[string]int{}, 0
    for _, item := range arr {
		if _, ok := hashMap[item]; ok {
			hashMap[item]++;
		}else{
			hashMap[item] = 1;
		}
	}


    for _, item := range arr {
		if value, ok := hashMap[item]; ok && value == 1 {
			kTh++;
		}

		if kTh == k {
			return item;
		}
	}

	return "";
}

// Space Complexity: O(n)
// Time Complexity: O(n)
