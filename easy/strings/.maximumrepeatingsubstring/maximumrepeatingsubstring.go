package main

import "fmt"

func maxRepeating(sequence string, word string) int {
	i, j, count := 0, 0, 0

	for i < len(word) {
		if sequence[j] == word[i] {
			fmt.Println(string(sequence[j]))
			fmt.Println(string(word[i]))
			i++
			j++
			if j == len(sequence) {
				j = 0
				count++
			}
		}else{
			i++
			j = 0
		}
	}

	return count
}

func main(){
	fmt.Println(maxRepeating("aaaba", "aaabaaaabaaabaaaabaaaabaaaabaaaaba"))
}
