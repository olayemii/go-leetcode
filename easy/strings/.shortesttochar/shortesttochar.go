package main

import (
	"fmt"
	"math"
)

func getMin(i int, j int, k int) int {
	if k < 0 {
		return j-i
	}
	if j - i < 0 {
		return i - k
	}

	return int(math.Min(float64(i-k), float64(j-i)))
}

func shortestToChar(s string, c byte) []int {
    i, j,k, ans := 0, 0,0, []int{}

	for j < len(s) {
		if s[j] == c {
			k = i-1
			for i <= j {
				fmt.Printf("%d and %d and %d\n", j, i, k)
				ans = append(ans, getMin(i, j, k))
				i++
			}
			j++
		}else{
			j++
		}
	}
	return ans
}

func main(){
	// shortestToChar("loveleetcode", 'e')
	shortestToChar("aaba", 'b')
}