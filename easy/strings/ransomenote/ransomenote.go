/*
	383. Ransom Note

	https://leetcode.com/problems/ransom-note/

	Given two stings ransomNote and magazine, return true if ransomNote can be constructed from magazine and false otherwise.

	Each letter in magazine can only be used once in ransomNote.



	Example 1:

	Input: ransomNote = "a", magazine = "b"
	Output: false
	Example 2:

	Input: ransomNote = "aa", magazine = "ab"
	Output: false
	Example 3:

	Input: ransomNote = "aa", magazine = "aab"
	Output: true


	Constraints:

	1 <= ransomNote.length, magazine.length <= 105
	ransomNote and magazine consist of lowercase English letters.
*/

package main

func canConstruct(ransomNote string, magazine string) bool {
	if (len(ransomNote) > len(magazine)) {
		return false;
	}
	hashMap := map[rune]int{}
	for _, char := range ransomNote {
		if _, ok := hashMap[char]; ok {
			hashMap[char]++;
		}else{
			hashMap[char] = 1;
		}

	}
	for _, char := range magazine {
		if val, ok := hashMap[char]; ok && val > 0{
			hashMap[char]--;
		}
	}
	for _, val := range hashMap {
		if (val != 0) {
			return false;
		}
	}
    return true;
}


// Space Complexity O(n)
// Time Complexity O(n)