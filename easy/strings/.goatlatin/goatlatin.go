package main

func toGoatLatin(sentence string) string {
	result, substr := []string{}, []rune{}

	for _, char := range sentence {
		substr = append(substr, char)
		if isVowel(substr[0]) {
		}
	}

	return string(result[0])
}

func isVowel(char rune) bool {
	if char >= 'a' && char <= 'z' {
		char-=32
	}
	return char == 'A' || char == 'E' || char == 'I' || char == 'O' || char == 'U'
}