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
	return fmt.Sprintf("%s at position %d",e.Message, e.Pos)
}

type Parser struct {
	tokens []tokenizer.Token
	curr   int
	errors []ParseError
}

func New(ts []tokenizer.Token) *Parser {
	return &Parser{
		tokens: ts,
		curr:   0,
		errors: []ParseError{},
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

func (p *Parser) addError(pos int, msg string) {
	p.errors = append(p.errors, ParseError{pos, msg})
}

func (p *Parser) Parse() (ast.Node, []ParseError) {
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
	case tokenizer.SPACE:
		t = p.next()
		return &ast.SpaceNode{Start: t.Pos, Value: t.Value}
	case tokenizer.SYMBOL:
		t = p.next()
		return &ast.SymbolNode{Start: t.Pos, Value: t.Value}
	case tokenizer.NUMBER:
		t = p.next()
		return &ast.NumberNode{Start: t.Pos, Value: t.Value}
	case tokenizer.OPERATOR:
		t = p.next()
		return &ast.OperatorNode{Start: t.Pos, Value: t.Value}
	default:
		p.addError(t.Pos, fmt.Sprintf("unexpected token %s", t.Type))
		return nil
	}
}
