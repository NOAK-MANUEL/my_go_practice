package main

func main() {

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

func CaeserEncrypt(text string, shift int) []byte {
	values := make([]byte, len(text))

	for i := 0; i < len(text); i++ {
		values = append(values, ShiftChar(text[i], shift))
	}
	return values
}

func CaeserDecrypt(text string, shift int) []byte {
	return CaeserEncrypt(text, -shift)
}
