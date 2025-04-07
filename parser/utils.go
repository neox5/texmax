package parser

import (
	"github.com/neox5/texmax/ast"
	"github.com/neox5/texmax/tokenizer"
)

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

func (p *Parser) parseGroupedStrict() ast.Node {
	if p.peek().Type != tokenizer.LBRACE {
		p.addError("expected '{'", p.peek().Pos)
		return nil
	}
	p.next() // consume '{'
	n := p.parseNode(LOWEST)
	if p.peek().Type != tokenizer.RBRACE {
		p.addError("expected '{'", p.peek().Pos)
		return n
	}
	p.next() // consume '}'
	return n
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
