package parser

import (
	"github.com/neox5/texmax/ast"
	"github.com/neox5/texmax/tokenizer"
)

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
	p.prefix[tokenizer.LBRACE] = p.parseGroupedStrict
	p.prefix[tokenizer.SYMBOL] = p.parseSymbol
	p.prefix[tokenizer.NUMBER] = p.parseNumber
	p.prefix[tokenizer.SPACE] = p.parseSpace

	// Infix registration
	p.infix[tokenizer.SUPERSCRIPT] = p.parseSuperscript
	p.infix[tokenizer.SUBSCRIPT] = p.parseSubscript
	return p
}

func (p *Parser) Parse() (ast.Node, []ParseError) {
	var nodes []ast.Node
	for p.peek().Type != tokenizer.EOF {
		n := p.parseNode(LOWEST)
		if n != nil {
			nodes = append(nodes, n)
		}
	}
	return &ast.RowNode{Elements: nodes}, p.errors
}

func (p *Parser) parseNode(precedence int) ast.Node {
	t := p.peek()

	prefix := p.prefix[t.Type]
	if prefix == nil {
		p.addError("no prefix token: "+t.Value, t.Pos)
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
