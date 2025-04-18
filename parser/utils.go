package parser

import (
	"github.com/neox5/texmax/ast"
	"github.com/neox5/texmax/tokenizer"
)

func (p *Parser) peek() tokenizer.Token {
	// Skip over any SPACE tokens
	for p.pos < len(p.tokens) && p.tokens[p.pos].Type == tokenizer.SPACE {
		p.pos++
	}

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

func (p *Parser) parseGroupedStrict() ast.Node {
	if p.peek().Type != tokenizer.LBRACE {
		p.addError("expected '{'", p.peek().Pos)
		return nil
	}

	p.next() // consume '{'

	// Parse the expression inside the braces
	expr := p.parseExpression()

	if p.peek().Type != tokenizer.RBRACE {
		p.addError("expected '}'", p.peek().Pos)
	}
	p.next() // consume '}'
	return expr
}

func (p *Parser) parseGroupedOrSingle() ast.Node {
	if p.peek().Type == tokenizer.LBRACE {
		return p.parseGroupedStrict()
	}
	return p.parseNode(HIGHEST)
}

func (p *Parser) parseLimits() (ast.Node, ast.Node) {
	var lower, upper ast.Node
	for {
		switch p.peek().Type {
		case tokenizer.SUPERSCRIPT:
			p.next()
			if upper != nil {
				p.addError("duplicate upper limit", p.peek().Pos)
			}
			upper = p.parseGroupedOrSingle()

		case tokenizer.SUBSCRIPT:
			p.next()
			if lower != nil {
				p.addError("duplicate lower limit", p.peek().Pos)
			}
			lower = p.parseGroupedOrSingle()

		default:
			return lower, upper
		}
	}
}

// parseOptionalArgument parses an optional argument enclosed in square brackets.
// Used for commands like \sqrt[n]{x} where [n] is an optional index.
// Returns nil if there is no optional argument.
func (p *Parser) parseOptionalArgument() ast.Node {
	// Check if the next token is a left square bracket
	if p.peek().Type != tokenizer.DELIMITER || p.peek().Value != "[" {
		return nil // No optional argument
	}

	p.next() // consume '['

	// Parse the expression inside the brackets
	expr := p.parseExpression()

	// Check for closing bracket
	if p.peek().Type != tokenizer.DELIMITER || p.peek().Value != "]" {
		p.addError("expected closing ']' for optional argument", p.peek().Pos)
		// Even if there's an error, we'll return what we parsed so far
	} else {
		p.next() // consume ']'
	}

	return expr
}
