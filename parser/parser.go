package parser

import (
	"fmt"

	"github.com/neox5/texmax/tokenizer"
)

type Parser struct {
	tokens []tokenizer.Token
	curr   int
	errors []error
}

func New(ts []tokenizer.Token) *Parser {
	return &Parser{
		tokens: ts,
		curr:   0,
		errors: []error{},
	}
}

func (p *Parser) peek() tokenizer.Token {
	if p.curr >= len(p.tokens) {
		return tokenizer.Token{Type: tokenizer.EOF, Value: "", Pos: -1}
	}
	return p.tokens[p.curr]
}

func (p *Parser) next() tokenizer.Token {
	t := p.peek()
	p.curr++
	return t
}

func (p *Parser) expect(expected tokenizer.TokenType) bool {
	if p.peek().Type == expected {
		p.next()
		return true
	}

	p.addError(fmt.Errorf("expected token %s, got %s at position %d",expected, p.peek().Type, p.peek().Pos))
	return false
}

func (p *Parser) addError(err error) {
	p.errors = append(p.errors, err)
}
