package parser

import "github.com/neox5/texmax/ast"

func (p *Parser) parseSuperscript(left ast.Node) ast.Node {
	p.next() // consume '^'
	exp := p.parseGroupedOrSingle()
	return &ast.SuperscriptNode{Base: left, Exponent: exp}
}

func (p *Parser) parseSubscript(left ast.Node) ast.Node {
	p.next() // consume '_'
	idx := p.parseGroupedOrSingle()
	return &ast.SubscriptNode{Base: left, Subscript: idx }
}
