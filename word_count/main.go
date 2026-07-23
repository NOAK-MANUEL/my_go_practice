package main

import "log"

func main() {
	text, err := readFile("words.txt")
	if err != nil {
		log.Fatal(err)
	}
	text = cleanWord(text)
	wordCount := countWords(text)
	for key, value := range wordCount {
		println(key, ": ", value)
	}
}
