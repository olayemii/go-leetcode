/*
1108. Defanging an IP Address

https://leetcode.com/problems/defanging-an-ip-address/

Given a valid (IPv4) IP address, return a defanged version of that IP address.

A defanged IP address replaces every period "." with "[.]".



Example 1:

Input: address = "1.1.1.1"
Output: "1[.]1[.]1[.]1"
Example 2:

Input: address = "255.100.50.0"
Output: "255[.]100[.]50[.]0"


Constraints:

The given address is a valid IPv4 address.

*/
package main

func defangIPaddr(address string) string {
	str := ""

	for _, char := range address {
		if char == '.' {
			str += "[.]"
		}else{
			str += string(char)
		}
	}
	return str
}

// Space complexity: O(n)
//  Time Complexity: O(n)