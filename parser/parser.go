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
	p.prefix[tokenizer.OPERATOR] = p.parseOperatorSymbol
	p.prefix[tokenizer.COMMAND] = p.parseCommand
	p.prefix[tokenizer.DELIMITER] = p.parseDelimiter

	// Infix registration
	p.infix[tokenizer.SUPERSCRIPT] = p.parseSuperscript
	p.infix[tokenizer.SUBSCRIPT] = p.parseSubscript
	return p
}

func (p *Parser) Parse() (ast.Node, []ParseError) {
	expr := p.parseExpression()
	return expr, p.errors
}

// parseExpression parses a sequence of nodes that form an expression
func (p *Parser) parseExpression() *ast.ExpressionNode {
	start := p.peek().Pos
	var elements []ast.Node

	for {
		curr := p.peek()
		// Exit conditions
		if curr.Type == tokenizer.EOF ||
			curr.Type == tokenizer.RBRACE ||
			(curr.Type == tokenizer.DELIMITER && curr.Value == "]") {
			break
		}

		n := p.parseNode(LOWEST)
		if n != nil {
			elements = append(elements, n)
		} else {
			// If parseNode returns nil, we should advance past the current token
			// to avoid infinite loops (unless we're at EOF)
			if p.peek().Type != tokenizer.EOF {
				p.next()
			}
		}
	}

	return &ast.ExpressionNode{Start: start, Elements: elements}
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
