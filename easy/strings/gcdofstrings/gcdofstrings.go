package main

import (
	"fmt"
	"math"
	"strings"
)

func gcdOfStrings(str1 string, str2 string) string {
    for i := 0; i < len(str2); i++ {
		curr := string(str2[0: len(str2)-i])
        if strings.Repeat(curr, int(math.Ceil(float64(len(str1)) / float64(len(curr))))) == str1 {
			return curr
		}
    }

	return ""
}

func main() {
	fmt.Println(gcdOfStrings("TAUXXTAUXXTAUXXTAUXXTAUXX", "TAUXXTAUXXTAUXXTAUXXTAUXXTAUXXTAUXXTAUXXTAUXX"))
}