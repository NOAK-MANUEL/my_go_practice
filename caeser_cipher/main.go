package main

import "strings"

func main() {
	text := Caeser_Encrypt("Hello how u, what's popping", 3)
	println(text)
	text = Caeser_Decrypt(text, 3)
	println(text)
}

func ShiftChar(ch byte, shift int) byte {
	if ch >= 'a' && ch <= 'z' {
		return byte((int(ch-'a')+shift+26)%26) + 'a'
	}
	if ch >= 'A' && ch <= 'Z' {
		return byte((int(ch-'A')+shift+26)%26) + 'A'
	}
	return ch
}

func Caeser_Encrypt(text string, shift int) string {
	var result strings.Builder

	for i := 0; i < len(text); i++ {
		result.WriteByte(ShiftChar(text[i], shift))
	}
	return result.String()
}

func Caeser_Decrypt(text string, shift int) string {
	return Caeser_Encrypt(text, -shift)
}
