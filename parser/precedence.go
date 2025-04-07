package parser

import "github.com/neox5/texmax/tokenizer"

const (
	LOWEST = iota
	SCRIPT // precedence for ^ and _
	HIGHEST
)

var precedences = map[tokenizer.TokenType]int{
	tokenizer.SUPERSCRIPT: SCRIPT,
	tokenizer.SUBSCRIPT:   SCRIPT,
}

func (p *Parser) peekPrecedence() int {
	if p, ok := precedences[p.peek().Type]; ok {
		return p
	}
	return LOWEST
}
