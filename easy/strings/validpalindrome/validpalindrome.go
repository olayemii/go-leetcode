/*
Given a string s, determine if it is a palindrome, considering only alphanumeric characters and ignoring cases.



Example 1:

Input: s = "A man, a plan, a canal: Panama"
Output: true
Explanation: "amanaplanacanalpanama" is a palindrome.
Example 2:

Input: s = "race a car"
Output: false
Explanation: "raceacar" is not a palindrome.


Constraints:

1 <= s.length <= 2 * 105
s consists only of printable ASCII characters.
*/

package validpalindrome

import "strings"

func isPalindrome(s string) bool {
    const alpha = "abcdefghijklmnopqrstuvwxyz0123456789"
	pointerA, pointerB := 0, len(s)-1;
	for (pointerA < pointerB) {
		var currentAChar = strings.ToLower(string(s[pointerA]));
		var currentBChar = strings.ToLower(string(s[pointerB]));
		for (!strings.Contains(alpha, currentAChar) && pointerA < pointerB) {
			pointerA++;
			currentAChar = strings.ToLower(string(s[pointerA]));
		}
		for (!strings.Contains(alpha, currentBChar) && pointerA < pointerB) {
			pointerB--;
			currentBChar = strings.ToLower(string(s[pointerB]));
		}

		if  (currentAChar != currentBChar) {
			return false;
		}
		pointerA++;
		pointerB--;
	}
    return true;
}


/*
	Time complexity O(n)
	Space complexity O(1)
*/