package parser

import (
	"fmt"

	"github.com/neox5/texmax/ast"
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

// func (p *Parser) expect(expected tokenizer.TokenType) bool {
// 	if p.peek().Type == expected {
// 		p.next()
// 		return true
// 	}
//
// 	p.addError(fmt.Errorf("expected token %s, got %s at position %d", expected, p.peek().Type, p.peek().Pos))
// 	return false
// }

func (p *Parser) addError(err error) {
	p.errors = append(p.errors, err)
}

func (p *Parser) Parse() (ast.Node, []error) {
	root := p.parseMathExpression()
	return root, p.errors
}

func (p *Parser) parseMathExpression() *ast.MathExpressionNode {
	start := 0
	if len(p.tokens) > 0 {
		start = p.tokens[0].Pos
	}

	root := &ast.MathExpressionNode{
		Start:    start,
		Elements: []ast.Node{},
	}

	for p.peek().Type != tokenizer.EOF {
		el := p.parseElement()
		if el != nil {
			root.Elements = append(root.Elements, el)
		} else {
			p.next() // skip problematic token to avoid infinite loop
		}
	}

	return root
}

func (p *Parser) parseElement() ast.Node {
	t := p.peek()

	switch t.Type {
	case tokenizer.SYMBOL:
		t = p.next()
		return &ast.SymbolNode{Start: t.Pos, Value: t.Value}
	case tokenizer.NUMBER:
		t = p.next()
		return &ast.NumberNode{Start: t.Pos, Value: t.Value}
	default:
		p.addError(fmt.Errorf("unexpected token %s at position %d", t.Type, t.Pos))
		return nil
	}
}
