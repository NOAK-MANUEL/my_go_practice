package main

import "strings"

func keyShift(key string, index int) int {
	keyShift := key[index%len(key)]
	return int(keyShift - 'A')
}

func shiftChar(ch byte, shift int) byte {
	if ch >= 'a' && ch <= 'z' {
		return byte((int(ch-'a')+shift+26)%26) + 'a'
	}
	if ch >= 'A' && ch <= 'Z' {
		return byte((int(ch-'A')+shift+26)%26) + 'A'
	}
	return ch
}

func vigenere_encrypt(text string, key string) string {
	var result strings.Builder
	key = strings.ToUpper(key)
	var index = 0

	for _, ch := range text {
		if ch >= 'a' || ch <= 'z' || ch >= 'A' || ch <= 'Z' {
			shift := keyShift(key, index)
			result.WriteByte(shiftChar(byte(ch), shift))
			index++
		} else {
			result.WriteRune(ch)
		}
	}
	return result.String()
}

func vigenere_decrypt(text string, key string) string {
	var result strings.Builder
	key = strings.ToUpper(key)

	var index = 0

	for _, ch := range text {
		if ch >= 'a' || ch <= 'z' || ch >= 'A' || ch <= 'Z' {
			shift := keyShift(key, index)
			result.WriteByte(shiftChar(byte(ch), -shift))
			index++
		} else {
			result.WriteRune(ch)
		}
	}
	return result.String()
}

func main() {
	text := vigenere_encrypt("Yoo what is up with u man", "MumBue")
	println(text)
	text = vigenere_decrypt(text, "MumBue")
	println(text)
}
