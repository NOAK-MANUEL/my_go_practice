package datatypes

import (
	"log"
	"regexp"
)

type Token int

const (
	Number Token = iota
	Subtract
	Add
	Divide
	Multiply
	Left_Bracket
	Right_Bracket
)

type TokenValue struct {
	Name  Token
	Value float64
}

func Tokenize(input string) []TokenValue {

	tokens := []TokenValue{}
	for i := 0; i < len(input); {
		ch := input[i]

		switch {
		case ch == '+':
			i++
			tokens = append(tokens, TokenValue{Name: Add})
		case ch == '-':
			i++
			tokens = append(tokens, TokenValue{Name: Subtract})
		case ch == '/':
			i++
			tokens = append(tokens, TokenValue{Name: Divide})
		case ch == '(':
			i++
			tokens = append(tokens, TokenValue{Name: Left_Bracket})
		case ch == ')':
			i++
			tokens = append(tokens, TokenValue{Name: Right_Bracket})
		case ch == ' ':
			i++
		case regexp.MustCompile("/d").MatchString(string(ch)):
			i++
			tokens = append(tokens, TokenValue{Name: Number, Value: float64(ch)})
		default:
			log.Fatal("Invalid figure", ch)
		}
	}
	return tokens
}
