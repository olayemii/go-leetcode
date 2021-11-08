package main

import "strings"

func numOfStrings(patterns []string, word string) int {
    count := 0;
	for _, pattern := range patterns {
		if strings.Index(word, pattern) > -1 {
			count++
		}
	}

	return count;
}