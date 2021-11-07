package main

import (
	"strings"
)

func longestNiceSubstring(s string) string {
	i, j, ans, s := 0, 0, "", s+" "

	for j < len(s) {
		curr := s[i]
		if curr >= 'a' && curr <= 'z' {
			curr-=32
		}else{
			curr+=32
		}

		if s[j] != s[i] && s[j] != curr {
			currSubstr := s[i:j]
			if strings.Index(currSubstr, string(curr)) > -1 && strings.Index(currSubstr, string(s[i])) > -1 {
				if len(currSubstr) > len(ans) {
					ans = currSubstr
				}
			}
			i = j;
		} else{
			j++
		}
	}


	return ans;
}

func main() {
	longestNiceSubstring("YazaAay")
}