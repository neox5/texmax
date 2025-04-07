package parser

import (
	"fmt"

	"github.com/neox5/texmax/ast"
	"github.com/neox5/texmax/tokenizer"
)

type ParseError struct {
	Pos     int
	Message string
}

func (e ParseError) String() string {
	return fmt.Sprintf("%s at position %d", e.Message, e.Pos)
}

type Parser struct {
	tokens []tokenizer.Token
	pos    int
	prefix map[tokenizer.TokenType]func() ast.Node
	infix  map[tokenizer.TokenType]func(ast.Node) ast.Node
	errors []ParseError
}

func New(ts []tokenizer.Token) *Parser {
	p := &Parser{
		tokens: ts,
		pos:    0,
		prefix: make(map[tokenizer.TokenType]func() ast.Node),
		infix:  make(map[tokenizer.TokenType]func(ast.Node) ast.Node),
		errors: []ParseError{},
	}

	// Prefix registration
	p.prefix[tokenizer.SYMBOL] = p.parseSymbol

	return p
}

func (p *Parser) addError(pos int, msg string) {
	p.errors = append(p.errors, ParseError{pos, msg})
}

func (p *Parser) peek() tokenizer.Token {
	if p.pos >= len(p.tokens) {
		return tokenizer.Token{Type: tokenizer.EOF, Value: "", Pos: -1}
	}
	return p.tokens[p.pos]
}

func (p *Parser) next() tokenizer.Token {
	t := p.peek()
	p.pos++
	return t
}

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

func (p *Parser) Parse() (ast.Node, []ParseError) {
	var nodes []ast.Node
	for p.peek().Type != tokenizer.EOF {
		n := p.parseExpression(LOWEST)
		if n != nil {
			nodes = append(nodes, n)
		}
	}
	return &ast.RowNode{Elements: nodes}, p.errors
}

func (p *Parser) parseExpression(precedence int) ast.Node {
	t := p.peek()

	prefix := p.prefix[t.Type]
	if prefix == nil {
		p.addError(t.Pos, "no prefix token: "+t.Value)
		return nil
	}

	left := prefix()

	for precedence < p.peekPrecedence() {
		infix := p.infix[p.peek().Type]
		if infix == nil {
			break
		}
		left = infix(left)
	}

	return left
}
