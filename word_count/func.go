package main

import (
	"os"
	"strings"
)

func readFile(path string) (string, error) {
	text, err := os.ReadFile(path)

	return strings.ToLower(string(text)), err

}

func cleanWord(text string) string {
	var words strings.Builder

	for _, ch := range text {
		if (ch >= 'a' && ch <= 'z') || ch == ' ' {
			words.WriteRune(ch)
		} else {
			words.WriteRune(' ')
		}
	}
	return words.String()
}
func countWords(text string) map[string]int {
	var counts = make(map[string]int)

	words := strings.Fields(text)
	for _, word := range words {
		counts[word]++
	}
	return counts
}
