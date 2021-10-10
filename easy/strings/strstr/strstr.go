package main

func strStr(haystack string, needle string) int {
	if (needle == "") {
		return 0;
	}
    if (len(needle) > len(haystack)) {
        return -1;
    }
    possibleSubstring := haystack[0 : len(haystack)-len(needle)+1]
	for pos, char := range possibleSubstring {
		if byte(char) == needle[0] && haystack[pos: pos+len(needle)] == needle {
			return pos;
		}
	}
	return -1;
}
