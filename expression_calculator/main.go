package main

import (
	"fmt"

	"github.com/NOAK-MANUEL/my_go_practice/expression_calculator/datatypes"
)

type Parser struct {
	tokens []datatypes.TokenValue
	pos    int
}

func (parser *Parser) Current() datatypes.TokenValue {
	return parser.tokens[parser.pos]
}
func (parser *Parser) Next() datatypes.TokenValue {
	if parser.pos >= len(parser.tokens) {
		return datatypes.TokenValue{}
	} else {

		parser.pos++
		return parser.tokens[parser.pos]
	}
}
func (p *Parser) Factor() float64 {
	current := p.Current()

	if current.Name == datatypes.Number {
		p.Next()
		return current.Value
	}

	if current.Name == datatypes.Left_Bracket {
		p.Next()
		result := p.Expression()
		p.Next()
		return result
	}

	if current.Name == datatypes.Subtract {
		p.Next()
		return -p.Factor()
	}
	panic("Invalid data")
}
func (p *Parser) Term() float64 {
	result := p.Factor()

	current := p.Current()

	if current.Name == datatypes.Multiply || current.Name == datatypes.Divide {
		op := p.Next()
		nextValue := p.Factor()
		if op.Name == datatypes.Multiply {
			result *= nextValue
		} else {
			result /= nextValue
		}
	}
	return result
}
func (p *Parser) Expression() float64 {
	result := p.Term()

	current := p.Current()
	if current.Name == datatypes.Add || current.Name == datatypes.Subtract {
		op := p.Next()
		nextValue := p.Term()
		if op.Name == datatypes.Add {
			result += nextValue
		} else {
			result -= nextValue
		}
	}
	return result
}
func evaluate(input string) float64 {
	tokens := datatypes.Tokenize(input)
	result := &Parser{tokens: tokens, pos: 0}
	return result.Expression()
}
func main() {

	fmt.Println(evaluate("2+3*2"))
	// fmt.Println(evaluate("2+3*(2+3)/2"))
}
