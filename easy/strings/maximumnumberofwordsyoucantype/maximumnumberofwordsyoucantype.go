package main

import (
	"strings"
)

func canBeTypedWords(text string, brokenLetters string) int {
    splitWord := strings.Split(text, " ")
    count := 0

	brokenMap := map[rune]bool{}
	for _, char := range brokenLetters {
		brokenMap[char] = true
	}
    
    for _, word := range splitWord {
        if similarities(brokenMap, word) {
			count++;
		}
    }


	return count;
}

func similarities(brokenMap map[rune]bool, text string) bool {
	for _, v := range text {
		if _, ok := brokenMap[v]; ok {
			return false
		}
	}
	
	return true;
}

func main() {
	canBeTypedWords("hello world", "ad")
}